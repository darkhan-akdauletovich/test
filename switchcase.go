package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1]
	for _, v := range arg {
		if v >= 'a' && v <= 'z' {
			v -= 32
		} else if v >= 'A' && v <= 'Z' {
			v += 32
		}
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
