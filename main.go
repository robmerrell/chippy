package main

import (
	"fmt"
	"github.com/robmerrell/chippy/system"
)

func main() {
	sys, err := system.NewSystem("./community_roms/INVADERS")
	if err != nil {
		fmt.Println(err)
	}

	sys.Run()
}
