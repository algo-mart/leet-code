package main

import (
	"fmt"
	"math"
)

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")



func Solution6(blocks []int) int {
	// write your code in Go 1.4
	a:= 0
	max := 0

	for a < len(blocks){
		right:= a
		left:= a
		for right < len(blocks)-1 && blocks[right] <= blocks[right+1]{
			right++
		}
		for left > 0 && blocks[left] <= blocks[left-1]{
			left--
		}
		if right-left > max{
			max = right-left
		}
		a = right+1
	}
	return max + 1
}

func Solution7(text string) int {
	// write your code in Go 1.4
	store := map[rune]int{'b': 0, 'a': 0, 'l': 0, 'o': 0, 'n': 0}
	a := 0
	min := math.MaxInt64
	for a < len(text) {
		if _, ok := store[rune(text[a])]; ok {
			store[rune(text[a])]++
		}
		a++
	}
	store['l'] = store['l'] / 2
	store['o'] = store['o'] / 2
	for _, v := range store {
		if v < min {
			min = v
		}
	}
	return min
}

func main() {
	fmt.Println(Solution7("nlaebolko"))
	//fmt.Println(Solution([]int {2, 5, 5, 2, 6}))
	//fmt.Println(Solution([]int {1,1}))
	//fmt.Println(Solution([]int {1,1}))
}
