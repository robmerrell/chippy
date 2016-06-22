package system

import (
	"strings"
	"testing"
)

func TestLoadingRom(t *testing.T) {
	s := &System{}
	s.loadRom("./testdata/inserting_rom")

	contents := string(s.memory[programStartOffset : programStartOffset+5])
	if contents != "12345" {
		t.Error("Expected contents to be 12345, but was", contents)
	}
}

func TestLoadingTooBigOfRom(t *testing.T) {
	s := &System{}
	err := s.loadRom("./testdata/rom_too_large")

	if err != nil && strings.Contains("too large", err.Error()) {
		t.Error("Expected a ROM too large error, but didn't get it.")
	}
}
