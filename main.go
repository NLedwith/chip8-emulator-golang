package main

import (
	"fmt"
	"log"

	//"runtime"
	"io/ioutil"
	"math/rand"
	"strconv"
	"github.com/mattn/go-tty"
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
	Stack   [16]uint16
	RAM     [4096]uint8
	Display [32][64]rune
}

func (emu *Chip8Emulator) start(program []uint8) {
	emu.load_program(program)
	emu.load_sprites()
	running := true
	for running {
		instruction := (uint16(emu.RAM[emu.PC]) << 8) | uint16(emu.RAM[emu.PC+1])
		emu.execute_instruction(instruction)
		emu.draw_screen()
		/*
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
			var w1 string
			_, err := fmt.Scanln(&w1)
			if err != nil {
				log.Fatal(err)
			}
			if w1 == "draw" {
				emu.draw_screen()
			}
		*/

	}
}
func (emu *Chip8Emulator) load_sprites() {
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

func (emu *Chip8Emulator) load_program(program []uint8) {
	i := int(emu.PC)
	for (i - 512) < len(program) {
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
		emu.run_2(instruction)
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
		emu.run_8(instruction)
	case 9:
		log.Fatal(fmt.Sprintf("%x", op) + " NOT IMPLEMENTED")
	case 10:
		emu.run_A(instruction)
	case 11:
		log.Fatal(fmt.Sprintf("%x", op) + " NOT IMPLEMENTED")
	case 12:
		emu.run_C(instruction)
	case 13:
		emu.run_D(instruction)
	case 14:
		log.Fatal(fmt.Sprintf("%x", op) + " NOT IMPLEMENTED")
	case 15:
		emu.run_F(instruction)

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
		emu.SP += 1
	default:
		fmt.Println(fmt.Sprintf("%x", instruction) + ": " + "SYS " + fmt.Sprintf("%x", instruction&4095))
		emu.PC = instruction & 4095
	}
}

func (emu *Chip8Emulator) run_1(instruction uint16) {
	fmt.Println(emu.PC)
	fmt.Println(fmt.Sprintf("%x", instruction) + ": " + "JP " + fmt.Sprintf("%x", (instruction&4095)))
	emu.PC = instruction & 4095
}

func (emu *Chip8Emulator) run_2(instruction uint16) {
	fmt.Println(fmt.Sprintf("%x", instruction) + ": " + "CALL " + fmt.Sprintf("%x", instruction&4095))
	emu.SP--
	emu.Stack[emu.SP] = emu.PC
	emu.PC = (instruction & 4095)
}

func (emu *Chip8Emulator) run_3(instruction uint16) {
	fmt.Println(fmt.Sprintf("%x", instruction) + ": " + "SE V" + fmt.Sprintf("%x", ((instruction&3840)>>8)) + ", " + fmt.Sprintf("%x", instruction&255))
	switch (instruction & 3840) >> 8 {
	case 0:
		if emu.V0 == uint8(instruction&255) {
			emu.PC += 4
		} else {
			emu.PC += 2
		}
	case 1:
		if emu.V1 == uint8(instruction&255) {
			emu.PC += 4
		} else {
			emu.PC += 2
		}
	case 2:
		if emu.V2 == uint8(instruction&255) {
			emu.PC += 4
		} else {
			emu.PC += 2
		}
	case 3:
		if emu.V3 == uint8(instruction&255) {
			emu.PC += 4
		} else {
			emu.PC += 2
		}
	case 4:
		if emu.V4 == uint8(instruction&255) {
			emu.PC += 4
		} else {
			emu.PC += 2
		}
	case 5:
		if emu.V5 == uint8(instruction&255) {
			emu.PC += 4
		} else {
			emu.PC += 2
		}
	case 6:
		if emu.V6 == uint8(instruction&255) {
			emu.PC += 4
		} else {
			emu.PC += 2
		}
	case 7:
		if emu.V7 == uint8(instruction&255) {
			emu.PC += 4
		} else {
			emu.PC += 2
		}
	case 8:
		if emu.V8 == uint8(instruction&255) {
			emu.PC += 4
		} else {
			emu.PC += 2
		}
	case 9:
		if emu.V9 == uint8(instruction&255) {
			emu.PC += 4
		} else {
			emu.PC += 2
		}
	case 10:
		if emu.VA == uint8(instruction&255) {
			emu.PC += 4
		} else {
			emu.PC += 2
		}
	case 11:
		if emu.VB == uint8(instruction&255) {
			emu.PC += 4
		} else {
			emu.PC += 2
		}
	case 12:
		if emu.VC == uint8(instruction&255) {
			emu.PC += 4
		} else {
			emu.PC += 2
		}
	case 13:
		if emu.VD == uint8(instruction&255) {
			emu.PC += 4
		} else {
			emu.PC += 2
		}
	case 14:
		if emu.VE == uint8(instruction&255) {
			emu.PC += 4
		} else {
			emu.PC += 2
		}
	case 15:
		if emu.VF == uint8(instruction&255) {
			emu.PC += 4
		} else {
			emu.PC += 2
		}
	default:
		log.Fatal(fmt.Sprintf("%x", ((instruction&3040)>>8)) + " REGISTER NOT RECOGNIZED")

	}
}

func (emu *Chip8Emulator) run_4(instruction uint16) {
	fmt.Println(fmt.Sprintf("%x", instruction) + ": " + "SE V" + fmt.Sprintf("%x", ((instruction&3840)>>8)) + ", " + fmt.Sprintf("%x", instruction&255))
	switch (instruction & 3840) >> 8 {
	case 0:
		if emu.V0 != uint8(instruction&255) {
			emu.PC += 4
		} else {
			emu.PC += 2
		}
	case 1:
		if emu.V1 != uint8(instruction&255) {
			emu.PC += 4
		} else {
			emu.PC += 2
		}
	case 2:
		if emu.V2 != uint8(instruction&255) {
			emu.PC += 4
		} else {
			emu.PC += 2
		}
	case 3:
		if emu.V3 != uint8(instruction&255) {
			emu.PC += 4
		} else {
			emu.PC += 2
		}
	case 4:
		if emu.V4 != uint8(instruction&255) {
			emu.PC += 4
		} else {
			emu.PC += 2
		}
	case 5:
		if emu.V5 != uint8(instruction&255) {
			emu.PC += 4
		} else {
			emu.PC += 2
		}
	case 6:
		if emu.V6 != uint8(instruction&255) {
			emu.PC += 4
		} else {
			emu.PC += 2
		}
	case 7:
		if emu.V7 != uint8(instruction&255) {
			emu.PC += 4
		} else {
			emu.PC += 2
		}
	case 8:
		if emu.V8 != uint8(instruction&255) {
			emu.PC += 4
		} else {
			emu.PC += 2
		}
	case 9:
		if emu.V9 != uint8(instruction&255) {
			emu.PC += 4
		} else {
			emu.PC += 2
		}
	case 10:
		if emu.VA != uint8(instruction&255) {
			emu.PC += 4
		} else {
			emu.PC += 2
		}
	case 11:
		if emu.VB != uint8(instruction&255) {
			emu.PC += 4
		} else {
			emu.PC += 2
		}
	case 12:
		if emu.VC != uint8(instruction&255) {
			emu.PC += 4
		} else {
			emu.PC += 2
		}
	case 13:
		if emu.VD != uint8(instruction&255) {
			emu.PC += 4
		} else {
			emu.PC += 2
		}
	case 14:
		if emu.VE != uint8(instruction&255) {
			emu.PC += 4
		} else {
			emu.PC += 2
		}
	case 15:
		if emu.VF != uint8(instruction&255) {
			emu.PC += 4
		} else {
			emu.PC += 2
		}
	default:
		log.Fatal(fmt.Sprintf("%x", ((instruction&3040)>>8)) + " REGISTER NOT RECOGNIZED")

	}
}

func (emu *Chip8Emulator) run_6(instruction uint16) {
	fmt.Println(fmt.Sprintf("%x", instruction) + ": " + "LD V" + fmt.Sprintf("%x", ((instruction&3840)>>8)) + ", " + fmt.Sprintf("%x", instruction&255))
	switch (instruction & 3840) >> 8 {
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
		log.Fatal(fmt.Sprintf("%x", ((instruction&3040)>>8)) + " REGISTER NOT RECOGNIZED")

	}
	emu.PC += 2
}

func (emu *Chip8Emulator) run_7(instruction uint16) {
	fmt.Println(fmt.Sprintf("%x", instruction) + ": " + "ADD V" + fmt.Sprintf("%x", ((instruction&3840)>>8)) + ", " + fmt.Sprintf("%x", instruction&255))
	switch (instruction & 3840) >> 8 {
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
		log.Fatal(fmt.Sprintf("%x", ((instruction&3840)>>8)) + " REGISTER NOT RECOGNIZED")

	}
	emu.PC += 2
}

func (emu *Chip8Emulator) run_8(instruction uint16) {
	fmt.Println(instruction)
	switch instruction & 15 {
	case 2:
		fmt.Println(fmt.Sprintf("%x", instruction) + ": " + "AND V" + fmt.Sprintf("%x", ((instruction&3840)>>8)) + ", V" + fmt.Sprintf("%x", ((instruction&240)>>4)))
		*emu.get_register(int(((instruction & 3840) >> 8))) = *emu.get_register(int(((instruction & 3840) >> 8))) & *emu.get_register(int(((instruction & 240) >> 4)))
	default:
		log.Fatal("NOT IMPLEMENTED")
	}
	emu.PC += 2
}

func (emu *Chip8Emulator) run_A(instruction uint16) {
	fmt.Println(fmt.Sprintf("%x", instruction) + ": " + "LD I,  " + fmt.Sprintf("%x", instruction&4095))
	emu.I = instruction & 4095
	emu.PC += 2
}

func (emu *Chip8Emulator) run_C(instruction uint16) {
	fmt.Println(fmt.Sprintf("%x", instruction) + ": " + "RND V" + fmt.Sprintf("%x", ((instruction&3840)>>8)) + ", " + fmt.Sprintf("%x", (instruction&255)))
	*emu.get_register(int(((instruction & 3840) >> 8))) = uint8(rand.Intn(255)) & uint8(instruction&255)
	emu.PC += 2
}

func (emu *Chip8Emulator) run_D(instruction uint16) {
	fmt.Println(fmt.Sprintf("%x", instruction) + ": " + "DRW V" + fmt.Sprintf("%x", ((instruction&3840)>>8)) + ", V" + fmt.Sprintf("%x", ((instruction&240)>>4)) + ", " + fmt.Sprintf("%x", instruction&15))
	x := *emu.get_register(int(((instruction & 3840) >> 8)))
	y := *emu.get_register(int(((instruction & 240) >> 4)))

	n := int(instruction & 15)
	i := 0
	j := emu.I
	for i < n {
		fmt.Print(emu.RAM[j])
		mask := uint8(128)
		bit_shift := 7
		cur_x := x
		for bit_shift != -1 {
			pixel := ((emu.RAM[j] & mask) >> bit_shift)
			if pixel == 0 && emu.Display[y][cur_x] == ' ' {
				emu.Display[y][cur_x] = ' '
			} else if pixel == 1 && emu.Display[y][cur_x] == ' ' {
				emu.Display[y][cur_x] = '*'
			} else if pixel == 0 && emu.Display[y][cur_x] == '*' {
				emu.Display[y][cur_x] = '*'
			} else {
				emu.Display[y][cur_x] = ' '
				emu.VF = 1
			}
			mask = mask / 2
			bit_shift--
			cur_x++
			if cur_x > 63 {
				cur_x = 0
			}
		}
		y++
		if y > 31 {
			y = 0
		}
		i++
		j++
	}
	fmt.Println()
	emu.PC += 2
}

func (emu *Chip8Emulator) run_F(instruction uint16) {
	fmt.Println(instruction)
	switch instruction & 255 {
	case 7:
		emu.debug(instruction, 0)
		*emu.get_register(int(((instruction & 3840) >> 8))) = emu.DT
		emu.debug(instruction, 1)
	
	case 10:
		fmt.Println(fmt.Sprintf("%x", instruction) + ": " + "LD V" + fmt.Sprintf("%x", (instruction&3840)>>8) + ", K")
		*emu.get_register(int(((instruction & 3840) >> 8))) = get_key_press()
	
	case 21:
		emu.debug(instruction, 0)
		reg := *emu.get_register(int(((instruction & 3840) >> 8)))
		emu.DT = reg
		emu.debug(instruction, 1)
	case 24:
		emu.debug(instruction, 0)
		reg := *emu.get_register(int(((instruction & 3840) >> 8)))
		emu.ST = reg
		emu.debug(instruction, 1)
	case 30:
		fmt.Println(fmt.Sprintf("%x", instruction) + ": " + "ADD I, V" + fmt.Sprintf("%x", (instruction&3840)>>8))
		reg := *emu.get_register(int(((instruction & 3840) >> 8)))
		emu.I += uint16(reg)
	case 41:
		fmt.Println(fmt.Sprintf("%x", instruction) + ": " + "LD F, V" + fmt.Sprintf("%x", (instruction&3840)>>8))
		reg := *emu.get_register(int(((instruction & 3840) >> 8)))
		emu.I = uint16(reg) * 5
	
	case 51:
		fmt.Println(fmt.Sprintf("%x", instruction) + ": " + "LD B, V" + fmt.Sprintf("%x", (instruction&3840)>>8))
		reg := *emu.get_register(int(((instruction & 3840) >> 8)))
		i := 2
		for i >= 0 {
			emu.RAM[emu.I + uint16(i)] = reg % 10
			reg = reg / 10
			i--
		}
	
	case 85:
		fmt.Println(fmt.Sprintf("%x", instruction) + ": " + "LD [I], V" + fmt.Sprintf("%x", (instruction&3840)>>8))
		i := uint16(0)
		n := ((instruction & 3840) >> 8)
		buf := emu.I
		for i <= n {
			reg := *emu.get_register(int(i))
			emu.RAM[buf] = reg
			i++
		}
	case 101:
		fmt.Println(fmt.Sprintf("%x", instruction) + ": " + "LD V" + fmt.Sprintf("%x", (instruction&3840)>>8) + ", [I]")
		i := uint16(0)
		n := ((instruction & 3840) >> 8)
		buf := emu.I
		for i <= n {
			*emu.get_register(int(i)) = emu.RAM[buf]
			buf++
			i++
		}
	default:
		log.Fatal("NOT IMPLEMENTED")
	}
	emu.PC += 2

}

func (emu *Chip8Emulator) get_register(val int) *uint8 {
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

const (
	width  = 64
	height = 32
)

func get_key_press() uint8{
	tty, err := tty.Open()
    if err != nil {
        log.Fatal(err)
    }
    defer tty.Close()

    for {
        r, err := tty.ReadRune()
        if err != nil {
            log.Fatal(err)
        }
		switch r {
			case 48:
				return 0
			case 49:
				return 1
			case 50:
				return 2
			case 51:
				return 3
			case 52:
				return 4
			case 53:
				return 5
			case 54:
				return 6
			case 55:
				return 7
			case 56:
				return 8
			case 57:
				return 9
			case 97:
				return 10
			case 98:
				return 11
			case 99:
				return 12
			case 100:
				return 13
			case 101:
				return 14
			case 102:
				return 15
		}
    }
	return 3
}

func main() {
	fmt.Println("Hello")
	file, err := ioutil.ReadFile("blitz.ch8")
	if err != nil {
		panic(err)
	}

	emu := Chip8Emulator{PC: 512, SP: 16}
	emu.start(file)

	/*
		byteArray := [4096]byte{}
		pc := 512
		for (pc-512) < len(file) {
			byteArray[pc] = file[pc-512]
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

		ic := 0
		i := 0

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
		fmt.Println(fmt.Sprintf("%x", instruction) + ": " + "SYS " + fmt.Sprintf("%x", instruction&4095))
	}
}

func run1(instruction uint16) {
	fmt.Println(fmt.Sprintf("%x", instruction) + ": " + "JP " + fmt.Sprintf("%x", instruction&4095))
}

func run6(instruction uint16) {
	fmt.Println(fmt.Sprintf("%x", instruction) + ": " + "LD V" + fmt.Sprintf("%x", ((instruction&3040)>>8)) + ", " + fmt.Sprintf("%x", instruction&255))
}

func run10(instruction uint16) {
	fmt.Println(fmt.Sprintf("%x", instruction) + ": " + "LD I,  " + fmt.Sprintf("%x", instruction&4095))
}

func getOpcode(instruction uint16) uint8 {
	return uint8(instruction >> 12)
}
