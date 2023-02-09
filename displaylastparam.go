package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 1 {
		param := os.Args[len(os.Args)-1]
		for _, v := range param {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}
