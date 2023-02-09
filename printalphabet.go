package main

import "github.com/01-edu/z01"

func main() {
	for sg := 'a'; sg <= 'z'; sg++ {
		z01.PrintRune(sg)
	}
	z01.PrintRune('\n')
}
