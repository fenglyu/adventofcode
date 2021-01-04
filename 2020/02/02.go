package main

import (
	"fmt"
	"strconv"
	"strings"

	//"strconv"

	"github.com/fenglyu/adventofcode/util"
)

func main() {
	report := util.ParseBasedOnEachLine()
	//fmt.Println(report)
	count := 0
	for _, v := range report {
		if validate(v) {
			count++
		}
	}

	fmt.Println(count)
}

func validate(line string) bool {
	res := strings.Split(line, ":")
	if len(res) < 2 {
		return false
	}
	rule, pass := res[0], res[1]
	return checkRule(rule, pass)
}

func checkRule(rule string, pass string) bool {
	res := strings.Fields(rule)
	scope, letter := res[0], res[1]

	ss := strings.Split(scope, "-")
	l, _ := strconv.Atoi(ss[0])
	r, _ := strconv.Atoi(ss[1])
	target := []byte(letter)[0]
	p := []byte(pass)
	return (p[l] == target || p[r] == target) && !(p[l] == target && p[r] == target)
}
