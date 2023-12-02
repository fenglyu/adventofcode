package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/fenglyu/adventofcode/util"
)

func main() {

	file := "input"
	input := util.ParseIntobyteArray(file)
	sum := 0
	for i := 0; i < len(input); i++ {
		j, k := 0, len(input[i])-1
		for j < k && (input[i][j] < '0' || input[i][j] > '9') {
			j++
		}
		for j < k && (input[i][k] < '0' || input[i][k] > '9') {
			k--
		}
		if j > k {
			break
		}
		res := int(input[i][j]-'0')*10 + int(input[i][k]-'0')
		sum += res
	}
	fmt.Println("Problem 1: ", sum)

	input2 := util.ParseBasedOnEachLine()

	sum2 := 0
	for _, v := range input2 {
		l, r := findNums(v)

		sum2 += l*10 + r
	}
	fmt.Println("Problem 2 ", sum2)
}

func findNums(s string) (int, int) {
	digits := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
	}

	f, l := math.MaxInt, math.MinInt
	fres, lres := -1, -1
	for k, _ := range digits {
		fi, si := strings.Index(s, k), strings.LastIndex(s, k)
		if fi >= 0 && fi < f {
			f = fi
			fres = digits[k]
			//	fmt.Println(k, f, fres)
		}

		if si >= 0 && si+len(k) > l {
			l = si + len(k)
			lres = digits[k]
		}
	}
	//	fmt.Println(f, l)
	return fres, lres
}
