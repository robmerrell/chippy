package main

import (
	"github.com/robmerrell/chippy/system"
	"github.com/robmerrell/chippy/ui"
	"runtime"
)

// drawScale sets the scale size of the window
const drawScale = 16

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {
	display, err := ui.NewDisplay(system.DisplayWidth, system.DisplayHeight, drawScale)
	if err != nil {
		panic(err)
	}
	defer display.Stop()

	sys, err := system.NewSystem("./test_roms/smiley.bin", display)
	if err != nil {
		panic(err)
	}

	sys.Run()
}
