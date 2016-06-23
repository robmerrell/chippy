package system

import (
	"testing"
)

// --------- Instructions -----------

// jump 0x1NNN
func TestInstJump(t *testing.T) {
	c := &cpu{programCounter: programStartOffset}
	c.process(0x1225)

	if c.programCounter != 549 {
		t.Error("Expected program counter to 549, but is", c.programCounter)
	}
}

// store NN in register X - 0x6XNN
func TestInstStoreValueInRegister(t *testing.T) {
	c := &cpu{programCounter: programStartOffset}
	c.process(0x6a08)

	if c.registers[10] != 8 {
		t.Error("Expected register 10 to be 8, but is", c.registers[10])
	}
}

// store NNN in the index register - 0xANNN
func TestInstStoreIndexRegister(t *testing.T) {
	c := &cpu{programCounter: programStartOffset}
	c.process(0xa3d3)

	if c.indexRegister != 979 {
		t.Error("Expected index register to be 979, but is", c.indexRegister)
	}
}

// add NN to register X - 0x7XNN
func TestInstAddToRegister(t *testing.T) {
	c := &cpu{programCounter: programStartOffset}
	c.registers[1] = byte(3)
	c.process(0x7104)

	if c.registers[1] != 7 {
		t.Error("Expected register 1 to be 7, but is", c.registers[1])
	}
}

// adding to a register and wrapping
func TestInstAddToRegisterWithWrapping(t *testing.T) {
	c := &cpu{programCounter: programStartOffset}
	c.registers[1] = byte(254)
	c.process(0x7104)

	if c.registers[1] != 2 {
		t.Error("Expected register 1 to be 2, but is", c.registers[1])
	}
}

// add register X to index register - 0xFX1E
func TestInstAddRegisterToIndexRegister(t *testing.T) {
	c := &cpu{programCounter: programStartOffset, indexRegister: 100}
	c.registers[0] = byte(12)
	c.process(0xf01e)

	if c.indexRegister != 112 {
		t.Error("Expected index register to be 112, but is", c.indexRegister)
	}
}

// skip the next instruction if register equals value - 0x3000
func TestInstSkipNextIfRegisterEqualsValue(t *testing.T) {
	c := &cpu{programCounter: programStartOffset}

	c.registers[0] = byte(12)
	c.process(0x300c)
	if c.programCounter != programStartOffset+4 {
		t.Error("Expected the program counter to advance 4 bytes, but advanced", c.programCounter-programStartOffset)
	}

	c.programCounter = programStartOffset
	c.registers[0] = byte(12)
	c.process(0x3007)
	if c.programCounter != programStartOffset+2 {
		t.Error("Expected the program counter to advance 2 bytes, but advanced", c.programCounter-programStartOffset)
	}
}
