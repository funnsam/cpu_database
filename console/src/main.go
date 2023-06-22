package main

import (
	reader "github.com/funnsam/cpu_db/reader"
	"fmt"
)

func main() {
	db, err := reader.ReadDatabase()
	if err != nil {
		panic(err)
	}
	fmt.Println(db)
}
