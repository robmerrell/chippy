package system

import (
	"fmt"
)

// CPU represents the current state of the chip-8 CPU. A pretty complete description of the system can be found here: https://en.wikipedia.org/wiki/CHIP-8#Virtual_machine_description
type cpu struct {
	// 16 registers V0 - VF, where VF is commonly the carry flag.
	registers [16]byte

	// Index register - store a memory address
	indexRegister uint16

	// Program counter to keep track of the next instruction to read
	programCounter uint16

	// Call stack - 16 levels.
	stack        [16]uint16
	stackPointer byte

	// Timers
	// once set counts down each cycle until 0
	delayTimer byte

	// Screen state of each pixel. Since there are no colors a pixel can either be on or off. Perhaps a bool would be better here...
	screenState [][]byte

	// We only need to draw after specific registers are processed. The system will read this draw flag, draw to the screen
	// and then set it to false after it has drawn.
	drawFlag bool
}

func (c *cpu) process(instruction uint16, memory []byte) {
	fmt.Printf("Instruction 0x%4x at %d\n", instruction, c.programCounter)

	// instruction handling. these are in alphabetical order to keep things easy to find.
	switch instruction & 0xF000 {
	case 0x0000:
		switch instruction & 0x00FF {

		// (0x00EE) return from the subroutine
		case 0x00EE:
			c.programCounter = c.stack[c.stackPointer-1]
			c.stackPointer--
			return

		// 0x0NNN is not implemented because it was used to execute machine code on the
		// the COSMAC VIP machine. Most modern emulators do not implement it.
		default:
			fmt.Println("Instruction not implemented")
		}

	// (0x1NNN) Jump to address NNN
	case 0x1000:
		c.programCounter = instruction & 0x0FFF
		return

	// (0x2NNN) Execute the subroutine starting at address NNN
	case 0x2000:
		c.stack[c.stackPointer] = c.programCounter + 2
		c.programCounter = instruction & 0x0FFF
		c.stackPointer++
		return

	// (3XNN) skip the following instruction if register X equals NN
	case 0x3000:
		register := (instruction & 0x0F00) >> 8
		value := instruction & 0x00FF
		if c.registers[register] == byte(value) {
			c.advancePC()
		}

	// (0x6XNN) store the value NN in register X
	case 0x6000:
		register := (instruction & 0x0F00) >> 8
		value := instruction & 0x00FF
		c.registers[register] = byte(value)

	// (0x7XNN) add the value NN to register X
	case 0x7000:
		register := (instruction & 0x0F00) >> 8
		value := instruction & 0x00FF
		c.registers[register] += byte(value)

	// (0xANNN) store the address NNN in the index register
	case 0xA000:
		c.indexRegister = instruction & 0x0FFF

	// (DXYN) draw a sprite at the position in registers X,Y with N bytes of sprite data
	case 0xD000:
		height := instruction & 0x000F
		regX := (instruction & 0x0F00) >> 8
		regY := (instruction & 0x00F0) >> 4
		x := c.registers[regX]
		y := c.registers[regY]

		c.registers[0xF] = 0 // hit detection flag
		for row := uint16(0); row < height; row++ {
			spriteData := memory[c.indexRegister+row]

			for col := uint16(0); col < 8; col++ {
				// expand out the sprite and place the pixel, which will be a 0 or a 1 in the screenstate.
				// if we were parsing 0x3C The screenstate for that section of the sprite will be 00111100
				inv := 7 - col // without reading the inverse (just col) we get our bits backwards. This fixes that.
				pixel := spriteData & (1 << inv) >> inv
				if pixel == 1 {
					xIndex := col + uint16(x)
					yIndex := row + uint16(y)

					// if the pixel was already one, set the last register to 1 to show that a collision happened
					if c.screenState[yIndex][xIndex] == 1 {
						c.registers[0xF] = 1
					}

					c.screenState[yIndex][xIndex] ^= 1
				}
			}
		}

		c.drawFlag = true

	case 0xF000:
		switch instruction & 0x00FF {

		// (0xFX1E) Add the value from register X to the index register
		// TODO: Wikipedia has a note about an undocumented feature that the carry flag should be set when overflowing
		// not sure if it is for chip-8 or just super chip-8
		case 0x001E:
			register := (instruction & 0x0F00) >> 8
			c.indexRegister += uint16(c.registers[register])

		// (0xFX07) Set register X to the value in the display timer
		case 0x0007:
			register := (instruction & 0x0F00) >> 8
			c.registers[register] = c.delayTimer

		// (0xFX15) Set the delay timer to the value of register X
		case 0x0015:
			register := (instruction & 0x0F00) >> 8
			c.delayTimer = c.registers[register]

		default:
			fmt.Println("Instruction not implemented")
		}

	default:
		fmt.Println("Instruction not implemented")
	}

	c.advancePC()
}

// advancePC advances the program counter by 2 bytes (because all instructions are 2 bytes)
func (c *cpu) advancePC() {
	c.programCounter += 2
}
