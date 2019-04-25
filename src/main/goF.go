/// FORTH implementation in Golang

package main

import (
	"fmt"
	"frame"
	"lexer"
	"os"
)

type Stack struct { /// unused
	size int   /// max size
	data []int /// data container
	top  int   /// top pointer
}

type FVM struct {
	/// configuration
	Msz int /// program/data memory size, bytes
	Rsz int /// return stack size, cells
	Dsz int /// data stack size, cells
	/// VM registers & program/data memory
	M  []byte
	Ip int  /// instruction pointer
	op byte /// current opcode
	Cp int  /// compiler/memory allocator pointer
	/// return stack
	R  []int
	Rp int /// pointer
	/// data stack
	D  []int
	Dp int /// pointer
}

func NewFVM(M, R, D int) FVM {
	return FVM{
		Msz: M, Rsz: R, Dsz: D,
		Ip: 0, Rp: 0, Dp: 0,
		Cp: 0}
}

var vm = FVM{Msz: 0x1000, Rsz: 0x100, Dsz: 0x10}

func main() {
	fmt.Print(os.Args)
	fmt.Println(vm)
	lexer.Parse("# hello")
	fmt.Println(frame.Frame{Tag: "hello", Val: "world"})
}
