package main

import (
	"log"
	"fmt"
	//"runtime"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"io/ioutil"
	//"strconv"
)

type Chip8Emulator struct {
	V0 uint8
	V1 uint8
	V2 uint8
	V3 uint8
	V4 uint8
	V5 uint8
	V6 uint8
	V7 uint8
	V8 uint8
	V9 uint8
	VA uint8
	VB uint8
	VC uint8
	VD uint8
	VE uint8
	VF uint8
	I uint16
	ST uint8
	DT uint8
	PC uint16
	SP uint8
	Stack [16]uint16
	RAM [4096]uint8
	Display [32][64]rune
}

func (emu *Chip8Emulator) start(program []uint8) {
	emu.load_program(program)
	running := true
	for running {
		instruction := (uint16(emu.RAM[emu.PC]) << 8) | uint16(emu.RAM[emu.PC+1])
		emu.execute_instruction(instruction)
		emu.draw_screen()
	}
}

func (emu *Chip8Emulator) load_program(program []uint8) {
	i := int(emu.PC)
	for (i-512) < len(program) {
		emu.RAM[i] = program[i-512]
		i++
	}
	log.Println("FILE LOADED")
}

func (emu *Chip8Emulator) execute_instruction(instruction uint16) {
	op := emu.get_opcode(instruction)
	switch op {
	case 0:
		emu.run_0(instruction)
	case 1:
		emu.run_1(instruction)
	case 2:
		log.Fatal(fmt.Sprintf("%x", op) + " NOT IMPLEMENTED")
	case 3:
		emu.run_3(instruction)
	case 4:
		emu.run_4(instruction)
	case 5:
		log.Fatal(fmt.Sprintf("%x", op) + " NOT IMPLEMENTED")
	case 6:
		emu.run_6(instruction)
	case 7:
		emu.run_7(instruction)
	case 8:
		log.Fatal(fmt.Sprintf("%x", op) + " NOT IMPLEMENTED")
	case 9:
		log.Fatal(fmt.Sprintf("%x", op) + " NOT IMPLEMENTED")
	case 10:
		emu.run_A(instruction)
	case 11:
		log.Fatal(fmt.Sprintf("%x", op) + " NOT IMPLEMENTED")
	case 12:
		log.Fatal(fmt.Sprintf("%x", op) + " NOT IMPLEMENTED")
	case 13:
		emu.run_D(instruction)
	case 14:
		log.Fatal(fmt.Sprintf("%x", op) + " NOT IMPLEMENTED")
	case 15:
		log.Fatal(fmt.Sprintf("%x", op) + " NOT IMPLEMENTED")
		
	default:
		log.Fatal(fmt.Sprintf("%x", op) + " NOT RECOGNIZED")
	}

}

func (emu *Chip8Emulator) get_opcode(instruction uint16) uint8 {
	return uint8(instruction >> 12)
}

func (emu *Chip8Emulator) run_0(instruction uint16) {
	check := instruction & 255
	switch check {
	case 224:
		fmt.Println(fmt.Sprintf("%x", instruction) + ": " + "CLS")
		emu.clear_screen()
		emu.PC += 2
	case 238:
		fmt.Println(fmt.Sprintf("%x", instruction) + ": " + "RET")
		emu.PC = emu.Stack[emu.SP]
		emu.SP -= 1
	default:
		fmt.Println(fmt.Sprintf("%x", instruction) + ": " + "SYS " + fmt.Sprintf("%x", instruction & 4095))
		emu.PC = instruction & 4095
	}
}

func (emu *Chip8Emulator) run_1(instruction uint16) {
	fmt.Println(emu.PC)
	fmt.Println(fmt.Sprintf("%x", instruction) + ": " + "JP " + fmt.Sprintf("%x", (instruction & 4095)))
	emu.PC = instruction & 4095
}

