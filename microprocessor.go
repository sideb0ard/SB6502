package main

// The SB6502, a virtual clone of the MOS 6502

import (
	"fmt"

	"github.com/mgutz/ansi"
)

// Instruction Loop:
// The instruction at the address given by the Instruction Pointer is loaded into the Instruction Store
// The Instruction Pointer is incremented by 1 or 2 depending on whether the instruction is a 1 or 2 byte instruction
// The instruction in the Instruction Store is executed
// This is repeated until the instruction in the Instruction Store equals 0 (Halt)

func (mp *MicroProcessor) executeInstruction() {
	if mp.debug {
		fmt.Println(mp)
	}

	switch mp.registers.IS {
	case 0:
		cprint("Halting now..")
		mp.registers.Status = 999 // arbitary
		return
	case 1:
		cprint(fmt.Sprintf("Adding R1: %d to R0: %d and storing in R0", mp.registers.R1, mp.registers.R0))
		mp.registers.R0 = mp.registers.R0 + mp.registers.R1
	case 2:
		cprint(fmt.Sprintf("Subtracting R1: %d from R0: %d and storing in R0", mp.registers.R1, mp.registers.R0))
		fmt.Println("Subtracting now..")
		mp.registers.R0 = mp.registers.R0 - mp.registers.R1
	case 3:
		cprint("Incrementing R0..")
		mp.registers.R0 = mp.registers.R0 + 1
	case 4:
		cprint("Incrementing R1..")
		mp.registers.R0 = mp.registers.R1 + 1
	case 5:
		cprint("decrementing R0..")
		mp.registers.R0 = mp.registers.R0 - 1
	case 6:
		cprint("decrementing R1..")
		mp.registers.R1 = mp.registers.R1 - 1
	case 7:
		cmsg := ansi.Color(fmt.Sprintf("R0 contains:: %d", mp.registers.R0), "cyan")
		fmt.Println(cmsg)

	// Instructions 8 and above are 2 byte instructions,
	// 2nd Byte is <data>

	case 8:
		cprint(fmt.Sprintf("Load <data:%d> to R0", mp.memory[mp.registers.IP-1]))
		mp.registers.R0 = mp.memory[mp.registers.IP-1]
	case 9:
		cprint(fmt.Sprintf("Load <data:%d> to R1", mp.memory[mp.registers.IP-1]))
		mp.registers.R1 = mp.memory[mp.registers.IP-1]
	case 10:
		cprint(fmt.Sprintf("Load R0 to Memory Address <data:%d>", mp.memory[mp.registers.IP-1]))
		mp.memory[mp.registers.IP-1] = mp.registers.R0
	case 11:
		cprint(fmt.Sprintf("Load R1 to Memory Address <data:%d>", mp.memory[mp.registers.IP-1]))
		mp.memory[mp.registers.IP-1] = mp.registers.R1
	case 12:
		cprint(fmt.Sprintf("Jump to Memory Address <data:%d>", mp.memory[mp.registers.IP-1]))
		mp.registers.IP = mp.memory[mp.registers.IP-1]
	case 13:
		cprint(fmt.Sprintf("Store val from Memory Address <data:%d> in R0", mp.memory[mp.registers.IP-1]))
		mp.registers.R0 = mp.memory[mp.registers.IP-1]
	case 14:
		cprint(fmt.Sprintf("Jump to Memory Address <data:%d> if R0 != 0", mp.memory[mp.registers.IP-1]))
		if mp.registers.R0 != 0 {
			mp.registers.IP = mp.memory[mp.registers.IP-1]
		}
	case 15:
		cprint(fmt.Sprintf("Jump to Memory Address <data:%d> if R0 == 0", mp.memory[mp.registers.IP-1]))
		if mp.registers.R0 == 0 {
			mp.registers.IP = mp.memory[mp.registers.IP-1]
		}
	}
}

func (mp *MicroProcessor) dumpMemory() {
	for v := range mp.memory {
		fmt.Println(v, " Val: ", mp.memory[v])
	}
}

func (mp *MicroProcessor) fetchExecuteLoop() {
	for {
		if mp.registers.Status == 999 {
			mp.reset()
			return
		}
		mp.registers.IS = mp.memory[mp.registers.IP]
		fmt.Println("\nInitial:: IP:", mp.registers.IP, " || IS:", mp.registers.IS)

		if mp.registers.IS >= 8 {
			mp.registers.IP += 2
		} else {
			mp.registers.IP += 1
		}
		fmt.Println("Incremented IP:", mp.registers.IP)
		mp.executeInstruction()
	}
}

func (mp *MicroProcessor) loadProgram(program [16]int) {
	for k, v := range program {
		if v < 0 || v > 15 {
			fmt.Println("bleurgh, buffer overflow!")
			return
		}
		// fmt.Println(k, v)
		mp.memory[k] = v
	}
}

func (mp *MicroProcessor) reset() {
	mp.registers.IP = 0
	mp.registers.IS = 0
	mp.registers.R0 = 0
	mp.registers.R1 = 0
	mp.registers.Status = 0
	for i, _ := range mp.memory {
		mp.memory[i] = 0
	}
}

func newMP(mpChan chan *MicroProcessor) {
	mp := new(MicroProcessor)
	mpChan <- mp
}

func (mp *MicroProcessor) String() string {
	return fmt.Sprintf("Registers:: IP: %d IS: %d R0: %d R1: %d, Status: %b\nMemory:: %d",
		mp.registers.IP, mp.registers.IS, mp.registers.R0, mp.registers.R1, mp.registers.Status, mp.memory)
}
