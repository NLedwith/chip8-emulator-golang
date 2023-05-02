package main

import (
	"testing"
	"fmt"
)

/*
0nnn - SYS addr
Jump to a machine code routine at nnn.

This instruction is only used on the old computers on which Chip-8 was originally implemented. It is ignored by modern interpreters.
*/
func TestTable0nnn(t *testing.T) {
	var tests = []struct {
		emuStartState Chip8Emulator
		instruction uint16
		emuEndState Chip8Emulator
	}{
		{Chip8Emulator{PC: 512}, uint16(728), Chip8Emulator{PC: 514}},
		{Chip8Emulator{PC: 728}, uint16(0), Chip8Emulator{PC: 730}},
		{Chip8Emulator{PC: 4092}, uint16(0), Chip8Emulator{PC: 4094}},

	}

	for _, test := range tests {
		test.emuStartState.executeInstruction(test.instruction)
		if test.emuStartState.PC != test.emuEndState.PC {
			t.Errorf("Test Failed: %v intruction executed, %v expected PC, received: %v", fmt.Sprintf("%X", test.instruction), test.emuEndState.PC, test.emuStartState.PC)
		}
	}
}

/*
00E0 - CLS
Clear the display.
*/
func TestTable00E0(t *testing.T) {
}

/*
00EE - RET
Return from a subroutine

The interpretor sets the program counter to the address at the top of the stack, then subtracts 1 from the stack pointer. 
*/
func TestTable00EE(t *testing.T) {

}

/*
1nnn - JP addr
Jump to location nnn.

The interpreter sets the program counter to nnn.
*/
func TestTable1nnn(t *testing.T) {

}

/*
2nnn - CALL addr
Call subroutine at nnn.

The interpretor increments the stack pointer, then puts the current PC on the top of the stack. The PC is then set to nnn.
*/
func TestTable2nnn(t *testing.T) {

}

/*
3xkk - SE Vx, byte
Skip next instruction if Vx = kk.

The interpretor compares register Vx to kk, and if they are equal, increments the program counter by 2.
*/
func TestTable3xkk(t *testing.T) {

}

/*
4xkk - SNE Vx, byte
Skip next instruction if Vx != kk.

The interpretor compares register Vx to kk, and if they are not equal, increments the program counter by 2.
*/
func TestTable4xkk(t *testing.T) {
}

/*
5xy0 - SE Vx, Vy
Skip next instruction if Vx = Vy.

The interpretor compares register Vx to register Vy, and if they are equal, increments the program counter by 2.
*/
func TestTable5xy0(t *testing.T) {

}

/*
6xkk - LD Vx, byte
Set Vx = kk.

The interpretor puts the value kk into register Vx.
*/
func TestTable6xkk(t *testing.T) {

}

/*
7xkk - ADD Vx, byte
Set Vx = Vx + kk.

Adds the value kk to the value of register Vx, then stores the result in Vx.
*/
func TestTable7xkk(t *testing.T) {

}

/*
8xy0 - LD Vx, Vy
Set Vx = Vy.

Stores the value of register Vy in register Vx.
*/
func TestTable8xy0(t *testing.T) {
}

/*
8xy1 - OR Vx, Vy
Set Vx = Vx OR Vy

Performs a bitwise OR on the values of Vx and Vy, then stores the result in Vx. A bitwise OR compares the corresponding bits from two values, and if either bit is 1, then the same bit in the result is also 1. Otherwise, it is 0.
*/
func TestTable8xy1(t *testing.T) {
}

/*
8xy2 - AND Vx, Vy
Set Vx = Vx AND Vy.

Performs a bitwise AND on the values of Vx and Vy, then stores the result in Vx. A bitwise AND compares the corresponding bits from two values, and if both bits are 1, then the same bit in the result is also 1. Otherwise it is 0.
*/
func TestTable8xy2(t *testing.T) {
}

/*
8xy3 - XOR Vx, Vy
Set Vx = Vx XOR Vy.

Performs a bitwise exclusive OR on the values of Vx and Vy, then stores the result in Vx. An exclusive OR compares the corresponding bits from two values, and if the bits are not both the same, then the corresponding bit in the result is set to 1. Otherwise, it is 0.
*/
func TestTable8xy3(t *testing.T) {
}

/*
8xy4 - ADD Vx, Vy
Set Vx = Vx + Vy, set VF = carry.

The values of Vx and Vy are added together. If the result is greater than 8 bits (i.e., > 255) VF is set to 1, otherwise 0. Only the lowest 8 bits of the result are kept, and stored in Vx.
*/
func TestTable8xy4(t *testing.T) {
}

/*
8xy5 - SUB Vx, Vy
Set Vx = Vx - Vy, set VF = NOT borrow.

If Vx > Vy, then VF is set to 1, otherwise 0. Then Vy is subtracted from Vx, and the results are stored in Vx.
*/
func TestTable8xy5(t *testing.T) {
}

