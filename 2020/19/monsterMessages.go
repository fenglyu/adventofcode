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

//func SpreadOut(row uint8, permut [][]uint8) [][]uint8 {
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