func (emu *Chip8Emulator) run_3(instruction uint16) {
	fmt.Println(fmt.Sprintf("%x", instruction) + ": " + "SE V" + fmt.Sprintf("%x", ((instruction & 3840) >> 8)) + ", " + fmt.Sprintf("%x", instruction & 255))
	switch ((instruction & 3840) >> 8) {
	case 0:
		if emu.V0 == uint8(instruction & 255) {
			emu.PC +=4
		} else {
			emu.PC += 2
		}
	case 1:
		if emu.V1 == uint8(instruction & 255) {
			emu.PC +=4
		} else {
			emu.PC += 2
		}
	case 2:
		if emu.V2 == uint8(instruction & 255) {
			emu.PC +=4
		} else {
			emu.PC += 2
		}
	case 3:
		if emu.V3 == uint8(instruction & 255) {
			emu.PC +=4
		} else {
			emu.PC += 2
		}
	case 4:
		if emu.V4 == uint8(instruction & 255) {
			emu.PC +=4
		} else {
			emu.PC += 2
		}
	case 5:
		if emu.V5 == uint8(instruction & 255) {
			emu.PC +=4
		} else {
			emu.PC += 2
		}
	case 6:
		if emu.V6 == uint8(instruction & 255) {
			emu.PC +=4
		} else {
			emu.PC += 2
		}
	case 7:
		if emu.V7 == uint8(instruction & 255) {
			emu.PC +=4
		} else {
			emu.PC += 2
		}
	case 8:
		if emu.V8 == uint8(instruction & 255) {
			emu.PC +=4
		} else {
			emu.PC += 2
		}
	case 9:
		if emu.V9 == uint8(instruction & 255) {
			emu.PC +=4
		} else {
			emu.PC += 2
		}
	case 10:
		if emu.VA == uint8(instruction & 255) {
			emu.PC +=4
		} else {
			emu.PC += 2
		}
	case 11:
		if emu.VB == uint8(instruction & 255) {
			emu.PC +=4
		} else {
			emu.PC += 2
		}
	case 12:
		if emu.VC == uint8(instruction & 255) {
			emu.PC +=4
		} else {
			emu.PC += 2
		}
	case 13:
		if emu.VD == uint8(instruction & 255) {
			emu.PC +=4
		} else {
			emu.PC += 2
		}
	case 14:
		if emu.VE == uint8(instruction & 255) {
			emu.PC +=4
		} else {
			emu.PC += 2
		}
	case 15:
		if emu.VF == uint8(instruction & 255) {
			emu.PC +=4
		} else {
			emu.PC += 2
		}
	default:
		log.Fatal(fmt.Sprintf("%x", ((instruction & 3040) >> 8)) + " REGISTER NOT RECOGNIZED")
		
	}
}

func (emu *Chip8Emulator) run_4(instruction uint16) {
	fmt.Println(fmt.Sprintf("%x", instruction) + ": " + "SE V" + fmt.Sprintf("%x", ((instruction & 3840) >> 8)) + ", " + fmt.Sprintf("%x", instruction & 255))
	switch ((instruction & 3840) >> 8) {
	case 0:
		if emu.V0 != uint8(instruction & 255) {
			emu.PC +=4
		} else {
			emu.PC += 2
		}
	case 1:
		if emu.V1 != uint8(instruction & 255) {
			emu.PC +=4
		} else {
			emu.PC += 2
		}
	case 2:
		if emu.V2 != uint8(instruction & 255) {
			emu.PC +=4
		} else {
			emu.PC += 2
		}
	case 3:
		if emu.V3 != uint8(instruction & 255) {
			emu.PC +=4
		} else {
			emu.PC += 2
		}
	case 4:
		if emu.V4 != uint8(instruction & 255) {
			emu.PC +=4
		} else {
			emu.PC += 2
		}
	case 5:
		if emu.V5 != uint8(instruction & 255) {
			emu.PC +=4
		} else {
			emu.PC += 2
		}
	case 6:
		if emu.V6 != uint8(instruction & 255) {
			emu.PC +=4
		} else {
			emu.PC += 2
		}
	case 7:
		if emu.V7 != uint8(instruction & 255) {
			emu.PC +=4
		} else {
			emu.PC += 2
		}
	case 8:
		if emu.V8 != uint8(instruction & 255) {
			emu.PC +=4
		} else {
			emu.PC += 2
		}
	case 9:
		if emu.V9 != uint8(instruction & 255) {
			emu.PC +=4
		} else {
			emu.PC += 2
		}
	case 10:
		if emu.VA != uint8(instruction & 255) {
			emu.PC +=4
		} else {
			emu.PC += 2
		}
	case 11:
		if emu.VB != uint8(instruction & 255) {
			emu.PC +=4
		} else {
			emu.PC += 2
		}
	case 12:
		if emu.VC != uint8(instruction & 255) {
			emu.PC +=4
		} else {
			emu.PC += 2
		}
	case 13:
		if emu.VD != uint8(instruction & 255) {
			emu.PC +=4
		} else {
			emu.PC += 2
		}
	case 14:
		if emu.VE != uint8(instruction & 255) {
			emu.PC +=4
		} else {
			emu.PC += 2
		}
	case 15:
		if emu.VF != uint8(instruction & 255) {
			emu.PC +=4
		} else {
			emu.PC += 2
		}
	default:
		log.Fatal(fmt.Sprintf("%x", ((instruction & 3040) >> 8)) + " REGISTER NOT RECOGNIZED")
		
	}
}

