package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/fenglyu/adventofcode/util"
)

func main() {

	report := util.ParseBasedOnEachLine()
	length := len(report)

	example := make([]int, length)
	for i := 0; i < length; i++ {
		v, _ := strconv.Atoi(report[i])
		example[i] = v
	}

	fmt.Println(len(example) == len(report))

	part1 := 0
	previous := 25
	for i := previous; i < length; i++ {
		if TwoSum(example[i-previous:i], example[i]) {
			//fmt.Println(i, example[i], "found")
		} else {
			fmt.Println(i, example[i], "not found")

			part1 = example[i]
			break
		}
	}
	fmt.Println("part1 > ", part1)

	l, r := 0, 0
	
	len := 1
	for i := 0; i < length-1; i++ {
		j := i + len
		sum := example[i]
		for j < length {
			sum += example[j]
			if sum == part1 {
				//fmt.Println(i, j)
				l, r = i, j
				goto breakHere
			} else if sum < part1 {
				j++
			} else if sum > part1 {
				break
			}
		}
	}

breakHere:
	min, max := maxAmin(example[l : r+1])
	fmt.Println("part2 >", min+max)
	fmt.Println("end")
}

func TwoSum(ar []int, target int) bool {
	lst := make([]int, len(ar))
	copy(lst, ar)
	sort.Ints(lst)

	i, j := 0, len(ar)-1
	for i < j {
		sum := lst[i] + lst[j]
		switch {
		case sum > target:
			j--
		case sum < target:
			i++
		case sum-target == 0:
			return true
		}
	}

	return false
}

func maxAmin(ar []int) (int, int) {
	lst := make([]int, len(ar))
	copy(lst, ar)
	sort.Ints(lst)
	return lst[0], lst[len(lst)-1]
}
