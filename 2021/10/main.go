package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/fenglyu/adventofcode/util"
)

var char map[uint8]uint8 = map[uint8]uint8{')': '(', ']': '[', '}': '{', '>': '<', '(': ')', '[': ']', '{': '}', '<': '>'}

func main() {
	report := util.ParseBasedOnEachLine()

	table := make(map[uint8]int, 0)
	//for _, v := range report[:1] {
	tps := make([]int, 0)
	for _, v := range report {
		i, c, q := navigation(v)
		//fmt.Printf("idx %d, c '%c' = %d \n", i, c, c)
		if i != len(v) {
			table[c]++
		} else {
			val := part2point(q)
			tps = append(tps, val)
		}
	}

	sort.Ints(tps)
	fmt.Println("part 1: ", part1point(table))
	fmt.Println("part 2: ", tps[int(math.Round(float64(len(tps)/2)))])
}

func part2point(queue *util.LIFO) int {
	var m map[uint8]int = map[uint8]int{')': 1, ']': 2, '}': 3, '>': 4}
	sum := 0
	for e := queue.Front(); e != nil; e = e.Next() {
		v := e.Value.(uint8)
		rv := char[v]
		sum = sum*5 + m[rv]
	}

	return sum
}

func part1point(table map[uint8]int) int {
	var m map[uint8]int = map[uint8]int{')': 3, ']': 57, '}': 1197, '>': 25137}
	sum := 0
	for k, v := range table {
		sum += m[k] * v
	}
	return sum
}

func navigation(subsys string) (int, uint8, *util.LIFO) {
	lifo := util.NewLIFO()
	for i := 0; i < len(subsys); i++ {
		v := subsys[i]
		switch v {
		case '{', '[', '(', '<':
			lifo.Enqueue(v)
		case '}':
			pop := lifo.Dequeue().(uint8)
			if pop != '{' {
				return i, v, lifo
			}
		case ']':
			pop := lifo.Dequeue().(uint8)
			if pop != '[' {
				return i, v, lifo
			}
		case ')':
			pop := lifo.Dequeue().(uint8)
			if pop != '(' {
				return i, v, lifo
			}
		case '>':
			pop := lifo.Dequeue().(uint8)
			if pop != '<' {
				return i, v, lifo
			}
		}
	}

	return len(subsys), '0', lifo
}
