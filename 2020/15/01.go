package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fenglyu/adventofcode/util"
)

func main() {

	report := util.ParseBasedOnEachLine()
	//fmt.Println(report)
	turns := make([]int, 0)

	for _, v := range strings.Split(report[0], ",") {
		val, _ := strconv.Atoi(v)
		turns = append(turns, val)
	}
	fmt.Println(turns)

	count, limit := 0, 2020

	i := 0
	freq := make(map[int]int)
	for count < limit {
		
	}
}
