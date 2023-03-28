package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fenglyu/adventofcode/util"
)

var mapp map[uint8]any

func main() {

	report := util.ParseBasedOnEmptyLine()
	//fmt.Println(report, len(report))

	mapp = make(map[uint8]any)
	for _, v := range strings.Split(report[0], "\n") {
		r := strings.Split(v, ": ")
		idInt32, _ := strconv.Atoi(r[0])
		idx := uint8(idInt32)
		if strings.Contains(r[1], "|") {
			// 93: 57 68 | 12 110
			r := strings.Split(r[1], " | ")
			n := [][]uint8{concatBytes(r[0]), concatBytes(r[1])}
			mapp[idx] = n
		} else if strings.Contains(r[1], "\"") {
			// 12: "a"
			v := []byte(strings.Trim(r[1], "\""))
			//mapp[idx] = bytes.Trim([]byte(r[1]), "\"")
			mapp[idx] = uint8(v[0])
		} else {
			// 0: 8 11
			mapp[idx] = concatBytes(r[1])
		}
	}

	fmt.Println(mapp)
	fmt.Println(Rec([]byte("ababbb"), 0, 0))
}

func concatBytes(res string) []uint8 {
	a := strings.Split(res, " ")
	u8a := make([]uint8, len(a))
	for i := 0; i < len(a); i++ {
		ia, _ := strconv.Atoi(a[i])
		u8a[i] = uint8(ia)
	}
	return u8a
}

func Rec(str []byte, row uint8, idx int) (bool, int) {
	value, Ok := mapp[row]
	if !Ok {
		return false, idx
	}
	switch value.(type) {
	case [][]uint8:
		rules := value.([][]uint8)
		r0, _ := Rec(str, rules[0][0], idx)
		r1, _ := Rec(str, rules[0][1], idx+1)
		r2, _ := Rec(str, rules[1][0], idx)
		r3, _ := Rec(str, rules[1][1], idx+1)
		return (r0 && r1) || (r2 && r3), idx + 2
	case []uint8:
		flag := true
		for _, v := range value.([]uint8) {
			res, i := Rec(str, v, idx)
			idx = i
			if !res {
				flag = false
				break
			}
		}
		return flag, idx
	case uint8:
		return str[idx] == value.(uint8), idx + 1
	}
	return true, idx
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
