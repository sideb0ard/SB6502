package main

import (
	"fmt"
	"strconv"
	"strings"
)

func newProgram(pChan chan *Program, strProg string) {
	fmt.Println("Interpreting new program...", strProg)
	splitProg := strings.Split(strProg, " ")
	fmt.Println("Compiled:::", splitProg)
	if len(splitProg) <= 16 {
		p := &Program{}
		for i, v := range splitProg {
			num, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println("Barfed on yer program, ain't no stinking number!")
				return
			}
			p.code[i] = num

		}
		pChan <- p
	} else {
		fmt.Println("Program too long, mate, up to 16 only")
		return
	}
}

func (p *Program) String() string {
	return fmt.Sprintf("Code:: %d\n", p.code)
}
