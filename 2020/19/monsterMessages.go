package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fenglyu/adventofcode/util"
)

type key struct {
	id  uint8
	pos int
}

var mapp map[uint8]any
var memo map[key][]int

func concatBytes(res string) []uint8 {
	a := strings.Split(res, " ")
	u8a := make([]uint8, len(a))
	for i := 0; i < len(a); i++ {
		ia, _ := strconv.Atoi(a[i])
		u8a[i] = uint8(ia)
	}
	return u8a
}

func main() {

	report := util.ParseBasedOnEmptyLine()
	//fmt.Println(report, len(report))

	mapp = make(map[uint8]any)
	for _, v := range strings.Split(report[0], "\n") {
		r := strings.Split(v, ": ")
		idInt32, _ := strconv.Atoi(r[0])
		pos := uint8(idInt32)
		if strings.Contains(r[1], "|") {
			// 93: 57 68 | 12 110
			r := strings.Split(r[1], " | ")
			n := [][]uint8{concatBytes(r[0]), concatBytes(r[1])}
			mapp[pos] = n
		} else if strings.Contains(r[1], "\"") {
			// 12: "a"
			v := []byte(strings.Trim(r[1], "\""))
			//mapp[pos] = bytes.Trim([]byte(r[1]), "\"")
			mapp[pos] = uint8(v[0])
		} else {
			// 0: 8 11
			mapp[pos] = concatBytes(r[1])
		}
	}

	fmt.Println(mapp)
	fmt.Println(matchRule([]byte("ababbb"), 0, 0))
}

func matchRule(str []byte, row uint8, pos int) []int {
	k := key{id: row, pos: pos}
	if res, Ok := memo[k]; Ok {
		return res
	}

	value, Ok := mapp[row]
	if !Ok {
		return nil
	}

	var out []int
	switch v := value.(type) {
	case [][]uint8:
		var results []int
		for _, seq := range v {
			// run the same sequence logic as above
			positions := []int{pos}
			for _, sub := range seq {
				var next []int
				for _, p := range positions {
					ends := matchRule(str, sub, p)
					if len(ends) > 0 {
						next = append(next, ends...)
					}
				}
				positions = next
				if len(positions) == 0 {
					break
				}
			}
			results = append(results, positions...)
		}
		out = results
	case []uint8:
		positions := []int{pos}
		for _, sub := range v {
			var next []int
			for _, p := range positions {
				ends := matchRule(str, sub, p)
				if len(ends) > 0 {
					next = append(next, ends...)
				}
			}
			positions = next
			if len(positions) == 0 {
				break
			}
		}
		out = positions
	case uint8:
		if pos < len(str) && str[pos] == byte(v) {
			out = []int{pos + 1}
		} else {
			out = nil
		}

	default:
		out = nil
	}
	memo[k] = out
	return out
}

/*
// func SpreadOut(row uint8, permut [][]uint8) [][]uint8 {
func SpreadOut(row uint8) [][]uint8 {
	value, Ok := mapp[row]
	if !Ok {
		return nil
	}
	switch value.(type) {
	case [][]uint8:
		fmt.Println(value)
	case []uint8:
		res := make([][]uint8, len(value))
		for _, v := range value {
			res = append(res, SpreadOut(v))
		}
		return res
	case uint8:
		return value
	}
	return [][]uint8{}
}

func Spread(row uint8) {
	return
}

// Cartesian product
/*
func product(){
	return
}
*/
