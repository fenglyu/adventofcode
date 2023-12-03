package main

import (
	"fmt"
	"strconv"

	"github.com/fenglyu/adventofcode/util"
)

var arrLen int

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

	// oxygen generator rating
	ogr := make([]int, len(report))
	// CO2 scrubber rating
	cosr := make([]int, len(report))

	for i := range ogr {
		ogr[i] = 0
		cosr[i] = 0
	}

	arrLen = len(report[0])

	m := filter(0, arr, ogr, mostCommon, 0)
	l := filter(0, arr, cosr, leastCommon, 0)

	part2 := 1
	for i := range ogr {
		if ogr[i] == m {
			//fmt.Println(i, arr[i])
			part2 *= int(arr[i])
		}
		if cosr[i] == l {
			//fmt.Println(i, arr[i])
			part2 *= int(arr[i])
		}
	}

	fmt.Println("part 2: ", part2)
}

type bitCriteria func(a, b, val int) int

func mostCommon(ogrC, cosrC int, val int) int {
	if ogrC >= cosrC {
		return val
	} else {
		return -val
	}
}

func leastCommon(ogrC, cosrC int, val int) int {
	if ogrC >= cosrC {
		return -val
	} else {
		return val
	}
}

func filter(i int, input []uint16, rate []int, f bitCriteria, refer int) int {
	if i >= arrLen {
		return refer
	}

	ogrC, cosrC := 0, 0

	for j := range input {
		if rate[j] != refer {
			continue
		}

		v := ((1 << (arrLen - 1 - i)) & (input[j])) >> (arrLen - 1 - i)
		if v == 1 {
			ogrC++
			rate[j] = i + 1
		} else {
			cosrC++
			rate[j] = -(i + 1)
		}
	}

	ref := f(ogrC, cosrC, i+1)

	if ogrC == 1 && cosrC == 0 {
		return i + 1
	} else if ogrC == 0 && cosrC == 1 {
		return -(i + 1)
	}

	return filter(i+1, input, rate, f, ref)
}
