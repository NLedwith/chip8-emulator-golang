package main

import (
	"fmt"
	"log"
	"runtime"
	"math/rand"
	"strconv"
	"time"
)


type Chip8Emulator struct {
	V0      uint8
	V1      uint8
	V2      uint8
	V3      uint8
	V4      uint8
	V5      uint8
	V6      uint8
	V7      uint8
	V8      uint8
	V9      uint8
	VA      uint8
	VB      uint8
	VC      uint8
	VD      uint8
	VE      uint8
	VF      uint8
	I       uint16
	ST      uint8
	DT      uint8
	PC      uint16
	SP      uint8
	ClockSpeed int64
	Stack   [16]uint16
	RAM     [4096]uint8
	ScreenState [32][64]bool
	Screen display
}

func (emu *Chip8Emulator) initialize() {
	emu.loadSprites()
	emu.PC = uint16(512)
	emu.SP = uint8(16)
	emu.ClockSpeed = int64(2)
	emu.Screen.initialize()
}

func (emu *Chip8Emulator) start(program []uint8) {
	emu.loadProgram(program)
	
	runtime.LockOSThread()
	emu.Screen.start()

	st := time.Now().UnixMilli()
	ct := time.Now().UnixMilli()
	
	for !emu.Screen.window.ShouldClose(){
		ct = time.Now().UnixMilli()
		if ct - st >= emu.ClockSpeed { 
			instruction := (uint16(emu.RAM[emu.PC]) << 8) | uint16(emu.RAM[emu.PC+1])
			emu.executeInstruction(instruction)
			emu.Screen.draw(emu.ScreenState, emu.Screen.cells, emu.Screen.window, emu.Screen.program)
			if emu.DT != 0 {
				emu.DT -= 1
			}
			st = ct
		}	
	}
}

func (emu *Chip8Emulator) loadProgram(program []uint8) {
	i := int(emu.PC)
	for (i - 512) < len(program) {
		emu.RAM[i] = program[i-512]
		i++
	}
}

/*
function that parses  a 16 bit instruction and executes the corresponding command
*/
func (emu *Chip8Emulator) executeInstruction(instr uint16) {

	// Parse opcode, lower byte, and lower nibble out of instruction
	var op, lb, ln uint8 = uint8(instr >> 12), uint8(instr & 255), uint8(instr & 15)
	
 	fmt.Println(instr, op, lb, ln)	
	switch op {
	case 0:
		switch lb {
		case 224:
			emu.execute00E0()
		case 238:
			emu.execute00EE()
		default:
			emu.execute0nnn()
		}
	case 1:
		nnn := instr & 4095
		emu.execute1nnn(nnn)
	case 2:
		nnn := instr & 4095
		emu.execute2nnn(nnn)
	case 3: 
		x := uint8((instr & 3840) >> 8)
		emu.execute3xkk(x, lb)	
	case 4:
		x := uint8((instr & 3840) >> 8)
		emu.execute4xkk(x, lb)
	case 5:
		x := uint8((instr & 3840) >> 8)
		y := uint8((instr & 240) >> 4)
		emu.execute5xy0(x, y)
	case 6: 
		x := uint8((instr & 3840) >> 8)
		emu.execute6xkk(x, lb)
	case 7:
		x := uint8((instr & 3840) >> 8)
		emu.execute7xkk(x, lb)
	case 8: 
		x := uint8((instr & 3840) >> 8)
		y := uint8((instr & 240) >> 4)
		switch ln {
			case 0:
				emu.execute8xy0(x, y)
			case 1:
				emu.execute8xy1(x, y)
			case 2:
				emu.execute8xy2(x, y)
			case 3:
				emu.execute8xy3(x, y)
			case 4: 
				emu.execute8xy4(x, y)
			case 5:
				emu.execute8xy5(x, y)
			case 6:	
				emu.execute8xy6(x, y)
			case 7:
				emu.execute8xy7(x, y)
			
			case 14:
				emu.execute8xyE(x, y)
			default:
				log.Fatal(instr, "INSTRUCTION NOT RECOGNIZED")
		}
	
	case 9:
		x := uint8((instr & 3840) >> 8)
		y := uint8((instr & 240) >> 4)
		emu.execute9xy0(x, y)
	case 10:
		nnn := instr & 4095
		emu.executeAnnn(nnn)
	case 11: 
		nnn := instr & 4095
		emu.executeBnnn(nnn)
	case 12:
		x := uint8((instr & 3840) >> 8)
		emu.executeCxkk(x, lb)
	case 13:
		x := uint8((instr & 3840) >> 8)
		y := uint8((instr & 240) >> 4)
		emu.executeDxyn(x, y, ln)
	case 14: 
		x := uint8((instr & 3840) >> 8)
		switch lb {
		case 158:
			emu.executeEx9E(x)
		case 161:
			emu.executeExA1(x)
		default:
			log.Fatal(instr, "INSTRUCTION NOT RECOGNIZED")
		}
	case 15:
		x := uint8((instr & 3840) >> 8)
		switch lb {
		case 7:
			emu.executeFx07(x)
		case 10:
			emu.executeFx0A(x)
		case 21:
			emu.executeFx15(x)
		case 24:
			emu.executeFx18(x)
		case 30: 
			emu.executeFx1E(x)
		case 41:
			emu.executeFx29(x)
		case 51:
			emu.executeFx33(x)
		case 85:
			emu.executeFx55(x)
		case 101:
			emu.executeFx65(x)
		default:
			log.Fatal(instr, "INSTRUCTION NOT RECOGNIZED")
		}
	default:
		log.Fatal(instr, "INSTRUCTION NOT RECOGNIZED")
	}
}

