package cpu

const (
	_ = iota
	implied
	accumulator
	immediate
	zero_page
	zero_pagex
	zero_pagey
	indirect_x
	indirect_y
	absolute
	absolute_x
	absolute_y
	indirect
	relative
)

const (
	_ = iota
	adc
	and
	asl
	bcc
	bcs
	beq
	bit
	bmi
	bne
	bpl
	brk
	bvc
	bvs
	clc
	cld
	cli
	clv
	cmp
	cpx
	cpy
	dec
	dex
	dey
	eor
	inc
	inx
	iny
	jmp
	jsr
	lda
	ldx
	ldy
	lsr
	nop
	ora
	pha
	php
	pla
	plp
	rol
	ror
	rti
	rts
	sbc
	sec
	sed
	sei
	sta
	stx
	sty
	tax
	tay
	tsx
	txa
	txs
	tya
)

// definition for an opcode
type _Opcode struct {

	// Opcode is the hexidecimal value of the 6502 opcode
	Opcode uint8

	// Mnemonic is the constant value for the opcode.  This value can be
	// used to retrieve a string that can be used for this opcode
	Mnemonic uint8

	// AddressingMode is the addressing mode of the opcode.  This value
	// is used by the cpu to determine how to address this opcode
	AddressingMode int8

	// Cycles is the number of cycles this opcode takes
	Cycles uint8

	// Bytes is the number of bytes this opcode uses
	Bytes uint8
}

