package main

import (
	"fmt"
	"strings"

	"github.com/fenglyu/adventofcode/util"
)

func main() {
	report := util.ParseBasedOnEachLine()
	l := strings.Split(report[0], ",")
	fmt.Println(l)
}