/*
0nnn - SYS addr
Jump to a machine code routine at nnn.

This instruction is only used on the old computers on which Chip-8 was originally implemented. It is ignored by modern interpretors.
*/
func (emu *Chip8Emulator) execute0nnn() {
	emu.PC += 2
}

/*
00E0 - CLS
Clear the display.
*/
func (emu *Chip8Emulator) execute00E0() {
	emu.Screen.clear()
	emu.PC += 2
}

/*
00EE - RET
Return from a subroutine.

The interpretor sets the program counter to the address at the top of the stack, then subtracts 1 from the stack pointer.
*/
func (emu *Chip8Emulator) execute00EE() {
	emu.PC = emu.Stack[emu.SP]
	emu.SP += 1
}

/*
1nnn - JP addr
Jump to location nnn.

The interpretor sets the program counter to nnn.
*/
func (emu *Chip8Emulator) execute1nnn(nnn uint16) {
	emu.PC = nnn
}

/*
2nnn - CALL addr
Call subroutine at nnn.

The interpreter increments the stack pointer, then puts the current PC on the top of the stack. The PC is then set to nnn.
*/
func (emu *Chip8Emulator) execute2nnn(nnn uint16) {
	// CHECK FOR STACK OVERFLOW AND THROW ERROR HERE

	emu.SP--
	emu.Stack[emu.SP] = emu.PC + 2
	emu.PC = nnn
}

/*
3xkk - SE Vx, byte
Skip next instruction if Vx = kk.

The interpretor compares register Vx to kk, and if they are equal, increments the program counter by 2.
*/
func (emu *Chip8Emulator) execute3xkk(x uint8, kk uint8) {
	if *emu.intToRegister(x) == kk {
		emu.PC += 4
	} else {
		emu.PC += 2
	}
}

/*
4xkk - SNE Vx, byte
Skip next instruction if Vx != kk.

The interpreter compares register Vx to kk, and if they are not equal, increments the program counter by 2.
*/
func (emu *Chip8Emulator) execute4xkk(x uint8, kk uint8) {
	if  *emu.intToRegister(x) != kk {
		emu.PC += 4
	} else {
		emu.PC += 2
	}
}

/*
5xy0 SE Vx, Vy
Skip next instruction if Vx = Vy.

The interpreter compares register Vx to register Vy, and if they are equal, increments the program counter by 2
*/
func (emu *Chip8Emulator) execute5xy0(x uint8, y uint8) {
	if *emu.intToRegister(x) == *emu.intToRegister(y) {
		emu.PC += 4
	} else {
		emu.PC += 2
	}
}

/*
6xkk - LD Vx, byte
Set Vx = kk.

The interpretor puts the value kk into register Vx
*/
func (emu *Chip8Emulator) execute6xkk(x uint8, kk uint8) {
	*emu.intToRegister(x) = kk
	emu.PC += 2
}

/*
7xkk - ADD Vx, byte
Set Vx = Vx + kk.

Adds the value kk to the value of register Vx, then stores the result in Vx.
*/
func (emu *Chip8Emulator) execute7xkk(x uint8, kk uint8) {
	*emu.intToRegister(x) += kk
	emu.PC += 2
}

