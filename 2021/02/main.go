package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fenglyu/adventofcode/util"
)

func main() {
	report := util.ParseBasedOnEachLine()

	horizontal, depth := 0, 0
	for _, v := range report {
		res := strings.Split(strings.Trim(v, " "), " ")
		val, _ := strconv.Atoi(res[1])
		switch res[0] {
		case "forward":
			horizontal += val
		case "down":
			depth -= val
		case "up":
			depth += val
		}
	}

	var dep int = 0
	if depth < 0 {
		dep = -depth
	}
	fmt.Println("part 1: ", horizontal*dep)

	horizontal, depth = 0, 0
	aim := 0
	for _, v := range report {
		res := strings.Split(strings.Trim(v, " "), " ")
		val, _ := strconv.Atoi(res[1])
		switch res[0] {
		case "forward":
			horizontal += val
			depth += aim * val
		case "down":
			aim += val
		case "up":
			aim -= val
		}
	}
	dep = depth
	if depth < 0 {
		dep = -depth
	}
	fmt.Println("part 2: ", horizontal*dep)
}
