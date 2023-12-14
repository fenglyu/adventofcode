package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/fenglyu/adventofcode/util"
)

func main() {
	startTime := time.Now()
	raw := util.ParseBasedOnEachLine()

	count := 0
	c2 := 0
	for _, v := range raw {
		arr := genArray(v)
		s, s2 := prediction(arr)
		count += s
		c2 += s2
		//fmt.Println("s: ", s)
	}

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Println("Execution time: ", duration)
	fmt.Println("problem 1: ", count)
	fmt.Println("problem 2: ", c2)
}

func genArray(s string) []int {
	arr := strings.Fields(strings.TrimSpace(s))
	res := make([]int, len(arr)+1)
	for i, v := range arr {
		val, _ := strconv.Atoi(v)
		res[i] = val
	}
	res[len(arr)] = -1
	return res
}

func initArray(s []int) {
	for i := 0; i < len(s); i++ {
		s[i] = -1
	}
}

func prediction(s []int) (int, int) {
	target := make([][]int, 0)
	target = append(target, s)
	j := 2
	l := -1
	for {
		sum := true
		next := make([]int, len(s))
		initArray(next)
		for i := 0; i < len(s)-j; i++ {
			next[i] = s[i+1] - s[i]
			sum = sum && next[i] == 0
			l = i
		}
		//fmt.Println("next :", next)
		s = next
		target = append(target, next)
		j++
		if sum {
			break
		}
	}

	val := 0
	m, n := len(target)-1, l+1
	target[m][n] = target[m][n-1]
	for m > 0 && n+1 < len(target[0]) {
		target[m-1][n+1] = target[m][n] + target[m-1][n]
		val = target[m-1][n+1]
		m--
		n++
	}

	//displayGrid(target)
	val2 := 0
	m = len(target) - 2
	for m >= 0 {
		val2 = target[m][0] - val2
		m--
	}
	//fmt.Println("val2:", val2)
	return val, val2
}

func displayGrid(grid [][]int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			fmt.Printf("%d ", grid[i][j])
		}
		fmt.Println("")
	}
}
