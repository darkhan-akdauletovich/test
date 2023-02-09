package main

import "github.com/01-edu/z01"

func main() {
	str1 := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	str2 := "abcdefghijklmnopqrstuvwxyz"
	st1 := []rune(str1)
	st2 := []rune(str2)
	for i := 25; i >= 0; i-- {
		if i%2 != 0 {
			z01.PrintRune(st1[i])
		} else {
			z01.PrintRune(st2[i])
		}
	}
	z01.PrintRune('\n')
}
