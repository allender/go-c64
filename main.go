package main

import (
	"fmt"

	"github.com/allender/go-c64/cpu"
	"github.com/allender/go-c64/memory"
)

func main() {
	memory.MemoryLoadFromFile("data/rom.bin", 0xe000)
	fmt.Printf("%v", cpu.Opcodes[5].Opcode)
}
