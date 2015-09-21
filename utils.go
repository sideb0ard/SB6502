package main

import (
	"fmt"
	"os"

	"github.com/mgutz/ansi"
)

func myexit() {
	fmt.Println("Later, dude...")
	os.Exit(0)
}

func cprint(msg string) {
	cmsg := ansi.Color(msg, "white")
	fmt.Println(cmsg)
}

func help() {
	cprint("VIRTUAL SB6502 - HELP!!")
	cprint("create a new cpu with 'new cpu'")
	cprint("create a new program with 'new prog \"<instructions>\" - see below for instruction set")
	fmt.Println()
	cprint("See whats loaded with 'ps'")
	cprint("Load Program 0 into CPU 0 with 'cpu 0 load 0'")
	cprint("Run CPU with 'cpu 0 run'")
	fmt.Println()
	cprint("Instruction Set - 16 of them:::")
	cprint("First 8 are 1 byte instructions:")
	fmt.Println("0 = Halt")
	fmt.Println("1 = Add (R0 = R0 + R1)")
	fmt.Println("2 = Subtract (R0 = R0 – R1)")
	fmt.Println("3 = Increment R0 (R0 = R0 + 1)")
	fmt.Println("4 = Increment R1 (R1 = R1 + 1)")
	fmt.Println("5 = Decrement R0 (R0 = R0 – 1)")
	fmt.Println("6 = Decrement R1 (R1 = R1 – 1)")
	fmt.Println("7 = Print Value in R0")
	cprint("Last 7 are 2 byte instructions (second byte is <data>):")
	fmt.Println("8 = Load <data> into R0")
	fmt.Println("9 = Load <data> into R1")
	fmt.Println("10 = Store R0 into address <data>")
	fmt.Println("11 = Store R1 into address <data>")
	fmt.Println("12 = Store val from address <data> in R0")
	fmt.Println("13 = Jump to address <data>")
	fmt.Println("14 = Jump to address <data> if R0 != 0")
	fmt.Println("15 = Jump to address <data> if R0 == 0")
}
