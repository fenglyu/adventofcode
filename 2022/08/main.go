package main

import (
	"fmt"
	"strings"

	set "github.com/deckarep/golang-set"
	"github.com/fenglyu/adventofcode/util"
)

func main() {
	report := util.ParseBasedOnEachLine()

	sum := 0
	part2 := 0
	for _, v := range report {
		res := strings.Split(v, "|")
		first, second := strings.Split(strings.TrimSpace(res[0]), " "), strings.Split(strings.TrimSpace(res[1]), " ")

		for j := range second {
			switch len(second[j]) {
			case 2, 3, 4, 7:
				sum++
			default:
			}
		}
		part2 += parse_signals(first, second)
	}

	fmt.Println("part 1: ", sum)
	fmt.Println("part 2: ", part2)

}

func parse_signals(input, ouput []string) int {
	// seq of the letters
	seq := make([]set.Set, 7)
	// number and the mapping for the characters
	mark := make([]set.Set, 10)
	for _, v := range input {
		arr := []byte(v)
		switch len(v) {
		case 2:
			mark[1] = setFromSlice(arr)
			//seq[2], seq[6] = arr[0], arr[1]
		case 3:
			mark[7] = setFromSlice(arr)
		case 4:
			mark[4] = setFromSlice(arr)
		case 7:
			mark[8] = setFromSlice(arr)
		}
	}

	i := 0
	for i < 2 {
		for _, v := range input {
			arr := []byte(v)
			switch i {
			case 0:
				seq[0] = mark[7].Difference(mark[1])
				tmp := mark[4].Union(mark[7])
				if len(v) == 6 && len(v) > tmp.Cardinality() {
					// 9 and 6 both have 6 elements
					sts := setFromSlice(arr)
					if !sts.IsSuperset(mark[1]) {
						mark[6] = setFromSlice(arr)
						seq[2] = mark[8].Difference(mark[6])
						seq[5] = mark[1].Difference(seq[2])
						i++
					}
				}
			case 1:
				sts := setFromSlice(arr)
				if len(v) == 5 && sts.IsSuperset(mark[1]) {
					// number 3
					mark[3] = sts
					seq[1] = mark[4].Difference(mark[3])
					seq[3] = mark[4].Difference(mark[1]).Difference(seq[1])
					seq[6] = mark[3].Difference(mark[4]).Difference(seq[0])
					seq[4] = mark[6].Difference(mark[3]).Difference(seq[1])
					i++
				}
			}
		}
	}

	mark[0] = setFromPos(seq, []int{0, 1, 2, 4, 5, 6})
	mark[2] = setFromPos(seq, []int{0, 2, 3, 4, 6})
	mark[5] = setFromPos(seq, []int{0, 1, 3, 5, 6})
	mark[9] = setFromPos(seq, []int{0, 1, 2, 3, 5, 6})

	val := 0
	for _, k := range ouput {
		a := []byte(k)
		r := setFromSlice(a)
		for i, v := range mark {
			if r.Equal(v) {
				val = val*10 + i
			}
		}
	}

	return val
}

func setFromSlice(arr []byte) set.Set {
	s := set.NewSet()
	for _, v := range arr {
		s.Add(v)
	}

	return s
}

func setFromPos(seq []set.Set, arr []int) set.Set {
	s := set.NewSet()
	for _, v := range arr {
		s = s.Union(seq[v])
	}
	return s
}
