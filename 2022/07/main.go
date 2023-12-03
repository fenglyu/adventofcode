package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fenglyu/adventofcode/util"
)

func main() {
	report := util.ParseBasedOnEachLine()
	//	fmt.Println(report)

	arr := strings.Split(report[0], ",")

	pos := make([]int, len(arr))

	for i := range arr {
		v, _ := strconv.Atoi(arr[i])
		pos[i] = v
	}

	//	fmt.Println(pos)

	uniq := make(map[int]int, 0)

	min := 0
	for i := 0; i < len(pos); i++ {
		sum := 0
		if _, Ok := uniq[pos[i]]; !Ok {
			uniq[pos[i]]++
		} else {
			continue
		}
		for j := 0; j < len(pos); j++ {
			v := pos[j] - pos[i]
			sum += util.Abs(v)
		}
		if i == 0 {
			min = sum
		}

		if sum < min {
			min = sum
		}
	}

	fmt.Println("part 1: ", min)

	l, r := util.Min(pos), util.Max(pos)

	min = 0
	for i := l; i <= r; i++ {
		sum := 0
		for j := 0; j < len(pos); j++ {
			v := fuelRate(util.Abs(pos[j] - i))
			sum += v
			//fmt.Println("Move from ", pos[j], " to ", i, v)
		}
		//fmt.Println("Total ", sum)
		if i == l {
			min = sum
		}

		if sum < min {
			min = sum
		}
	}
	fmt.Println("part 2: ", min)
}

func fuelRate(s int) int {
	return s * (1 + s) / 2
}