/*
8xy6 = SHR Vx {, Vy}
Set Vx = Vx SHR 1.

If the least-significant bit of Vx is 1, then VF is set to 1, otherwise 0. Then Vx is divided by 2.
*/
func TestTable8xy6(t *testing.T) {
}

/*
8xy7 SUBN Vx, Vy
Set Vx = Vy - Vx, set VF = NOT borrow.

If Vy > Vx, then VF is set to 1, otherwise 0. Then Vx is subtracted from Vy, and the results stored in Vx.
*/
func TestTable8xy7(t *testing.T) {

}

/*
8xyE - SHL Vx {, Vy}
Set Vx = Vx SHL 1.

If the most-significant bit of Vx is 1, then VF is set to 1, otherwise to 0. Then Vx is multiplied by 2.
*/
func TestTable8xyE(t *testing.T) {
}

/*
9xy0 - SNE Vx, Vy
Skip next instruction if Vx != Vy.

The values of Vx and Vy are compared, and if they are not equal, the program counter is increased by 2.
*/
func TestTable9xy0(t *testing.T) {
}

/*
Annn - LD I, addr
Set I = nnn.

The value of register I is set to nnn.
*/
func TestTableAnnn(t *testing.T) {
}

/*
Bnnn - JP V0, addr
Jump to location nnn + V0.

The program counter is set to nnn plus the value of V0.
*/
func TestTableBnnn(t *testing.T) {
}

/*
Cxkk - RND Vx, byte
Set Vx = random byte AND kk.

The interpreter generates a random number from 0 to 255, which is then ANDed with the value kk. The results are stored in Vx. See instruction 8xy2 for more information on AND.
*/
func TestTableCxkk(t *testing.T) {
}

/*
Dxyn - DRW Vx, Vy, nibble
Display n-byte sprite starting at memory location I at (Vx, Vy), set VF = collision.

The interpreter reads n bytes from memory, starting at the address stored in I. These bytes are then displayed as sprites on screen at coordinates (Vx, Vy). Sprites are XORed onto the existing screen. If this causes any pixels to be erased, VF is set to 1, otherwise it is set to 0. If the Sprite is positioned so part of it is outside the coordinates of the display, it wraps around to the opposite side of the screen. See instruction 8xy3 for more information on XOR, and section 2.4, Display, for more information on the Chip-8 screen and sprites.
*/
func TestTableDxyn(t *testing.T) {
}

/*
Ex9E - SKP Vx
Skip next instruction if key with the value of Vx is pressed.

Checks the keyboard, and if the key corresponding to the value of Vx is currently in the down position, PC is increased by 2.
*/
func TestTableEx9E(t *testing.T) {
}

/*
ExA1 - SKNP Vx
Skip next instruction if key with the value of Vx is not pressed.

Checks the keyboard, and if the key corresponding to the value of Vx is currently in the up position, PC is increased by 2.
*/
func TestTableExA1(t *testing.T) {
}

/*
Fx07 - LD Vx, DT
Set Vx = delay timer value.

The value of DT is placed into Vx.
*/
func TestTableFx07(t *testing.T) {
}

/*
Fx0A - LD Vx, K
Wait for a key press, store the value of the key in Vx.

All execution stops until a key is pressed, then the value of that key is stored in Vx.
*/
func TestTableFx0A(t *testing.T) {
}

/*
Fx15 - LD DT, Vx
Set delay timer = Vx.

DT is set equal to the value of Vx.
*/
func TestTableFx15(t *testing.T) {
}

/*
Fx18 - LD ST, Vx
Set sound timer = Vx.

ST is set equal to the value of Vx.
*/
func TestTableFx18(t *testing.T) {
}

/*
Fx1E - ADD I, Vx
Set I = I + Vx.

The values of I and Vx are added, and the results are stored in I.
*/
func TestTableFx1E(t *testing.T) {
}

/*
Fx29 - LD F, Vx
Set I = location of sprite for digit Vx.

The value of I is set to the location for the hexadecimal sprite corresponding to the value of Vx. See section 2.4, Display, for more information on the Chip-8 hexadecimal font.
*/
func TestTableFx29(t *testing.T) {
}

/*
Fx33 - LD B, Vx
Store BCD representation of Vx in memory locations I, I+1, and I+2.

The interpreter takes the decimal value of Vx, and places the hundreds digit in memory at location in I, the tens digit at location I+1, and the ones digit at location I+2.
*/
func TestTableFx33(t *testing.T) {

}

/*
Fx55 - LD [I], Vx
Store registers V0 through Vx in memory starting at location I.

The interpretor copies the values of registers V0 through Vx into memory, starting at the address in I.
*/
func TestTableFx55(t *testing.T) {
	
}

/*
Fx65 - LD Vx, [I]
Read registers V0 through Vx from memory starting at location I.

The interpretor reads values from memory starting at location I into registers V0 through Vx.
*/
func TestTableFx65(t *testing.T) {
}





