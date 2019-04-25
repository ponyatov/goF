package frame

import "fmt"

type Frame struct {
	tag  string /// class/type tag
	val  string /// primitive value
	attr *Frame /// slots = atrributes = associative array
	nest *Frame /// nested elements = vector = stack
}

var Pool *Frame

func NewFrame(tag, val string) *Frame { Pool = &Frame{tag: tag, val: val}; return Pool }
func Head(F *Frame)                   { fmt.Println("<", F.tag, ":", F.val, ">") }
