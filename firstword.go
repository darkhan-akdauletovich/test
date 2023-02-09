package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}

func main() {
	arg := os.Args[1:]
	if len(arg) == 1 {
		s := ""
		for _, v := range arg[0] {
			if v == ' ' {
				printStr(s)
				return
			} else {
				s = s + string(v)
			}
		}
		if arg[0] != "" {
			printStr(arg[0])
		}
		z01.PrintRune('\n')
	}
}
