package system

import (
	"errors"
	"io/ioutil"
)

const (
	// The maximum size of memory available to the system in bytes
	memorySize = 4096

	// Game data starts at 0x200. The ROM should be dumped into memory starting at this location. This is also
	// where the emulator should start executing instructions from.
	programStartOffset = 0x200
)

// System is the emulator
type System struct {
	// handles emulating the CPU and it's instructions
	cpu *cpu

	// Game data starts at 0x200. 0x00 - 0x1FF are reserved by the system.
	// the contents of the ROM will be dumped into here.
	memory [memorySize]byte
}

// NewSystem initializes a new Chip-8 emulator system and returns it
func NewSystem(romFile string) (*System, error) {
	sys := &System{cpu: &cpu{}}

	// place the rom into the system's memory
	if err := sys.loadRom(romFile); err != nil {
		return sys, err
	}

	return sys, nil
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
