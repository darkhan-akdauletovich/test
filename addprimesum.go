package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isPrime(n int) bool {
	for x := 2; x < n; x++ {
		if n%x == 0 {
			return false
		}
	}
	return true
}
func myAtoi(s string) int {
	n := 0
	for _, v := range s {
		if v >= '0' && v <= '9' {
			n = n*10 + int(v-'0')
		} else {
			return 0
		}
	}
	return n
}
func main() {
	arg := os.Args[1:]
	defer z01.PrintRune('\n')

	if len(arg) != 1 {
		z01.PrintRune('0')
		return
	}
	num := myAtoi(arg[0])
	sum := 0
	for x := 2; x <= num; x++ {
		if isPrime(x) {
			sum += x
		}
	}
	result := ""
	n := 0
	for sum > 0 {
		n = sum % 10
		for _, v := range string(n) {
			result = string(v+'0') + result
		}
		sum = sum / 10
	}
	for _, v := range result {
		z01.PrintRune(v)
	}
}
