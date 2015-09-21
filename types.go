package main

type RegisterBank struct {
	IP     int // instruction pointer
	IS     int // instruction store
	R0     int // register zero
	R1     int // register one
	Status int
}

type MicroProcessor struct {
	registers RegisterBank
	memory    [32]int
	debug     bool
}

type Program struct {
	code [32]int
}

type B0ard struct {
	MicroProcessors []*MicroProcessor
	Programs        []*Program
}