/*
8xy0 - LD Vx, Vy
Set Vx = Vy.

Stores the value of register Vy in register Vx.
*/
func (emu *Chip8Emulator) execute8xy0(x uint8, y uint8) {
	*emu.intToRegister(x) = *emu.intToRegister(y)
	emu.PC += 2
}

/*
8xy1 - OR Vx, Vy
Set Vx = Vx OR Vy

Performs a bitwise OR on the values of Vx and Vy, then stores the result in Vx. A bitwise OR compares the corresponding bits from two values, and if either bit is 1, then the same bit in the result is also 1. Otherwise it is 0.
*/
func (emu *Chip8Emulator) execute8xy1(x uint8, y uint8) {
	*emu.intToRegister(x) |= *emu.intToRegister(y)
	emu.PC += 2
}

/*
8xy2 - AND Vx, Vy
Set Vx = Vx AND Vy.

Performs a bitwise AND on the values of Vx and Vy, then stores the result in Vx. A bitwise AND compares the corresponding bits from two values, and if both bits are 1, then the same bit in the result is also 1. Otherwise it is 0.
*/
func (emu *Chip8Emulator) execute8xy2(x uint8, y uint8) {
	*emu.intToRegister(x) &= *emu.intToRegister(y)
	emu.PC += 2
}

/*
8xy3 - XOR Vx, Vy
Set Vx = Vx XOR Vy

Performs a bitwise exclusive OR on the values of Vx and Vy, then stores the result in Vx. An exclusive OR compares the corresponding bits from two values, and if the bits are not both the same, then the corresponding bit in the result is set to 1. Otherwise, it is 0.
*/
func (emu *Chip8Emulator) execute8xy3(x uint8, y uint8) {
	*emu.intToRegister(x) ^= *emu.intToRegister(y)
	emu.PC += 2
}

/*
8xy4 - ADD Vx, Vy
Set Vx = Vx + Vy, set VF = carry.

The values of Vx and Vy are added together. If the result is greater than 8 bits (i.e., > 255) VF is set to 1, otherwise 0. Only the lowest 8 bits of the result are kept, and stored in Vx.
*/
func (emu *Chip8Emulator) execute8xy4(x uint8, y uint8) {
	if (*emu.intToRegister(x) & 128) == 128 && (*emu.intToRegister(y) & 128) == 128 {
		emu.VF = 1
	} else {
		emu.VF = 0
	}
	*emu.intToRegister(x) += *emu.intToRegister(y)
	emu.PC += 2
}

/*
8xy5 - SUB Vx, Vy
Set Vx = Vx - Vy, set VF = NOT borrow.

If Vx > Vy, then VF is set to 1, otherwise 0. Then Vy is subtracted from Vx, and the results are stored in Vx.
*/
func (emu *Chip8Emulator) execute8xy5(x uint8, y uint8) {
	if *emu.intToRegister(x) > *emu.intToRegister(y) {
		emu.VF = 1
	} else {
		emu.VF = 0
	}
	*emu.intToRegister(x) -= *emu.intToRegister(y)
	emu.PC += 2
}

/*
8xy6 - SHR Vx {, Vy}
Set Vx = Vx SHR 1.

If the least-significant bit of Vx is 1, the VF is set to 1, otherwise 0. Then Vx is divided by 2.
*/
func (emu *Chip8Emulator) execute8xy6(x uint8, y uint8) {
	if (*emu.intToRegister(x) & 1) == 1{
		emu.VF = 1
	} else {
		emu.VF = 0
	}
	*emu.intToRegister(x) = *emu.intToRegister(x) >> 1
	emu.PC += 2
}

/*
8xy7 SUBN Vx, Vy
Set Vx = Vy - Vx, set VF = NOT borrow.

If Vy > Vx, then VF is set to 1, otherwise 0. Then Vx is subtracted from Vy, and the results stored in Vx.
*/
func (emu *Chip8Emulator) execute8xy7(x uint8, y uint8) {
	if *emu.intToRegister(y) > *emu.intToRegister(x) {
		emu.VF = 1
	} else {
		emu.VF = 0
	}
	*emu.intToRegister(x) = *emu.intToRegister(y) - *emu.intToRegister(x)
	emu.PC += 2
}

/*
8xyE - SHL Vx {, Vy}
Set Vx = Vx SHL 1.

If the most-significant bit to Vx is 1, the Vf is set to 1, otherwise to 0. the Vx is multiplied by 2.
*/
func (emu *Chip8Emulator) execute8xyE(x uint8, y uint8) {
	if (*emu.intToRegister(x) & 128) == 128 {
		emu.VF = 1
	} else {
		emu.VF = 0
	}
	*emu.intToRegister(x) = *emu.intToRegister(x) << 1
	emu.PC += 2
}

