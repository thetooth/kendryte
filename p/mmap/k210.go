// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build k210

// Package mmap provides base memory adresses for all peripherals.
package mmap

const (
	SPI2_BASE uintptr = 0x50240000 // Serial Peripheral Interface 2 (slave)
)

// AES
const (
	AES_BASE uintptr = 0x50450000 // AES Accelerator
)

// APU
const (
	APU_BASE uintptr = 0x50250200 // Audio Processor
)

// DMAC
const (
	DMAC_BASE uintptr = 0x50000000 // Direct Memory Access Controller
)

// DVP
const (
	DVP_BASE uintptr = 0x50430000 // Digital Video Port
)

// FFT
const (
	FFT_BASE uintptr = 0x42000000 // Fast Fourier Transform Accelerator
)

// FPIOA
const (
	FPIOA_BASE uintptr = 0x502B0000 // Field Programmable IO Array
)

// GPIO
const (
	GPIO_BASE uintptr = 0x50200000 // General Purpose Input/Output Interface
)

// GPIOHS
const (
	GPIOHS_BASE uintptr = 0x38001000 // High-speed GPIO
)

// I2C
const (
	I2C0_BASE uintptr = 0x50280000 // Inter-Integrated Circuit Bus 0
	I2C1_BASE uintptr = 0x50290000 // Inter-Integrated Circuit Bus 1
	I2C2_BASE uintptr = 0x502A0000 // Inter-Integrated Circuit Bus 2
)

// I2S
const (
	I2S0_BASE uintptr = 0x50250000 // Inter-Integrated Sound Interface 0
	I2S1_BASE uintptr = 0x50260000 // Inter-Integrated Sound Interface 1
	I2S2_BASE uintptr = 0x50270000 // Inter-Integrated Sound Interface 2
)

// KPU
const (
	KPU_BASE uintptr = 0x40800000 // Neural Network Accelerator
)

// OTP
const (
	OTP_BASE uintptr = 0x50420000 // One-Time Programmable Memory Controller
)

// RTC
const (
	RTC_BASE uintptr = 0x50460000 // Real Time Clock
)

// SHA256
const (
	SHA256_BASE uintptr = 0x502C0000 // SHA256 Accelerator
)

// SPI
const (
	SPI0_BASE uintptr = 0x52000000 // Serial Peripheral Interface 0 (master)
	SPI1_BASE uintptr = 0x53000000 // Serial Peripheral Interface 1 (master)
	SPI3_BASE uintptr = 0x54000000 // Serial Peripheral Interface 3 (master)
)

// SYSCTL
const (
	SYSCTL_BASE uintptr = 0x50440000 // System Controller
)

// TIMER
const (
	TIMER0_BASE uintptr = 0x502D0000 // Timer 0
	TIMER1_BASE uintptr = 0x502E0000 // Timer 1
	TIMER2_BASE uintptr = 0x502F0000 // Timer 2
)

// UART
const (
	UART1_BASE uintptr = 0x50210000 // Universal Asynchronous Receiver-Transmitter 1
	UART2_BASE uintptr = 0x50220000 // Universal Asynchronous Receiver-Transmitter 2
	UART3_BASE uintptr = 0x50230000 // Universal Asynchronous Receiver-Transmitter 3
)

// UARTHS
const (
	UARTHS_BASE uintptr = 0x38000000 // High-speed UART
)

// WDT
const (
	WDT0_BASE uintptr = 0x50400000 // Watchdog Timer 0
	WDT1_BASE uintptr = 0x50410000 // Watchdog Timer 1
)
