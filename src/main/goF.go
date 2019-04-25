/// FORTH implementation in Golang

package main

import (
	"fmt"
	"lexer"
	"os"
)

/// main memory size, bytes
const Msz = 0x1000

/// instruction pointer
var Ip = 0

/// current instruction opcode
var op byte = 0

/// compiler/memory allocator pointer
var Cp = 0

/// return stack, cells
const Rsz = 0x100

/// return stack pointer
var Rp = 0

/// data stack, cells
const Dsz = 0x10

/// data stack pointer
var Dp = 0

type Stack struct {
	data int
	ptr  int
}

//var M byte[];

func VM() {
	//op = M[Ip] // fetch next command
	Ip += 1 // autoincrement to next comment
	switch op {
	default:
		os.Exit(-1)
	}
}

func (s *Stack) push(value int) { s.data += value; s.ptr++ }

// func (s Stack) push(value int)  { fmt.Println("pushmanda", s) } // no poly

type FVM struct {
	Msz int
	Rsz int
	Dsz int
	Ip  int
	Cp  int
	op  byte
	Dp  int
	Rp  int
}

func NewFVM(M, R, D int) FVM {
	return FVM{
		Msz: M, Rsz: R, Dsz: D,
		Ip: 0, Rp: 0, Dp: 0,
		Cp: 0}
}

func main() {
	x := Stack{}
	fmt.Println(x)
	x.push(11)
	fmt.Println(x)
	x.push(11)
	fmt.Println(x)
	power := 42
	fmt.Println(os.Args, power, x)
	lexer.Parse("# hello")
	fmt.Println(NewFVM(Msz, Rsz, Dsz))
}
