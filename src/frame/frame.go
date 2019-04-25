package frame

import "fmt"

type Frame struct {
	Tag  string /// class/type tag
	Val  string /// primitive value
	attr *Frame /// slots = atrributes = associative array
	nest *Frame /// nested elements = vector = stack
}

var Pool *Frame

func NewFrame(tag, val string) *Frame { Pool = &Frame{Tag: tag, Val: val}; return Pool }
func Head(F *Frame)                   { fmt.Println("<", F.Tag, ":", F.Val, ">") }
