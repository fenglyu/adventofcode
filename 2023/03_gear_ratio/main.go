package main

import (
	"fmt"

	"github.com/fenglyu/adventofcode/util"
)

var raw [][]uint8

func main() {
	raw = util.ParseIntobyteArray("input")
	gearRatio := make(map[int]string)

	sum := 0
	for i := 0; i < len(raw); i++ {
		j := 0
		for j < len(raw[i]) {
			num := 0
			flag := false
			x, y := 0, 0
			f := false
			var c uint8
			brake := true
			for j < len(raw[i]) && raw[i][j] >= '0' && raw[i][j] <= '9' {
				num = num*10 + int(raw[i][j]-'0')
				if brake {
					f, x, y, c = partNum(i, j)
				}
				if f {
					flag = true
					brake = false
				}
				j++
			}
			if flag {
				//fmt.Println(num)
				sum += num
				if c == '*' {
					gearRatio[num] = fmt.Sprintf("%d_%d", x, y)
				}

			}
			j++
		}
	}
	fmt.Println("Problem 1:", sum)

	fmt.Println(gearRatio)

	reverseMap := make(map[string][]int)
	for k, v := range gearRatio {
		if a, Ok := reverseMap[v]; Ok {
			a = append(a, k)
			reverseMap[v] = a
		} else {
			reverseMap[v] = []int{k}
		}
	}
	fmt.Println(reverseMap)
	var ratio uint64 = 0
	for _, v := range reverseMap {
		if len(v) == 2 {
			power := v[0] * v[1]
			ratio += uint64(power)
		}
	}
	fmt.Println("Problem 2:", ratio)

}

func partNum(i, j int) (bool, int, int, uint8) {
	direction := [][]int{{-1, 0}, {-1, -1}, {0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}}
	flag := false
	x, y := 0, 0
	var c uint8
	for k := 0; k < len(direction); k++ {
		xset, yset := direction[k][0], direction[k][1]
		if i+xset < 0 || i+xset >= len(raw[i]) || j+yset < 0 || j+yset >= len(raw) {
			continue
		}
		v := raw[i+xset][j+yset]
		switch v {
		case '*':
			flag = true
			x = i + xset
			y = j + yset
			c = '*'
		case '#', '$', '+', '%', '@', '/', '&', '!', '^', '(', ')', '-', '=', '[', ']', '?', '\\', '|':
			flag = true
			c = '#'
		case '.':
			continue
		}
	}

	return flag, x, y, c
}
