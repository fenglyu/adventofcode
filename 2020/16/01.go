package main

import (
	"fmt"
	"strconv"
	"strings"

	set "github.com/deckarep/golang-set"
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

func (s *Seg) String() string {
	return fmt.Sprintf("[%d-%d]", s.min, s.max)
}

type Rule struct {
	sgs []*Seg
}

func (r *Rule) Valid(v int) bool {
	s1, s2 := r.sgs[0], r.sgs[1]
	return s1.Valid(v) || s2.Valid(v)
}

func (r *Rule) String() string {
	return fmt.Sprintf("%s | %s", r.sgs[0], r.sgs[1])
}

// the mapping of rules name and its applied fileds count
var gstat map[string]set.Set

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
	gstat = make(map[string]set.Set)

	for j := 0; j < len(nearbyticket[0])-1; j++ {
		ar := make([]int, 0)
		for i := 0; i < len(nearbyticket)-1; i++ {
			if _, ok := errorLine[i]; ok {
				continue
			}
			ar = append(ar, nearbyticket[i][j])
		}
		validField(fieldsDict, ar, j)
	}
	/*
		for k, v := range gstat {
			fmt.Printf("%s -> %v\n", k, v)
		}
	*/
	output := make(map[string]set.Set)
	for {
		if len(gstat) < 1 {
			break
		}

		var t set.Set
		for k, v := range gstat {
			slice := v.ToSlice()
			if len(slice) == 1 {
				output[k] = v
				t = v.Clone()
				delete(gstat, k)
			}
		}

		for key, val := range gstat {
			temp := val.Difference(t)
			gstat[key] = temp
		}
	}
	/*
		for k, v := range output {
			fmt.Printf("%s -> %v\n", k, v)
		}
	*/
	mul := 1
	for k, v := range output {
		if strings.Contains(k, "departure") {
			i := v.Pop().(int)
			mul *= myticket[i]
		}
	}

	fmt.Println("part 2", mul)
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

func validField(fieldsDict map[string]interface{}, col []int, colNum int) {
	for k, val := range fieldsDict {
		r := val.(*Rule)
		flag := true
		for _, v := range col {
			if !r.Valid(v) {
				flag = false
				//break on expression colNum == 18 && v == 138
				//fmt.Println("colNum: ", colNum, v, " invalid:", k, "rule :", r, idx)
				break
			}
		}
		if flag {
			if _, Ok := gstat[k]; Ok {
				fields := gstat[k]
				fields.Add(colNum)
				gstat[k] = fields
			} else {
				fields := set.NewSet()
				fields.Add(colNum)
				gstat[k] = fields
			}
		}
	}
}
