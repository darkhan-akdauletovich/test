package main

import "os"

func main() {
	if len(os.Args) == 2 {
		arg := os.Args[1]
		for _, v:= range arg{
			if v>='a' && v<='z'{
				for i:=0;i<=int(v)-'a';i++{
					z01.PrintRune(v)
				}
			}
			if v>='A' && v<='Z'{
				for i:=0;i<=int(v)-'A';i++{
					z01.PrintRune(v)
				}
			} else {
				z01.PrintRune(v)
			}
		}
		z01.PrintRune('\n')
	}
	}
}