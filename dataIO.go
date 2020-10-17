package main

import (
	"encoding/gob"
	"log"
	"os"
	"sync"
)

var gobfile = "pages.gob"

var m = make(map[string]*Page)
var lock sync.Mutex

func savePages() {
	lock.Lock()
	defer lock.Unlock()
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
	lock.Lock()
	defer lock.Unlock()
	// open data file
	dataFile, err := os.Open(gobfile)
	if err != nil {
		log.Println(err)
	}
	defer dataFile.Close()

	dataDecoder := gob.NewDecoder(dataFile)
	err = dataDecoder.Decode(&m)

	if err != nil {
		log.Println(err)
	}
}