// list of the opcodes, in order of opcode value
var _opcodes = []_Opcode{
	{0x00, brk, implied, 7, 0},
	{0x01, ora, indirect_x, 6, 2},
	{0x05, ora, zero_page, 3, 2},
	{0x06, asl, zero_page, 5, 2},
	{0x08, php, implied, 3, 1},
	{0x09, ora, immediate, 2, 2},
	{0x0a, asl, accumulator, 2, 1},
	{0x0d, ora, absolute, 4, 3},
	{0x0e, asl, absolute, 6, 3},
	{0x10, bpl, relative, 2, 2},
	{0x11, ora, indirect_y, 5, 2},
	{0x15, ora, zero_pagex, 4, 2},
	{0x16, asl, zero_pagex, 6, 2},
	{0x18, clc, implied, 2, 1},
	{0x19, ora, absolute_y, 4, 3},
	{0x1d, ora, absolute_x, 4, 3},
	{0x1e, asl, absolute_x, 7, 3},
	{0x20, jsr, absolute, 6, 3},
	{0x21, and, indirect_x, 6, 2},
	{0x24, bit, zero_page, 3, 2},
	{0x25, and, zero_page, 3, 2},
	{0x26, rol, zero_page, 5, 2},
	{0x28, plp, implied, 4, 1},
	{0x29, and, immediate, 2, 2},
	{0x2a, rol, accumulator, 2, 1},
	{0x2c, bit, absolute, 4, 3},
	{0x2d, and, absolute, 4, 3},
	{0x2e, rol, absolute, 6, 3},
	{0x30, bmi, relative, 2, 2},
	{0x31, and, indirect_y, 5, 2},
	{0x35, and, zero_pagex, 4, 2},
	{0x36, rol, zero_pagex, 6, 2},
	{0x38, sec, implied, 2, 1},
	{0x39, and, absolute_y, 4, 3},
	{0x3d, and, absolute_x, 4, 3},
	{0x3e, rol, absolute_x, 7, 3},
	{0x40, rti, implied, 6, 1},
	{0x41, eor, indirect_x, 6, 2},
	{0x45, eor, zero_page, 3, 2},
	{0x46, lsr, zero_page, 5, 2},
	{0x48, pha, implied, 3, 1},
	{0x49, eor, immediate, 2, 2},
	{0x4a, lsr, accumulator, 2, 1},
	{0x4c, jmp, absolute, 3, 3},
	{0x4d, eor, absolute, 4, 3},
	{0x4e, lsr, absolute, 6, 3},
	{0x50, bvc, relative, 2, 2},
	{0x51, eor, indirect_y, 5, 2},
	{0x55, eor, zero_pagex, 4, 2},
	{0x56, lsr, zero_pagex, 6, 2},
	{0x58, cli, implied, 2, 1},
	{0x59, eor, absolute_y, 4, 3},
	{0x5d, eor, absolute_x, 4, 3},
	{0x5e, lsr, absolute_x, 7, 3},
	{0x60, rts, implied, 6, 1},
	{0x61, adc, indirect_x, 6, 2},
	{0x65, adc, zero_page, 3, 2},
	{0x66, ror, zero_page, 5, 2},
	{0x68, pla, implied, 4, 1},
	{0x69, adc, immediate, 2, 2},
	{0x6a, ror, accumulator, 2, 1},
	{0x6c, jmp, indirect, 5, 3},
	{0x6d, adc, absolute, 4, 3},
	{0x6e, ror, absolute, 6, 3},
	{0x70, bvs, relative, 2, 2},
	{0x71, adc, indirect_y, 5, 2},
	{0x75, adc, zero_pagex, 4, 2},
	{0x76, ror, zero_pagex, 6, 2},
	{0x78, sei, implied, 2, 1},
	{0x79, adc, absolute_y, 4, 3},
	{0x7d, adc, absolute_x, 4, 3},
	{0x7e, ror, absolute_x, 7, 3},
	{0x81, sta, indirect_x, 6, 2},
	{0x84, sty, zero_page, 3, 2},
	{0x85, sta, zero_page, 3, 2},
	{0x86, stx, zero_page, 3, 2},
	{0x88, dey, implied, 2, 1},
	{0x8a, txa, implied, 2, 1},
	{0x8c, sty, absolute, 4, 3},
	{0x8d, sta, absolute, 4, 3},
	{0x8e, stx, absolute, 4, 3},
	{0x90, bcc, relative, 2, 2},
	{0x91, sta, indirect_y, 6, 2},
	{0x94, sty, zero_pagex, 4, 2},
	{0x95, sta, zero_pagex, 4, 2},
	{0x96, stx, zero_pagey, 4, 2},
	{0x98, tya, implied, 2, 1},
	{0x99, sta, absolute_y, 5, 3},
	{0x9a, txs, implied, 2, 1},
	{0x9d, sta, absolute_x, 5, 3},
	{0xa0, ldy, immediate, 2, 2},
	{0xa1, lda, indirect_x, 6, 2},
	{0xa2, ldx, immediate, 2, 2},
	{0xa4, ldy, zero_page, 3, 2},
	{0xa5, lda, zero_page, 3, 2},
	{0xa6, ldx, zero_page, 3, 2},
	{0xa8, tay, implied, 2, 1},
	{0xa9, lda, immediate, 2, 2},
	{0xaa, tax, implied, 2, 1},
	{0xac, ldy, absolute, 4, 3},
	{0xad, lda, absolute, 4, 3},
	{0xae, ldx, absolute, 4, 3},
	{0xb0, bcs, relative, 2, 2},
	{0xb1, lda, indirect_y, 5, 2},
	{0xb4, ldy, zero_pagex, 4, 2},
	{0xb5, lda, zero_pagex, 4, 2},
	{0xb6, ldx, zero_pagey, 4, 2},
	{0xb8, clv, implied, 2, 1},
	{0xb9, lda, absolute_y, 4, 3},
	{0xba, tsx, implied, 2, 1},
	{0xbc, ldy, absolute_x, 4, 3},
	{0xbd, lda, absolute_x, 4, 3},
	{0xbe, ldx, absolute_y, 4, 3},
	{0xc0, cpy, immediate, 2, 2},
	{0xc1, cmp, indirect_x, 6, 2},
	{0xc4, cpy, zero_page, 3, 2},
	{0xc5, cmp, zero_page, 3, 2},
	{0xc6, dec, zero_page, 5, 2},
	{0xc8, iny, implied, 2, 1},
	{0xc9, cmp, immediate, 2, 2},
	{0xca, dex, implied, 2, 1},
	{0xcc, cpy, absolute, 4, 3},
	{0xcd, cmp, absolute, 4, 3},
	{0xce, dec, absolute, 6, 3},
	{0xd0, bne, relative, 2, 2},
	{0xd1, cmp, indirect_y, 5, 2},
	{0xd5, cmp, zero_pagex, 4, 2},
	{0xd6, dec, zero_pagex, 6, 2},
	{0xd8, cld, implied, 2, 1},
	{0xd9, cmp, absolute_y, 4, 3},
	{0xdd, cmp, absolute_x, 4, 3},
	{0xde, dec, absolute_x, 7, 3},
	{0xe0, cpx, immediate, 2, 2},
	{0xe1, sbc, indirect_x, 6, 2},
	{0xe4, cpx, zero_page, 3, 2},
	{0xe5, sbc, zero_page, 3, 2},
	{0xe6, inc, zero_page, 5, 2},
	{0xe8, inx, implied, 2, 1},
	{0xe9, sbc, immediate, 2, 2},
	{0xea, nop, implied, 2, 1},
	{0xec, cpx, absolute, 4, 3},
	{0xed, sbc, absolute, 4, 3},
	{0xee, inc, absolute, 6, 3},
	{0xf0, beq, relative, 2, 2},
	{0xf1, sbc, indirect_y, 5, 2},
	{0xf5, sbc, zero_pagex, 4, 2},
	{0xf6, inc, zero_pagex, 6, 2},
	{0xf8, sed, implied, 2, 1},
	{0xf9, sbc, absolute_y, 4, 3},
	{0xfd, sbc, absolute_x, 4, 3},
	{0xfe, inc, absolute_x, 7, 3},
}

// Opcodes is the global array mapping O(1) the hex value to the runtime execution.
var Opcodes [256]*_Opcode

func init() {
	for i := range _opcodes {
		val := _opcodes[i].Opcode
		Opcodes[val] = &_opcodes[i]
	}
}
