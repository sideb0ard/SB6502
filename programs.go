package main

import (
	"fmt"
	"strconv"
	"strings"
)

func newProgram(pChan chan *Program, strProg string) {
	fmt.Println("Interpreting new program...", strProg)
	p := &Program{}
	p.Update(strProg)
	pChan <- p

	//splitProg := strings.Split(strProg, " ")
	//fmt.Println("Compiled:::", splitProg)
	//if len(splitProg) <= 16 {
	//	p := &Program{}
	//	for i, v := range splitProg {
	//		num, err := strconv.Atoi(v)
	//		if err != nil {
	//			fmt.Println("Barfed on yer program, ain't no stinking number!")
	//			return
	//		}
	//		p.code[i] = num

	//	}
	//	pChan <- p
	//} else {
	//	fmt.Println("Program too long, mate, up to 16 only")
	//	return
	//}
}

func (p *Program) Update(strProg string) {
	for i := range p.code {
		p.code[i] = 0
	}
	splitProg := strings.Split(strProg, " ")
	if len(splitProg) <= 16 {
		for i, v := range splitProg {
			num, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println("Barfed on yer program, ain't no stinking number!")
				return
			}
			p.code[i] = num
		}
	}
}

func (p *Program) String() string {
	return fmt.Sprintf("Code:: %d\n", p.code)
}