/*
9xy0 - SNE Vx, Vy
Skip next instruction if Vx != Vy.

The values of Vx and Vy are compared, and if they are not equal, the program counter is increased by 2.
*/
func (emu *Chip8Emulator) execute9xy0(x uint8, y uint8) {
	if *emu.intToRegister(x) != *emu.intToRegister(y) {
		emu.PC += 4
	} else {
		emu.PC += 2
	}
}

/*
Annn - LD I, addr
Set I = nnn.

The value of register I is set to nnn.
*/
func (emu *Chip8Emulator) executeAnnn(nnn uint16) {
	emu.I = nnn
	emu.PC += 2
}

/*
Bnnn = JP V0, addr
Jump to location nnn + V0

The program counter is set to nnn plus the value of V0.
*/
func (emu *Chip8Emulator) executeBnnn(nnn uint16) {
	emu.PC = nnn + uint16(emu.V0)
}

/*
Cxkk - RND Vx, byte
Set Vx = random byte AND kk.

The interpreter generates a random number from 0 to 255, which is then ANDed with the value kk. The results are stored in Vx. See instruction 8xy2 for more information on AND.
*/
func (emu *Chip8Emulator) executeCxkk(x uint8, kk uint8) {
	*emu.intToRegister(x) = uint8(rand.Intn(255)) & kk
	emu.PC += 2
}

/*
Dxyn - DRW Vx, Vy, nibble
Display n-byte sprite starting at memory location I at (Vx, Vy), set VF = collision.

The interpreter reads n bytes from memory, starting at the address stored in I. These bytes are then displayed as sprites on screen at coordinates (Vx, Vy). Sprites are XORed onto the existing screen. If this causes any pixels to be erased, VF is set to 1, otherwise it is set to 0. If the Sprite is positioned so part of it is outside the coordinates of the display, it wraps around to the opposite side of the screen. See instruction 8xy3 for more information on XOR, and section 2.4, Display, for more information on the Chip-8 screen and sprites.
*/
func (emu *Chip8Emulator) executeDxyn(x uint8, y uint8, n uint8) {
	vX := *emu.intToRegister(x)
	vY := *emu.intToRegister(y)
	i := uint8(0)
	j := emu.I
	f := false
	fmt.Println("Drawing", n, "bytes starting at", x, 31-y)
	for i < n {
		if vY >= 32 {
			i++
			j++
		} else {
			c := emu.Screen.updateScreenState(emu.RAM[j], vX, vY)
			f = f || c
			vY++
			i++
			j++
		}
	}
	if f {
		emu.VF = uint8(1)
	} else {
		emu.VF = uint8(0)
	}
	emu.PC += 2
}

/*
Ex9E - SKP Vx
Skip next instruction if key with the value of Vx is pressed.

Checks the keyboard, and if the key corresponding to the value of Vx is currently in the down position, PC is increased by 2.
*/
func (emu *Chip8Emulator) executeEx9E(x uint8) {
	if emu.Screen.checkKeyStatus(*emu.intToRegister(x))  {
		emu.PC += 4
	} else {
		emu.PC += 2
	}
}

/*
ExA1 - SKNP Vx
Skip next instruction if the key with the value of Vx is not pressed.

Checks the keyboard, and if the key corresponding to the value of Vx is currently in the up position, PC is increased by 2.
*/
func (emu *Chip8Emulator) executeExA1(x uint8) {
	if !emu.Screen.checkKeyStatus(*emu.intToRegister(x)) {
		emu.PC += 4
	} else {
		emu.PC += 2
	}
}

/*
Fx07 - LD Vx, DT
Set Vx = delay timer value.

The value of DT is placed into Vx.
*/
func (emu *Chip8Emulator) executeFx07(x uint8) {
	*emu.intToRegister(x) = emu.DT
	emu.PC += 2
}

/*
Fx0A - LD Vx, K
Wait for a key press, store the value of the key in Vx.
*/
func (emu *Chip8Emulator) executeFx0A(x uint8) {
	*emu.intToRegister(x) = emu.Screen.checkAnyKeyDown()
	emu.PC += 2
}

/*
Fx15 - LD DT, Vx
Set delay timer = Vx.

DT is set equal to the value of Vx.
*/
func (emu *Chip8Emulator) executeFx15(x uint8) {
	emu.DT = *emu.intToRegister(x)
	emu.PC += 2
}

