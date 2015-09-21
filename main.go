package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/mgutz/ansi"
)

func main() {

	PS2 := ansi.Color("#SB6502> ", "blue")

	// Check GOMAXPROCS
	max_procs := runtime.GOMAXPROCS(-1)
	num_procs_to_set := runtime.NumCPU()
	if max_procs != num_procs_to_set {
		runtime.GOMAXPROCS(num_procs_to_set)
	}

	mpChan := make(chan *MicroProcessor)
	pChan := make(chan *Program)
	b0ard := newB0ard()
	go b0ard.run(mpChan, pChan)

	shellReader := bufio.NewReader(os.Stdin)
	for {
		// SETUP && LOOP
		fmt.Printf(PS2)
		input, err := shellReader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				myexit()
			} else {
				fmt.Println("Got an err for ye: ", err)
			}
		}
		input = strings.TrimSpace(input)

		// ALL DA REGEX FROM HERE ON OUT..

		ncp, _ := regexp.MatchString("(?i)^new CPU$", input)
		if ncp {
			fmt.Println("New CPU added")
			newMP(mpChan)
		}

		lpx := regexp.MustCompile("^cpu ([0-9]+) load ([0-9]+)$")
		lpxr := lpx.FindStringSubmatch(input)
		if len(lpxr) == 3 {
			cpuNum, err := strconv.Atoi(lpxr[1])
			if err != nil {
				fmt.Println("CPU not an int, ya numpty")
				continue
			}
			progNum, err := strconv.Atoi(lpxr[2])
			if err != nil {
				fmt.Println("ProgNum not an int, ya numpty")
				continue
			}

			if len(b0ard.MicroProcessors) > cpuNum && len(b0ard.Programs) > progNum {
				b0ard.MicroProcessors[cpuNum].loadProgram(b0ard.Programs[progNum].code)
			} else {
				fmt.Println("good try mate, but something is fucked")
				continue
			}

		}

		crx := regexp.MustCompile("^cpu ([0-9]+) run$")
		crxr := crx.FindStringSubmatch(input)
		if len(crxr) == 2 {
			cpuNum, _ := strconv.Atoi(crxr[1])
			if len(b0ard.MicroProcessors) > cpuNum {
				b0ard.MicroProcessors[cpuNum].fetchExecuteLoop()
			}
		}

		cdx := regexp.MustCompile("^cpu ([0-9]+) debug$")
		cdxr := cdx.FindStringSubmatch(input)
		if len(cdxr) == 2 {
			cpuNum, _ := strconv.Atoi(cdxr[1])
			if len(b0ard.MicroProcessors) > cpuNum {
				b0ard.MicroProcessors[cpuNum].debug = true
			}
		}

		npx := regexp.MustCompile("^new prog \"([0-9 ]+)\"$")
		npxr := npx.FindStringSubmatch(input)
		if len(npxr) == 2 {
			newProgram(pChan, npxr[1])
		}

		rpx := regexp.MustCompile("^prog ([0-9]+) \"([0-9 ]+)\"$") // r for replace
		rpxr := rpx.FindStringSubmatch(input)
		if len(rpxr) == 3 {
			progNum, _ := strconv.Atoi(rpxr[1])
			if len(b0ard.Programs) > progNum {
				cprint("Updating program")
				b0ard.Programs[progNum].Update(rpxr[2])
			}
		}

		//// PROCESS LIST
		psx, _ := regexp.MatchString("^ps$", input)
		if psx {
			// Processors
			b0ard.listMPs()
			fmt.Println()
			// Programs
			b0ard.listPs()
			fmt.Println()
		}

		hx, _ := regexp.MatchString("^help$", input)
		if hx {
			help()
		}
	}
}
