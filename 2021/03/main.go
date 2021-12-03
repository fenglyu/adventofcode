package main

import (
	"fmt"
	"strconv"

	"github.com/fenglyu/adventofcode/util"
)

func main() {
	report := util.ParseBasedOnEachLine()

	arr := make([]uint16, len(report))
	for i, r := range report {
		v, _ := strconv.ParseUint(r, 2, 16)
		arr[i] = uint16(v)
	}

	count := make([]int, len(report[0]))
	for j := range arr {
		for i := 0; i < len(count); i++ {
			//v := ((1 << (len(count) - 1 - i)) & (0b111110110111)) >> (len(count) - 1 - i)
			v := ((1 << (len(count) - 1 - i)) & (arr[j])) >> (len(count) - 1 - i)
			if v == 1 {
				count[i]++
			}
		}
	}

	gamma, epsilon := 0, 0
	for _, v := range count {
		if v > len(report)/2 {
			gamma = (gamma << 1) + 1
			epsilon = (epsilon << 1) + 0
		} else {
			epsilon = (epsilon << 1) + 1
			gamma = (gamma << 1) + 0
		}
	}

	fmt.Println("gamma ", gamma, "epsilon ", epsilon)
	fmt.Println("part 1: ", gamma*epsilon)

	// part 2

}
