package system

import (
	"testing"
)

// --------- Instructions -----------

// jump 0x1NNN
func TestInstJump(t *testing.T) {
	c := &cpu{programCounter: programStartOffset}
	c.process(0x1225, []byte{})

	if c.programCounter != 549 {
		t.Error("Expected program counter to 549, but is", c.programCounter)
	}
}

// store NN in register X - 0x6XNN
func TestInstStoreValueInRegister(t *testing.T) {
	c := &cpu{programCounter: programStartOffset}
	c.process(0x6a08, []byte{})

	if c.registers[10] != 8 {
		t.Error("Expected register 10 to be 8, but is", c.registers[10])
	}
}

// store NNN in the index register - 0xANNN
func TestInstStoreIndexRegister(t *testing.T) {
	c := &cpu{programCounter: programStartOffset}
	c.process(0xa3d3, []byte{})

	if c.indexRegister != 979 {
		t.Error("Expected index register to be 979, but is", c.indexRegister)
	}
}

// add NN to register X - 0x7XNN
func TestInstAddToRegister(t *testing.T) {
	c := &cpu{programCounter: programStartOffset}
	c.registers[1] = byte(3)
	c.process(0x7104, []byte{})

	if c.registers[1] != 7 {
		t.Error("Expected register 1 to be 7, but is", c.registers[1])
	}
}

// adding to a register and wrapping
func TestInstAddToRegisterWithWrapping(t *testing.T) {
	c := &cpu{programCounter: programStartOffset}
	c.registers[1] = byte(254)
	c.process(0x7104, []byte{})

	if c.registers[1] != 2 {
		t.Error("Expected register 1 to be 2, but is", c.registers[1])
	}
}

// add register X to index register - 0xFX1E
func TestInstAddRegisterToIndexRegister(t *testing.T) {
	c := &cpu{programCounter: programStartOffset, indexRegister: 100}
	c.registers[0] = byte(12)
	c.process(0xf01e, []byte{})

	if c.indexRegister != 112 {
		t.Error("Expected index register to be 112, but is", c.indexRegister)
	}
}

// skip the next instruction if register equals value - 0x3000
func TestInstSkipNextIfRegisterEqualsValue(t *testing.T) {
	c := &cpu{programCounter: programStartOffset}

	c.registers[0] = byte(12)
	c.process(0x300c, []byte{})
	if c.programCounter != programStartOffset+4 {
		t.Error("Expected the program counter to advance 4 bytes, but advanced", c.programCounter-programStartOffset)
	}

	c.programCounter = programStartOffset
	c.registers[0] = byte(12)
	c.process(0x3007, []byte{})
	if c.programCounter != programStartOffset+2 {
		t.Error("Expected the program counter to advance 2 bytes, but advanced", c.programCounter-programStartOffset)
	}
}

// set the delay timer to the value of a register - 0xFX15
func TestSetDelayTimer(t *testing.T) {
	c := &cpu{programCounter: programStartOffset}

	c.registers[3] = 0xfe
	c.process(0xf315, []byte{})

	if c.delayTimer != 0xfe {
		t.Error("Expected the delay timer to be 0xfe, but was", c.delayTimer)
	}
}

// set a register to the value of the delay timer - 0xFX07
func TestSetRegisterFromDelayTimer(t *testing.T) {
	c := &cpu{programCounter: programStartOffset}

	c.delayTimer = 0xb3
	c.process(0xf207, []byte{})

	if c.registers[2] != 0xb3 {
		t.Error("Expected the register to be 0xb3, but was", c.registers[2])
	}
}

func TestDrawInstructionXorPixels(t *testing.T) {
	c := &cpu{programCounter: programStartOffset}
	c.screenState = make([][]byte, DisplayHeight)
	for i := range c.screenState {
		c.screenState[i] = make([]byte, DisplayWidth)
	}

	// set up the memory
	mem := make([]byte, memorySize)
	mem[0x204] = 0x80 // pixel

	c.registers[0] = 10
	c.registers[1] = 10
	c.indexRegister = 0x204

	// draw twice to make sure the hit flag is set and the sprite is removed
	c.process(0xd011, mem)
	if c.screenState[10][10] != 1 {
		t.Error("Expected the pixel to be set, but it was not")
	}
	if c.registers[15] != 0 {
		t.Error("Expected the last register to be 0, but it was not")
	}

	c.process(0xd011, mem)
	if c.screenState[10][10] != 0 {
		t.Error("Expected the pixel to be unset, but it was set")
	}
	if c.registers[15] != 1 {
		t.Error("Expected the last register to be 1, but it was not")
	}
}
