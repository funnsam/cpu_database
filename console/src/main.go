package main

import (
	"github.com/funnsam/cpu_db/reader"
	"github.com/rivo/tview"
	"github.com/gdamore/tcell/v2"
	"fmt"
)

func center(p tview.Primitive) tview.Primitive {
	return tview.NewGrid().
		SetColumns(0, -4, 0).
		SetRows(0, -4, 0).
		AddItem(p, 1, 1, 1, 1, 0, 0, true)
}

func main() {
	db, err := reader.ReadDatabase()
	if err != nil {
		panic(err)
	}

	app := tview.NewApplication()

	pages := tview.NewPages()
	list := tview.NewList()
	details := tview.NewTextView()

	list.SetHighlightFullLine(true)
	list.SetMainTextStyle(tcell.StyleDefault.Attributes(tcell.AttrBold))
	list.SetMainTextColor(tcell.ColorWhite)
	list.SetSecondaryTextColor(tcell.NewHexColor(0xAFAFAF))

	details.SetBorder(true)
	details.SetDynamicColors(true)
	details.SetTitleAlign(tview.AlignLeft)

	pages.AddPage("CPU List", list, true, true)
	pages.AddPage("Details", center(details), true, false)

	for _, v := range *db {
		list.AddItem(v.Name, fmt.Sprintf("%s, %.2f Hz", v.Author, v.Speed), 0, nil)
	}

	list.SetSelectedFunc(func(i int, _, _ string, _ rune) {
		pages.ShowPage("Details")

		w := details.BatchWriter()
		defer w.Close()
		w.Clear()
		d := (*db)[i]

		ns := ""
		ne := ""
		fs := "[#AFAFAF::]"
		fe := "[-:-:-]"

		details.SetTitle(fmt.Sprintf("%s (by %s)", d.Name, d.Author))
		fmt.Fprintf(w, "%sSpeed:%s %s%.2f Hz%s\n", ns, ne, fs, d.Speed, fe)
		if d.Pipeline != 0 {
			fmt.Fprintf(w, "%sPipeline stages:%s %s%d%s\n", ns, ne, fs, d.Pipeline, fe)
		}
		if d.Registers != 0 {
			fmt.Fprintf(w, "%sRegisters:%s %s%d%s\n", ns, ne, fs, d.Registers, fe)
		}
		fmt.Fprintf(w, "%sRAM Size:%s %s%d words%s\n", ns, ne, fs, d.RAM, fe)
		fmt.Fprintf(w, "%sROM Size:%s %s%d words%s\n", ns, ne, fs, d.ROM, fe)
		fmt.Fprintf(w, "%sData Size:%s %s%d bits%s\n", ns, ne, fs, d.DWBits, fe)
		if d.IWBits != 0 {
			fmt.Fprintf(w, "%sInstruction Size:%s %s%d bits%s\n", ns, ne, fs, d.IWBits, fe)
		}
		fmt.Fprintf(w, "%sImage:%s %s%s%s\n", ns, ne, fs, d.Image, fe)
		fmt.Fprintf(w, "%sVideo:%s %s%s%s\n", ns, ne, fs, d.Video, fe)
		fmt.Fprintf(w, "%sISA:%s %s%s%s\n", ns, ne, fs, d.ISA, fe)
		fmt.Fprintf(w, "%sDescription:%s %s%s%s\n", ns, ne, fs, d.Description, fe)
	})

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		key := event.Key()
		if key == tcell.KeyESC {
			pages.HidePage("Details")
			return nil
		}

		return event
	})

	if err := app.SetRoot(pages, true).SetFocus(pages).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
