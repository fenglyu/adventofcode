package main

import (
	"fmt"
	"strconv"

	"github.com/fenglyu/adventofcode/util"
)

func main() {
	report := util.ParseBasedOnEachLine()
	//fmt.Println(report)

	measurements := make([]int, len(report))
	count := 0
	for i := 0; i < len(report); i++ {
		v, _ := strconv.Atoi(report[i])
		measurements[i] = v
		if i > 0 && measurements[i] > measurements[i-1] {
			count++
		}
	}
	//fmt.Println(measurements)
	fmt.Println("part 1:", count)

	pre, next, incr := 0, 0, 0
	slidingWindow := 3
	for i := slidingWindow - 1; i < len(measurements); i++ {
		next = measurements[i] + measurements[i-1] + measurements[i-2]
		if pre != 0 && next > pre {
			incr++
		}
		if pre != next {
			pre = next
		}
	}

	fmt.Println("part 2:", incr)

}
