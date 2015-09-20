package main

import (
	"fmt"

	"github.com/mgutz/ansi"
)

func newB0ard() *B0ard {
	b := &B0ard{}
	return b
}

func (b *B0ard) run(mpChan chan *MicroProcessor, pChan chan *Program) {
	for {
		select {
		case mp := <-mpChan:
			b.MicroProcessors = append(b.MicroProcessors, mp)
		case p := <-pChan:
			b.Programs = append(b.Programs, p)
		}
	}
}

func (b *B0ard) listMPs() {
	PL := ansi.Color("*** Processor List ***", "magenta")
	fmt.Println(PL)
	for i, mp := range b.MicroProcessors {
		fmt.Printf("CPU: %d\n", i)
		fmt.Println(mp)
	}
}
func (b *B0ard) listPs() {
	PL := ansi.Color("*** Program List ***", "magenta")
	fmt.Println(PL)
	for i, p := range b.Programs {
		fmt.Printf("Program: %d\n", i)
		fmt.Println(p)
	}
}
