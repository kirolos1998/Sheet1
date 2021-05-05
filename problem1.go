package main

import (
	"fmt"
)

func main() {
	fmt.Println("kirolos")
	PerfectNumber()
}
func PerfectNumber() {
	var maxnumber int = 10000
	for i := 1; i < maxnumber; i++ {
		var sum int = 0
		for j := 1; j < i; j++ {
			if i%j == 0 {
				sum = sum + j
			}
		}
		if sum == i {
			fmt.Println(i)
		}
	}
}
