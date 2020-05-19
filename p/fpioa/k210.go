// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build k210

// Package fpioa provides access to the registers of the FPIOA peripheral.
//
// Instances:
//  FPIOA  FPIOA_BASE  APB0  -  Field Programmable IO Array
// Registers:
//  0x000 32  IO[48]      FPIOA GPIO multiplexer io array
//  0x0C0 32  TIE_EN[8]   FPIOA GPIO multiplexer tie enable array
//  0x0E0 32  TIE_VAL[8]  FPIOA GPIO multiplexer tie value array
// Import:
//  github.com/embeddedgo/kendryte/p/bus
//  github.com/embeddedgo/kendryte/p/mmap
package fpioa

const (
	CH_SEL IO = 0xFF << 0  //+ Channel select from 256 input
	DS     IO = 0x0F << 8  //+ Driving selector
	OE_EN  IO = 0x01 << 12 //+ Static output enable, will AND with OE_INV
	OE_INV IO = 0x01 << 13 //+ Invert output enable
	DO_SEL IO = 0x01 << 14 //+ Data output select: 0 for DO, 1 for OE
	DO_INV IO = 0x01 << 15 //+ Invert the result of data output select (DO_SEL)
	PU     IO = 0x01 << 16 //+ Pull up enable. 0 for nothing, 1 for pull up
	PD     IO = 0x01 << 17 //+ Pull down enable. 0 for nothing, 1 for pull down
	SL     IO = 0x01 << 19 //+ Slew rate control enable
	IE_EN  IO = 0x01 << 20 //+ Static input enable, will AND with IE_INV
	IE_INV IO = 0x01 << 21 //+ Invert input enable
	DI_INV IO = 0x01 << 22 //+ Invert Data input
	ST     IO = 0x01 << 23 //+ Schmitt trigger
	PAD_DI IO = 0x01 << 31 //+ Read current IO's data input
)

const (
	CH_SELn = 0
	DSn     = 8
	OE_ENn  = 12
	OE_INVn = 13
	DO_SELn = 14
	DO_INVn = 15
	PUn     = 16
	PDn     = 17
	SLn     = 19
	IE_ENn  = 20
	IE_INVn = 21
	DI_INVn = 22
	STn     = 23
	PAD_DIn = 31
)
