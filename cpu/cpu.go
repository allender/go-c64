package cpu

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
