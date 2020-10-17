package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

var gobfile = "pages.gob"

var m = make(map[string]*Page)

func savePages() {

	dataFile, err := os.Create(gobfile)
	if err != nil {
		log.Println(err)
	}
	defer dataFile.Close()

	// serialize the data
	dataEncoder := gob.NewEncoder(dataFile)
	dataEncoder.Encode(m)
}

func getPages() {

	// open data file
	dataFile, err := os.Open(gobfile)
	if err != nil {
		log.Println(err)
	}
	defer dataFile.Close()

	dataDecoder := gob.NewDecoder(dataFile)
	err = dataDecoder.Decode(&m)

	if err != nil {
		fmt.Println(err)
	}
}