func (emu *Chip8Emulator) run_6(instruction uint16) {
	fmt.Println(fmt.Sprintf("%x", instruction) + ": " + "LD V" + fmt.Sprintf("%x", ((instruction & 3840) >> 8)) + ", " + fmt.Sprintf("%x", instruction & 255))
	switch ((instruction & 3840) >> 8) {
	case 0:
		emu.V0 = uint8(instruction & 255)
	case 1:
		emu.V1 = uint8(instruction & 255)
	case 2:
		emu.V2 = uint8(instruction & 255)
	case 3:
		emu.V3 = uint8(instruction & 255)
	case 4:
		emu.V4 = uint8(instruction & 255)
	case 5:
		emu.V5 = uint8(instruction & 255)
	case 6:
		emu.V6 = uint8(instruction & 255)
	case 7:
		emu.V7 = uint8(instruction & 255)
	case 8:
		emu.V8 = uint8(instruction & 255)
	case 9:
		emu.V9 = uint8(instruction & 255)
	case 10:
		emu.VA = uint8(instruction & 255)
	case 11:
		emu.VB = uint8(instruction & 255)
	case 12:
		emu.VC = uint8(instruction & 255)
	case 13:
		emu.VD = uint8(instruction & 255)
	case 14:
		emu.VE = uint8(instruction & 255)
	case 15:
		emu.VF = uint8(instruction & 255)
	default:
		log.Fatal(fmt.Sprintf("%x", ((instruction & 3040) >> 8)) + " REGISTER NOT RECOGNIZED")
		
	}
	emu.PC += 2
}

func (emu *Chip8Emulator) run_7(instruction uint16) {
	fmt.Println(fmt.Sprintf("%x", instruction) + ": " + "ADD V" + fmt.Sprintf("%x", ((instruction & 3840) >> 8)) + ", " + fmt.Sprintf("%x", instruction & 255))
	switch ((instruction & 3840) >> 8) {
	case 0:
		emu.V0 += uint8(instruction & 255)
	case 1:
		emu.V1 += uint8(instruction & 255)
	case 2:
		emu.V2 += uint8(instruction & 255)
	case 3:
		emu.V3 += uint8(instruction & 255)
	case 4:
		emu.V4 += uint8(instruction & 255)
	case 5:
		emu.V5 += uint8(instruction & 255)
	case 6:
		emu.V6 += uint8(instruction & 255)
	case 7:
		emu.V7 += uint8(instruction & 255)
	case 8:
		emu.V8 += uint8(instruction & 255)
	case 9:
		emu.V9 += uint8(instruction & 255)
	case 10:
		emu.VA += uint8(instruction & 255)
	case 11:
		emu.VB += uint8(instruction & 255)
	case 12:
		emu.VC += uint8(instruction & 255)
	case 13:
		emu.VD += uint8(instruction & 255)
	case 14:
		emu.VE += uint8(instruction & 255)
	case 15:
		emu.VF += uint8(instruction & 255)
	default:
		log.Fatal(fmt.Sprintf("%x", ((instruction & 3040) >> 8)) + " REGISTER NOT RECOGNIZED")
		
	}
	emu.PC += 2
}

func (emu *Chip8Emulator) run_A(instruction uint16) {
	fmt.Println(fmt.Sprintf("%x", instruction) + ": " + "LD I,  " + fmt.Sprintf("%x", instruction & 4095))
	emu.I = instruction & 4095
	emu.PC += 2
}

func (emu *Chip8Emulator) run_D(instruction uint16) {
	fmt.Println(fmt.Sprintf("%x", instruction) + ": " + "DRW V" + fmt.Sprintf("%x", ((instruction & 3840) >> 8)) + ", V" + fmt.Sprintf("%x", ((instruction & 240) >> 4)) + ", " + fmt.Sprintf("%x", instruction & 15))
	x := int(((instruction & 3840) >> 8))
	y := int(((instruction & 240) >> 4))
	n := int(instruction & 15)
	i := 0
	j := emu.I
	for i < n {
		mask := uint8(128)
		bit_shift := 7
		for bit_shift != -1 {
			pixel := ((emu.RAM[j] & mask) >> bit_shift)
			if pixel == 0 && emu.Display[y][x] == ' ' {
				emu.Display[y][x] = ' '
			} else if pixel == 1 && emu.Display[y][x] == ' ' {
				emu.Display[y][x] = '*'
			} else if pixel == 0 && emu.Display[y][x] == '*' {
				emu.Display[y][x] = '*'
			} else {
				emu.Display[y][x] = ' '
				emu.VF = 1
			}
			mask = mask/2
			bit_shift--
			x++
			if x > 63 {
				x = 0
			}
		}
		i++
		j++
	}
	emu.PC += 2
}

func (emu *Chip8Emulator) clear_screen() {
	i := 0
	for i < len(emu.Display) {
		j := 0
		for j < len(emu.Display[i]) {
			emu.Display[i][j] = ' ' 
			j++
		}
		i++
	}
}

func (emu *Chip8Emulator) draw_screen() {
	i := 0
	for i < len(emu.Display) {
		j := 0
		for j < len(emu.Display[i]) {
			fmt.Print(string(emu.Display[i][j]))
			j++
		}
		fmt.Println()
		i++
	}
}








