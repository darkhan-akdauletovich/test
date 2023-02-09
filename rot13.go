package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	if len(arg) == 1 {
		for _, v := range arg[0] {
			if v >= 'a' && v <= 'm' || v >= 'A' && v <= 'M' {
				z01.PrintRune(v + 13)
			} else if v >= 'n' && v <= 'z' || v >= 'N' && v <= 'Z' {
				z01.PrintRune(v - 13)
			} else {
				z01.PrintRune(v)
			}
		}
		z01.PrintRune('\n')
	}
}
