package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/fenglyu/adventofcode/util"
)

func main() {
	startTime := time.Now()
	raw := util.ParseBasedOnEmptyLine()

	ruleStr := strings.TrimSpace(raw[0])
	rule := make([]int, len(ruleStr))

	for j, v := range []byte(ruleStr) {
		var i int
		if v == 'R' {
			i = 1
		} else if v == 'L' {
			i = 0
		}
		rule[j] = i
	}

	nodeArray := strings.Split(strings.TrimSpace(raw[1]), "\n")
	nodes := make(map[string][]string, 0)

	startSteps := make([]string, 0)
	for _, v := range nodeArray {
		nStr := strings.Split(v, " = ")
		key := nStr[0]
		dirStr := strings.Split(nStr[1], ", ")
		ls, rs := strings.TrimLeft(dirStr[0], "("), strings.TrimRight(dirStr[1], ")")
		nodes[key] = []string{ls, rs}

		if strings.HasSuffix(key, "A") {
			startSteps = append(startSteps, key)
		}
	}
	/*
		fmt.Println("rule: ", rule)
		for k, v := range nodes {
			fmt.Println("key: ", k, "value: ", v)
		}
	*/

	key := "AAA"
	count := 0
	i := 0
	length := len(rule) - 1
	for {
		//fmt.Println(i, rule[i], length, key)

		if key == "ZZZ" {
			break
		}

		if _, Ok := nodes[key]; !Ok {
			break
		}
		key = nodes[key][rule[i]]
		count++

		if i == length {
			i = 0
			continue
		}
		i++
	}
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Println("Execution time: ", duration)
	fmt.Println("Problem 1: ", count)

	var keyList []string = startSteps
	// bruteforce is silly
	// enlights from https://elixirforum.com/t/advent-of-code-2023-day-8/60244/2

	var l int64 = 1
	for _, v := range keyList {
		val := findTarget(v, "Z", rule, nodes)
		l = util.Lcm(int64(val), l)
	}
	endTime = time.Now()
	duration = endTime.Sub(startTime)
	fmt.Println("Execution time: ", duration)
	fmt.Println("Problem 2: ", l)
}

func findTarget(start, end string, rule []int, nodes map[string][]string) int {
	key := start
	count := 0
	i := 0
	length := len(rule) - 1
	for {
		//fmt.Println("find; ", i, rule[i], length, key)

		if strings.HasSuffix(key, end) {
			break
		}

		if _, Ok := nodes[key]; !Ok {
			break
		}
		key = nodes[key][rule[i]]
		count++

		if i == length {
			i = 0
			continue
		}
		i++
	}
	return count
}
