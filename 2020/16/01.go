package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fenglyu/adventofcode/util"
)

type Seg struct {
	min int
	max int
}

func (s *Seg) Valid(v int) bool {
	return s.min <= v && v <= s.max
}

func newSeg(seg string) *Seg {
	segs := strings.Split(seg, "-")
	min, _ := strconv.Atoi(segs[0])
	max, _ := strconv.Atoi(segs[1])
	return &Seg{min: min, max: max}
}

type Rule struct {
	segs []*Seg
}

func main() {

	report := util.ParseBasedOnEmptyLine()
	//fmt.Println(report)
	//fmt.Println(len(report))
	rawFields, rawMyTicket, rawNearbyTickets := report[0], report[1], report[2]

	fieldsDict := make(map[string]interface{}, 0)

	for _, rf := range strings.Split(rawFields, "\n") {
		eachRf := strings.Split(rf, ": ")
		field, rawRanges := eachRf[0], eachRf[1]

		ranges := strings.Split(rawRanges, " or ")
		res := make([]*Seg, 0)
		for _, v := range ranges {
			res = append(res, newSeg(v))
		}
		fieldsDict[field] = res
	}

	myticket := make([]int, 0)
	for i, rf := range strings.Split(rawMyTicket, "\n") {
		if i == 0 || strings.Contains(rf, ":") {
			continue
		}
		for _, v := range strings.Split(rf, ",") {
			if strings.EqualFold(v, "") {
				continue
			}
			num, _ := strconv.Atoi(v)
			myticket = append(myticket, num)
		}
	}

	nearbyticket := make([][]int, 0)
	for i, rf := range strings.Split(rawNearbyTickets, "\n") {

		if i == 0 || strings.Contains(rf, ":") {
			continue
		}
		nbticket := make([]int, 0)
		for _, v := range strings.Split(rf, ",") {
			num, _ := strconv.Atoi(v)
			nbticket = append(nbticket, num)
		}
		nearbyticket = append(nearbyticket, nbticket)
	}

	//fmt.Println("fieldsDict: ", fieldsDict)
	//fmt.Println("myticket: ", myticket)
	//fmt.Println("nearbyticket: ", nearbyticket)

	errorRate := make([]int, 0)
	for _, nt := range nearbyticket {
		for _, val := range nt {
			if !valid(fieldsDict, val) {
				errorRate = append(errorRate, val)
			}
		}
	}

	sum := 0
	for _, v := range errorRate {
		sum += v
	}

	fmt.Println("part 1", sum)

	stats := make(map[string]int)

	for j := 0; j < len(nearbyticket[0])-1; j++ {
		for i := 0; i < len(nearbyticket)-1; i++ {

		}
	}
}

func valid(fieldsDict map[string]interface{}, value int) bool {
	for _, v := range fieldsDict {
		for _, seg := range v.([]*Seg) {
			if seg.Valid(value) {
				return true
			}
		}
	}

	return false
}

func validField(fieldsDict map[string]interface{}, nearbyticket [][]int, col int) bool {

	for k, v := range fieldsDict {
		for _, seg := range v.([]*Seg) {
			if seg.Valid(nearbyticket[i][j]) {
				flag = false
				break
			}
		}
	}

}