/*
Fx18 - LD ST, Vx
Set sound timer = Vx.

ST is set equal to the value of Vx.
*/
func (emu *Chip8Emulator) executeFx18(x uint8) {
	emu.ST = *emu.intToRegister(x)
	emu.PC += 2
}

/*
Fx1E - ADD I, Vx
Set I = I + Vx

The values of I and Vx are added, and the results are stored in I.
*/
func (emu *Chip8Emulator) executeFx1E(x uint8) {
	emu.I += uint16(*emu.intToRegister(x))
	emu.PC += 2
}

/*
Fx29 - LD F, Vx
Set I = location of sprite for digit Vx.

The value of I is set to the location for the hexadecimal sprite corresponding to the value of Vx. See section 2.4, Display, for more information on the Chip-8 hexadecimal font.
*/
func (emu *Chip8Emulator) executeFx29(x uint8) {
	emu.I = uint16(*emu.intToRegister(x)) * 5
	emu.PC += 2
}

/*
Fx33 - LD B, Vx
Store BCD representation of Vx in memory locations I, I+1, and I+2.

The interpreter takes the decimal value of Vx, and places the hundreds digit in memory at location in I, the tens digit at location I + 1, and the ones digit at location I+2.
*/
func (emu *Chip8Emulator) executeFx33(x uint8) {
	vX := *emu.intToRegister(x)
	i := 2
	for i >= 0 {
		emu.RAM[emu.I + uint16(i)] = vX % 10
		vX = vX / 10
		i--
	}
	emu.PC += 2
}

/*
Fx55 - LD [I], Vx
Store registers V0 through Vx in memory starting at location I.

The interpreter copies the values of registers V0 through Vx into memory, starting at the address in I.
*/
func (emu *Chip8Emulator) executeFx55(x uint8) {
	i := uint16(0)
	n := uint16(x)
	buf := emu.I
	for i <= n {
		reg := *emu.intToRegister(uint8(i))
		emu.RAM[buf] = reg
		buf++
		i++
	}
	emu.PC += 2
}

/*
Fx65 - LD Vx, [I]
Read registers V0 through Vx from memory starting at location I.

The interpreter reads values from memory starting at location I into registers V0 through Vx.
*/
func (emu *Chip8Emulator) executeFx65(x uint8) {
	i := uint16(0)
	n := uint16(x)
	buf := emu.I
	for i <= n {
		*emu.intToRegister(uint8(i)) = uint8(emu.RAM[buf])
		buf++
		i++
	}
	emu.PC += 2
}


func (emu *Chip8Emulator) intToRegister(val uint8) *uint8 {
	var register *uint8
	switch val {
	case 0:
		register = &emu.V0
	case 1:
		register = &emu.V1
	case 2:
		register = &emu.V2
	case 3:
		register = &emu.V3
	case 4:
		register = &emu.V4
	case 5:
		register = &emu.V5
	case 6:
		register = &emu.V6
	case 7:
		register = &emu.V7
	case 8:
		register = &emu.V8
	case 9:
		register = &emu.V9
	case 10:
		register = &emu.VA
	case 11:
		register = &emu.VB
	case 12:
		register = &emu.VC
	case 13:
		register = &emu.VD
	case 14:
		register = &emu.VE
	case 15:
		register = &emu.VF
	default:
		log.Fatal(fmt.Sprintf("%x", (val)) + " REGISTER NOT RECOGNIZED")
	}
	return register
}

