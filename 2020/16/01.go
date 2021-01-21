package main

import (
	"strings"

	"github.com/fenglyu/adventofcode/util"
)

type seg struct {
	min int
	max int
}

fucn (s *seg) Valid() bool{
	
}

func main() {

	report := util.ParseBasedOnEmptyLine()
	//fmt.Println(len(report))
	rawFields, rawMyTicket, rawNearbyTickets := report[0], report[1], report[2]

	fieldsDict = make(map[string][]int, 0)

	for _, rf := range strings.Split(rawFields, "\n") {
		eachRf := strings.Split(rf, ": ")
		field, rawRanges := eachRf[0], eachRf[1]

		eachRg := strings.Split(rawRanges, " or ")
		make()
	}

}
