package reader

import (
	"encoding/csv"
	"strconv"
	"strings"
	"fmt"
	"io"
)

type CPUData struct {
	Name 		string
	Author 		string
	Speed 		float64
	Pipeline  	uint64
	Registers  	uint64
	RAM 		uint64
	ROM 		uint64
	DWBits 		uint64
	IWBits 		uint64
	Image 		string
	Video 		string
	ISA 		string
	Description string
}

func ReadDatabase() (*[]CPUData, error) {
	db_raw, err := ReadDatabaseContent()
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(strings.NewReader(*db_raw))
	_, _ = r.Read() // HACK: ignore first row

	db := make([]CPUData, 0, 100)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("Unable to read database: %v", err)
		}

		data, err := dataFromStrings(record)
		if err != nil {
			return nil, fmt.Errorf("Unable to read data from database: %v", err)
		}

		db = append(db, *data)
	}
	
	return &db, nil
}

func dataFromStrings(strs []string) (*CPUData, error) {
	if len(strs) < 13 {
		return nil, fmt.Errorf("Not enough fields, only got %v (%v)", len(strs), strs)
	}
	d := CPUData {}

	speed, _ 	:= strconv.ParseFloat(strs[2], 64)
	pipe, _  	:= strconv.ParseUint(strs[3], 10, 64)
	regs, _  	:= strconv.ParseUint(strs[4], 10, 64)
	ram, _  	:= strconv.ParseUint(strs[5], 10, 64)
	rom, _  	:= strconv.ParseUint(strs[6], 10, 64)
	dw, _  		:= strconv.ParseUint(strs[7], 10, 64)
	iw, _  		:= strconv.ParseUint(strs[8], 10, 64)

	d.Name 		= strs[0]
	d.Author 	= strs[1]
	d.Speed 	= speed 	// 2
	d.Pipeline  = pipe 		// 3
	d.Registers = regs 		// 4
	d.RAM 		= ram 		// 5
	d.ROM 		= rom 		// 6
	d.DWBits 	= dw 		// 7
	d.IWBits 	= iw   		// 8
	d.Image 	= strs[9]
	d.Video 	= strs[10]
	d.ISA 		= strs[11]
	d.Description = strs[12]

	return &d, nil
}
