package memory

import (
	"io/ioutil"
	"log"
)

var Ram [0x10000]uint8

func Read(a uint16) uint8 {
	return Ram[a]
}

func MemoryLoadFromFile(filename string, start_addr uint16) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	copy(Ram[start_addr:], data)
}