const (
	width = 64
	height = 32
)

func main() {
	file, err := ioutil.ReadFile("puzzle.ch8")
	if err != nil {
		panic(err)
	}

	emu := Chip8Emulator{PC: 512}
	emu.start(file)

	/*
	pc := 512
	for (pc-512) < len(file) {
		RAM[pc] = file[pc-512]
		pc++
	}
	running := true
	pc = 512
	for running {
		var next int
		instruction := (uint16(RAM[pc]) << 8) | uint16(RAM[pc+1])
		executeInstruction(instruction)
		pc += 2
		_, err := fmt.Scan(&next)
		if err != nil {
			panic(err)
		}
	}
	*/
	/*
	ic := 512
	i = 0
	
	for i < len(byteArray) {
		upperByte := fmt.Sprintf("%x", byteArray[i])
		lowerByte := fmt.Sprintf("%x", byteArray[i+1])
		if len(upperByte) == 1 {
			upperByte = "0" + upperByte
		}
		if len(lowerByte) == 1 {
			lowerByte = "0" + lowerByte
		}
		instruction := upperByte + lowerByte
		fmt.Println(strconv.FormatInt(int64(ic), 10) + ": " + instruction)
		ic+=2
		i += 2
	}
	*/
	/*
	runtime.LockOSThread()

	window := initGlfw()
	defer glfw.Terminate()

	program := initOpenGL()
	
	for !window.ShouldClose() {
		draw(window, program)
	}
	*/
}

func executeInstruction(instruction uint16) {
	opCode := getOpcode(instruction)
	switch opCode {
	case 0:
		run0(instruction)
	case 1:
		run1(instruction)
	case 2:
		log.Fatal(fmt.Sprintf("%x", opCode) + " NOT IMPLEMENTED")
	case 3:
		log.Fatal(fmt.Sprintf("%x", opCode) + " NOT IMPLEMENTED")
	case 4:
		log.Fatal(fmt.Sprintf("%x", opCode) + " NOT IMPLEMENTED")
	case 5:
		log.Fatal(fmt.Sprintf("%x", opCode) + " NOT IMPLEMENTED")
	case 6:
		run6(instruction)
	case 7:
		log.Fatal(fmt.Sprintf("%x", opCode) + " NOT IMPLEMENTED")
	case 8:
		log.Fatal(fmt.Sprintf("%x", opCode) + " NOT IMPLEMENTED")
	case 9:
		log.Fatal(fmt.Sprintf("%x", opCode) + " NOT IMPLEMENTED")
	case 10:
		run10(instruction)
	case 11:
		log.Fatal(fmt.Sprintf("%x", opCode) + " NOT IMPLEMENTED")
	case 12:
		log.Fatal(fmt.Sprintf("%x", opCode) + " NOT IMPLEMENTED")
	case 13:
		log.Fatal(fmt.Sprintf("%x", opCode) + " NOT IMPLEMENTED")
	case 14:
		log.Fatal(fmt.Sprintf("%x", opCode) + " NOT IMPLEMENTED")
	case 15:
		log.Fatal(fmt.Sprintf("%x", opCode) + " NOT IMPLEMENTED")
		
	default:
		log.Fatal(fmt.Sprintf("%x", opCode) + " NOT RECOGNIZED")
	}	
}

func run0(instruction uint16) {
	check := instruction & 255
	switch check {
	case 224:
		fmt.Println(fmt.Sprintf("%x", instruction) + ": " + "CLS")
	case 238:
		fmt.Println(fmt.Sprintf("%x", instruction) + ": " + "RET")
	default:
		fmt.Println(fmt.Sprintf("%x", instruction) + ": " + "SYS " + fmt.Sprintf("%x", instruction & 4095)) 
	}
}

func run1(instruction uint16) {
	fmt.Println(fmt.Sprintf("%x", instruction) + ": " + "JP " + fmt.Sprintf("%x", instruction & 4095))
}

func run6(instruction uint16) {
	fmt.Println(fmt.Sprintf("%x", instruction) + ": " + "LD V" + fmt.Sprintf("%x", ((instruction & 3040) >> 8)) + ", " + fmt.Sprintf("%x", instruction & 255))
}

func run10(instruction uint16) {
	fmt.Println(fmt.Sprintf("%x", instruction) + ": " + "LD I,  " + fmt.Sprintf("%x", instruction & 4095))
}


func getOpcode(instruction uint16) uint8 {
	return uint8(instruction >> 12)
}

func initGlfw() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, "Conway's Game of Life", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	return window
}

func initOpenGL() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	prog := gl.CreateProgram()
	gl.LinkProgram(prog)
	return prog
}

func draw(window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	glfw.PollEvents()
	window.SwapBuffers()
}

