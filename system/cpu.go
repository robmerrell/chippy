package system

// CPU represents the current state of the chip-8 CPU. A pretty complete description of the system can be found here: https://en.wikipedia.org/wiki/CHIP-8#Virtual_machine_description
type cpu struct {
	// 16 registers V0 - VF, where VF is commonly the carry flag.
	registers [16]byte

	// Index register - store a memory address
	indexRegister uint16

	// program counter to keep track of the next instruction to read
	programCounter uint16

	// call stack - 16 levels.
	stack        [16]uint16
	stackPointer byte

	// timers
}
