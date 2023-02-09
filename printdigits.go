package main

import "github.com/01-edu/z01"

func main() {
	for sg := '0'; sg <= '9'; sg++ {
		z01.PrintRune(sg)
	}
	z01.PrintRune('\n')
}
