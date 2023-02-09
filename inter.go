package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isUniq(r rune, str string) bool {
	for _, b := range str {
		if r == b {
			return false
		}
	}
	return true
}
func main() {
	defer z01.PrintRune('\n')
	if len(os.Args) == 3 {
		var res, fin string

		for _, b := range os.Args[1] {
			for _, r := range os.Args[2] {
				if r == b {
					res += string(r)
				}
			}
		}
		for _, r := range res {
			if isUniq(r, fin) {
				fin = fin + string(r)
			}
		}
		for _, r := range fin {
			z01.PrintRune(r)
		}
	}
}
