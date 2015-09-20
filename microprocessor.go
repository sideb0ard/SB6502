package main

// The SB6502, a virtual clone of the MOS 6502

import "fmt"

// Instruction Loop:
// The instruction at the address given by the Instruction Pointer is loaded into the Instruction Store
// The Instruction Pointer is incremented by 1 or 2 depending on whether the instruction is a 1 or 2 byte instruction
// The instruction in the Instruction Store is executed
// This is repeated until the instruction in the Instruction Store equals 0 (Halt)

func (mp *MicroProcessor) executeInstruction() {
	fmt.Printf("Execution :: IP: %d, IS: %d, R0: %d, R1: %d\n",
		mp.registers.IP, mp.registers.IS,
		mp.registers.R0, mp.registers.R1)
	fmt.Println("Memory: ", mp.memory)

	switch mp.registers.IS {
	case 0:
		fmt.Println("Halting now..")
		mp.registers.Status = 999 // arbitary
		return
	case 1:
		fmt.Println("Adding")
		mp.registers.R0 = mp.registers.R0 + mp.registers.R1
	case 2:
		fmt.Println("Subtracting now..")
		mp.registers.R0 = mp.registers.R0 - mp.registers.R1
	case 3:
		fmt.Println("Incrementing R0..")
		mp.registers.R0 = mp.registers.R0 + 1
	case 4:
		fmt.Println("Incrementing R1..")
		mp.registers.R0 = mp.registers.R1 + 1
	case 5:
		fmt.Println("decrementing R0..")
		mp.registers.R0 = mp.registers.R0 - 1
	case 6:
		fmt.Println("decrementing R1..")
		mp.registers.R1 = mp.registers.R1 - 1
	case 7:
		fmt.Println("**Ring Bell**")

	case 8:
		fmt.Println("Print Data in R0..")
		fmt.Println(mp.registers.R0)

	// Instructions 9 and above are 2 byte instructions,
	// 2nd Byte is <data>

	case 9:
		fmt.Println("Load <data> to R0", mp.memory[mp.registers.IP-1])
		mp.registers.R0 = mp.memory[mp.registers.IP-1]
	case 10:
		fmt.Println("Load <data> to R1")
		mp.registers.R1 = mp.memory[mp.registers.IP-1]
	case 11:
		fmt.Println("Store R0 to Address")
		mp.memory[mp.registers.IP-1] = mp.registers.R0
	case 12:
		fmt.Println("Store R1 to Address ")
		mp.memory[mp.registers.IP-1] = mp.registers.R1
	case 13:
		fmt.Println("Jump To Address..")
		mp.registers.IP = mp.memory[mp.registers.IP-1]
	case 14:
		fmt.Println("Jump To Address if R0..")
		if mp.registers.R0 == 0 {
			mp.registers.IP = mp.memory[mp.registers.IP-1]
		}
	case 15:
		fmt.Println("Jump To Address if Not R0..")
		if mp.registers.R0 != 0 {
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
			// reset CPU and return
			mp.registers.Status = 0
			return
		}
		mp.registers.IS = mp.memory[mp.registers.IP]
		fmt.Println("\nInitial:: IP:", mp.registers.IP, " || IS:", mp.registers.IS)

		if mp.registers.IS >= 9 {
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

func newMP(mpChan chan *MicroProcessor) {
	mp := new(MicroProcessor)
	mpChan <- mp
}

func (mp *MicroProcessor) String() string {
	return fmt.Sprintf("Registers:: IP: %d IS: %d R0: %d R1: %d\nMemory:: %d",
		mp.registers.IP, mp.registers.IS, mp.registers.R0, mp.registers.R1, mp.memory)
}
