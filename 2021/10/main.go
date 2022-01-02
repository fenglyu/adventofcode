package main

import (
	"fmt"

	"github.com/fenglyu/adventofcode/util"
)

func main() {
	report := util.ParseBasedOnEachLine()

	table := make(map[uint8]int, 0)
	//for _, v := range report[:1] {
	for _, v := range report {
		i, c := navigation(v)
		//fmt.Printf("idx %d, c '%c' = %d \n", i, c, c)
		if i != len(v) {
			table[c]++
		}
	}

	fmt.Println("part 1: ", point(table))

}

func point(table map[uint8]int) int {
	var m map[uint8]int = map[uint8]int{')': 3, ']': 57, '}': 1197, '>': 25137}

	sum := 0
	for k, v := range table {
		sum += m[k] * v
	}
	return sum
}

func navigation(subsys string) (int, uint8) {
	lifo := util.NewLIFO()
	for i := 0; i < len(subsys); i++ {
		v := subsys[i]
		switch v {
		case '{', '[', '(', '<':
			lifo.Enqueue(v)
		case '}':
			pop := lifo.Dequeue().(uint8)
			if pop != '{' {
				return i, v
			}
		case ']':
			pop := lifo.Dequeue().(uint8)
			if pop != '[' {
				return i, v
			}
		case ')':
			pop := lifo.Dequeue().(uint8)
			if pop != '(' {
				return i, v
			}
		case '>':
			pop := lifo.Dequeue().(uint8)
			if pop != '<' {
				return i, v
			}
		}
	}

	return len(subsys), '0'
}
