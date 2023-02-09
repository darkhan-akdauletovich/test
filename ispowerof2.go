package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	defer z01.PrintRune('\n')
	flag := "true"
	if len(os.Args) == 2 {
		num, _ := strconv.Atoi(os.Args[1])
		for num > 1 {
			if num%2 != 0 {
				flag = "false"
				break
			}
			num /= 2
		}
	}
	for _, r := range flag {
		z01.PrintRune(r)
	}
}
