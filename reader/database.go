package cpu_database_reader

import (
	"encoding/csv"
	"strconv"
	"strings"
	"fmt"
	"io"
)

type CPUData struct {
	name 		string
	author 		string
	speed 		float64
	registers  	uint64
	ram 		uint64
	rom 		uint64
	dw_bits 	uint64
	iw_bits 	uint64
	image 		string
	video 		string
	isa 		string
	description string
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
	if len(strs) < 12 {
		return nil, fmt.Errorf("Not enough fields, only got %v (%v)", len(strs), strs)
	}
	d := CPUData {}

	speed, _ 	:= strconv.ParseFloat(strs[2], 64)
	regs, _  	:= strconv.ParseUint(strs[3], 10, 64)
	ram, _  	:= strconv.ParseUint(strs[4], 10, 64)
	rom, _  	:= strconv.ParseUint(strs[5], 10, 64)
	dw, _  		:= strconv.ParseUint(strs[6], 10, 64)
	iw, _  		:= strconv.ParseUint(strs[7], 10, 64)

	d.name 		= strs[0]
	d.author 	= strs[1]
	d.speed 	= speed 	// 2
	d.registers = regs 		// 3
	d.ram 		= ram 		// 4
	d.rom 		= rom 		// 5
	d.dw_bits 	= dw 		// 6
	d.iw_bits 	= iw   		// 7
	d.image 	= strs[8]
	d.video 	= strs[9]
	d.isa 		= strs[10]
	d.description = strs[11]

	return &d, nil
}
