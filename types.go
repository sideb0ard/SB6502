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
	memory    [16]int
}

type B0ard struct {
	MicroProcessors []*MicroProcessor
	Programs        []*Program
}

type Program struct {
	code [16]int
}
