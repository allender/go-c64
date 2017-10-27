package cpu

import "github.com/allender/go-c64/memory"

// Cpu is the main structure for the 6502 emulation
type Cpu struct {

	// PC is the program counter
	PC uint16

	// SP is the stack pointer
	SP uint8

	// AC is the accumulator register
	AC uint8

	// X is the x index register
	X uint8

	// Y is the y index register
	Y uint8

	// SR is the status register
	SR uint8
}

// Addr_mode_t is a function type for addressing modes
type Addr_mode_t func(*Cpu) uint8

func (c *Cpu) Imm() uint8 {
	return memory.Read(c.PC)
}

func (c *Cpu) Zpg() uint8 {
	return memory.Read(uint16(memory.Read(c.PC)))
}

func (c *Cpu) Zpx() uint8 {
	return memory.Read(uint16(memory.Read(c.PC) + c.X))
}

func (c *Cpu) Zpy() uint8 {
	return memory.Read(uint16(memory.Read(c.PC) + c.Y))
}

func (c *Cpu) Inx() uint8 {
	addr := uint16(memory.Read(c.PC)) + uint16(c.X)
	low := memory.Read(addr)
	high := memory.Read(addr + 1)
	return memory.Read((uint16(high) << 8) | uint16(low))
}

func (c *Cpu) Iny() uint8 {
	addr := uint16(memory.Read(c.PC))
	low := memory.Read(addr)
	high := memory.Read(addr + 1)
	return memory.Read(((uint16(high) << 8) | uint16(low)) + uint16(c.Y))
}

func (c *Cpu) Abs() uint8 {
	low := memory.Read(c.PC)
	high := memory.Read(c.PC + 1)
	return memory.Read((uint16(high) << 8) | uint16(low))
}

func (c *Cpu) Abx() uint8 {
	low := memory.Read(c.PC)
	high := memory.Read(c.PC + 1)
	return memory.Read((uint16(high) << 8) | uint16(low) + uint16(c.X))
}

func (c *Cpu) Aby() uint8 {
	low := memory.Read(c.PC)
	high := memory.Read(c.PC + 1)
	return memory.Read((uint16(high) << 8) | uint16(low) + uint16(c.Y))
}

func (c *Cpu) Ind() uint8 {
	low := memory.Read(c.PC)
	high := memory.Read(c.PC + 1)
	addr := (uint16(high) << 8) | uint16(low)
	return memory.Read(uint16(memory.Read(addr)) | (uint16(memory.Read(addr+1)) << 8))
}
