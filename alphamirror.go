package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	if len(arg) == 1 {
		str := arg[0]
		for _, v := range str {
			if v >= 'a' && v <= 'z' {
				z01.PrintRune('z' - (v - 'a'))
			} else if v >= 'A' && v <= 'Z' {
				z01.PrintRune('Z' - (v - 'A'))
			} else {
				z01.PrintRune(v)
			}
		}
		z01.PrintRune('\n')
	}
}
