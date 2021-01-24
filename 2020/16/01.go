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
	return s.min <= v || v <= s.max
}

func newSeg(seg string) *Seg {
	segs := strings.Split(seg, "-")
	min, _ := strconv.Atoi(segs[0])
	max, _ := strconv.Atoi(segs[1])
	return &Seg{min: min, max: max}
}

type Rule struct {
	sgs []*Seg
}

func (r *Rule) Valid(v int) bool {
	for _, s := range r.sgs {
		if !s.Valid(v) {
			return false
		}
	}
	return true
}

var gstat map[string]int

func main() {

	report := util.ParseBasedOnEmptyLine()
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
		fieldsDict[field] = &Rule{sgs: res}
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

	errorRate := make([]int, 0)
	//errorLine := make([]int, 0)
	errorLine := make(map[int]bool)
	for i, nt := range nearbyticket {
		for _, val := range nt {
			if !valid(fieldsDict, val) {
				errorRate = append(errorRate, val)
				//errorLine = append(errorLine, i)
				errorLine[i] = true
			}
		}
	}

	sum := 0
	for _, v := range errorRate {
		sum += v
	}
	fmt.Println("part 1", sum)

	nearbyticket = append(nearbyticket, myticket)
	//fmt.Println("errorLine> ", errorLine)
	//stats := make(map[int]string)
	gstat = make(map[string]int)

	stats := make([]string, len(nearbyticket[0]))
	for j := 0; j < len(nearbyticket[0])-1; j++ {
		ar := make([]int, 0)
		for i := 0; i < len(nearbyticket)-1; i++ {
			if _, ok := errorLine[i]; ok {
				continue
			}
			ar = append(ar, nearbyticket[i][j])
		}
		//fmt.Println(j, len(ar))
		if ok, con := validField(fieldsDict, ar); ok {
			stats[j] = con
		}
	}

	fmt.Println(len(stats))
	fmt.Println(stats)
	mul := 1
	for i, v := range stats {
		if strings.Contains(v, "departure") {
			mul *= myticket[i]
			fmt.Println("i :", i, v)
		}
	}
	fmt.Println(gstat)
	fmt.Println(myticket)
	fmt.Println(mul)
}

func valid(fieldsDict map[string]interface{}, value int) bool {
	for _, v := range fieldsDict {
		for _, seg := range v.(*Rule).sgs {
			if seg.Valid(value) {
				return true
			}
		}
	}
	return false
}

func validField(fieldsDict map[string]interface{}, col []int) (bool, string) {
	for k, v := range fieldsDict {
		r := v.(*Rule)
		flag := true
		for _, v := range col {
			if !r.Valid(v) {
				flag = false
				break
			}
		}
		if flag {
			gstat[k]++
			return flag, k
		}
	}
	return false, ""
}
