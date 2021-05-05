package main

import (
	"fmt"
)

func main() {
	numbers := []int{2, 4, 1, 3, 5}
	fmt.Println(numbers)
	bubble_sort(numbers)
	//fmt.Println(numbers)
}
func bubble_sort(numbers []int) {
	//var sorted bool ;
	//sorted=false
	var size int
	size = len(numbers)
	var iterator int
	var end int = size
	for end = size; end > 1; end-- {
		for iterator = 0; iterator < end-1; iterator++ {
			if numbers[iterator] > numbers[iterator+1] {
				numbers[iterator], numbers[iterator+1] = numbers[iterator+1], numbers[iterator]
			}
			iterator++
		}
	}
}
