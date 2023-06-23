package main

import (
	"github.com/funnsam/cpu_db/reader"
	"github.com/rivo/tview"
	"github.com/gdamore/tcell/v2"
	"fmt"
)

func center(p tview.Primitive, width, height int) tview.Primitive {
	return tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(p, height, 1, true).
		AddItem(nil, 0, 1, false), width, 1, true).
		AddItem(nil, 0, 1, false)
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
	_, _, w, h := pages.GetInnerRect()

	list.SetHighlightFullLine(true)
	list.SetMainTextStyle(tcell.StyleDefault.Attributes(tcell.AttrBold))
	list.SetMainTextColor(tcell.ColorWhite)
	list.SetSecondaryTextColor(tcell.NewHexColor(0xBFBFBF))

	details.SetBorder(true)
	details.SetDynamicColors(true)

	pages.AddPage("CPU List", list, true, true)
	pages.AddPage("Details", center(details, w*w/h, h*w/h), true, false)

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
		fs := "[DarkSlateGray::b]"
		fe := "[-:-:-]"

		fmt.Fprintf(w, "%sName:%s %s%s%s\n", ns, ne, fs, d.Name, fe)
		fmt.Fprintf(w, "%sAuthor:%s %s%s%s\n", ns, ne, fs, d.Author, fe)
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
