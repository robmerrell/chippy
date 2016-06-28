package system

import (
	"encoding/binary"
	"errors"
	"github.com/robmerrell/chippy/ui"
	"io/ioutil"
)

const (
	// The maximum size of memory available to the system in bytes
	memorySize = 4096

	// Game data starts at 0x200. The ROM should be dumped into memory starting at this location. This is also
	// where the emulator should start executing instructions from.
	programStartOffset = 0x200

	// DisplayWidth is the width in pixels
	DisplayWidth = 64

	// DisplayHeight is the height in pixels
	DisplayHeight = 32
)

// System is the emulator
type System struct {
	// Handles emulating the CPU and it's instructions
	cpu *cpu

	// Game data starts at 0x200. 0x00 - 0x1FF are reserved by the system.
	// the contents of the ROM will be dumped into here.
	memory [memorySize]byte

	// Display to draw on
	display *ui.Display
}

// NewSystem initializes a new Chip-8 emulator system and returns it
func NewSystem(romFile string, display *ui.Display) (*System, error) {
	sys := &System{cpu: &cpu{programCounter: programStartOffset}, display: display}

	// place the rom into the system's memory
	if err := sys.loadRom(romFile); err != nil {
		return sys, err
	}

	// initialize the screen. Not too worried about locality here, on a more intensive system I would be.
	sys.cpu.screenState = make([][]byte, DisplayHeight)
	for i := range sys.cpu.screenState {
		sys.cpu.screenState[i] = make([]byte, DisplayWidth)
	}

	return sys, nil
}

// Run starts the Chip-8 emulator
func (s *System) Run() {
	for {
		// each instruction is 2 bytes
		instruction := binary.BigEndian.Uint16(s.memory[s.cpu.programCounter : s.cpu.programCounter+2])
		s.cpu.process(instruction)
	}
}

// loadRom inserts the contents of the given ROM file into the system's memory
func (s *System) loadRom(romFile string) error {
	romContents, err := ioutil.ReadFile(romFile)
	if err != nil {
		return err
	}

	// ensure that the ROM will fit into memory
	if len(romContents) > memorySize-programStartOffset {
		return errors.New("ROM too large to fit into system memory")
	}

	// insert the rom into the system's memory
	insertAt := programStartOffset
	for _, b := range romContents {
		s.memory[insertAt] = b
		insertAt++
	}

	return nil
}
