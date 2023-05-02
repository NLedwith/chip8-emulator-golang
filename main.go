package main

import (
	"os"
	"io/ioutil"
)

func main() {
	filePath := ""
	if len(os.Args) > 1 {
		filePath = "./roms/" + string(os.Args[1]) + ".ch8"
	} else {
		filePath = "./roms/blitz.ch8"
	}

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	emu := Chip8Emulator{}
	emu.initialize()
	emu.start(file)
}