func (emu *Chip8Emulator) debug(instruction uint16, flag int) {
	fmt.Println("INSTRUCTION: " + fmt.Sprintf("%x", instruction))
	fmt.Println("V0: " + strconv.FormatInt(int64(emu.V0), 10))
	fmt.Println("V1: " + strconv.FormatInt(int64(emu.V1), 10))
	fmt.Println("V2: " + strconv.FormatInt(int64(emu.V2), 10))
	fmt.Println("V3: " + strconv.FormatInt(int64(emu.V3), 10))
	fmt.Println("V4: " + strconv.FormatInt(int64(emu.V4), 10))
	fmt.Println("V5: " + strconv.FormatInt(int64(emu.V5), 10))
	fmt.Println("V6: " + strconv.FormatInt(int64(emu.V6), 10))
	fmt.Println("V7: " + strconv.FormatInt(int64(emu.V7), 10))
	fmt.Println("V8: " + strconv.FormatInt(int64(emu.V8), 10))
	fmt.Println("V9: " + strconv.FormatInt(int64(emu.V9), 10))
	fmt.Println("VA: " + strconv.FormatInt(int64(emu.VA), 10))
	fmt.Println("VB: " + strconv.FormatInt(int64(emu.VB), 10))
	fmt.Println("VC: " + strconv.FormatInt(int64(emu.VC), 10))
	fmt.Println("VD: " + strconv.FormatInt(int64(emu.VD), 10))
	fmt.Println("VE: " + strconv.FormatInt(int64(emu.VE), 10))
	fmt.Println("VF: " + strconv.FormatInt(int64(emu.VF), 10))
	fmt.Println("I: " + strconv.FormatInt(int64(emu.I), 10))
	fmt.Println("ST: " + strconv.FormatInt(int64(emu.ST), 10))
	fmt.Println("DT: " + strconv.FormatInt(int64(emu.DT), 10))
	fmt.Println("PC: " + strconv.FormatInt(int64(emu.PC), 10))
	fmt.Println("SP: " + strconv.FormatInt(int64(emu.SP), 10))
	fmt.Println("Stack:", emu.Stack)
	if flag == 1 {
		log.Fatal()
	}
}


func (emu *Chip8Emulator) loadSprites() {
	// "0" Sprite
	emu.RAM[0] = 240
	emu.RAM[1] = 144
	emu.RAM[2] = 144
	emu.RAM[3] = 144
	emu.RAM[4] = 240

	// "1" Sprite
	emu.RAM[5] = 32
	emu.RAM[6] = 96
	emu.RAM[7] = 32
	emu.RAM[8] = 32
	emu.RAM[9] = 112

	// "2" Sprite
	emu.RAM[10] = 240
	emu.RAM[11] = 16
	emu.RAM[12] = 240
	emu.RAM[13] = 128
	emu.RAM[14] = 240

	// "3" Sprite
	emu.RAM[15] = 240
	emu.RAM[16] = 16
	emu.RAM[17] = 240
	emu.RAM[18] = 16
	emu.RAM[19] = 240

	// "4" Sprite
	emu.RAM[20] = 144
	emu.RAM[21] = 144
	emu.RAM[22] = 240
	emu.RAM[23] = 16
	emu.RAM[24] = 16

	// "5" Sprite
	emu.RAM[25] = 240
	emu.RAM[26] = 128
	emu.RAM[27] = 240
	emu.RAM[28] = 16
	emu.RAM[29] = 240

	// "6" Sprite
	emu.RAM[30] = 240
	emu.RAM[31] = 128
	emu.RAM[32] = 240
	emu.RAM[33] = 144
	emu.RAM[34] = 240

	// "7" Sprite
	emu.RAM[35] = 240
	emu.RAM[36] = 16
	emu.RAM[37] = 32
	emu.RAM[38] = 64
	emu.RAM[39] = 64

	// "8" Sprite
	emu.RAM[40] = 240
	emu.RAM[41] = 144
	emu.RAM[42] = 240
	emu.RAM[43] = 144
	emu.RAM[44] = 240

	// "9" Sprite
	emu.RAM[45] = 240
	emu.RAM[46] = 144
	emu.RAM[47] = 240
	emu.RAM[48] = 16
	emu.RAM[49] = 240

	// "A" Sprite
	emu.RAM[50] = 240
	emu.RAM[51] = 144
	emu.RAM[52] = 240
	emu.RAM[53] = 144
	emu.RAM[54] = 144

	// "B" Sprite
	emu.RAM[55] = 224
	emu.RAM[56] = 144
	emu.RAM[57] = 224
	emu.RAM[58] = 144
	emu.RAM[59] = 224

	// "C" Sprite
	emu.RAM[60] = 240
	emu.RAM[61] = 128
	emu.RAM[62] = 128
	emu.RAM[63] = 128
	emu.RAM[64] = 240

	// "D" Sprite
	emu.RAM[65] = 224
	emu.RAM[66] = 144
	emu.RAM[67] = 144
	emu.RAM[68] = 144
	emu.RAM[69] = 224

	// "E" Sprite
	emu.RAM[70] = 240
	emu.RAM[71] = 128
	emu.RAM[72] = 240
	emu.RAM[73] = 128
	emu.RAM[74] = 240

	// "F" Sprite
	emu.RAM[75] = 240
	emu.RAM[76] = 128
	emu.RAM[77] = 240
	emu.RAM[78] = 128
	emu.RAM[79] = 128
}
