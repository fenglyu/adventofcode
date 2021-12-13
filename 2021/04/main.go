package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fenglyu/adventofcode/util"
)

// all numbers in any row or any column of a board are marked
var RowColSeq [][]int

func main() {
	report := util.ParseBasedOnEmptyLine()

	//fmt.Println(report[0], report[1], report[2], report[3])

	maxI := 0
	numbers := make([]int, 0)
	for _, v := range strings.Split(report[0], ",") {
		r, _ := strconv.Atoi(v)
		if r > maxI {
			maxI = r
		}
		numbers = append(numbers, r)
	}

	Cmap := make([]int, maxI+1)
	for i, v := range numbers {
		Cmap[v] = i
	}
	//	fmt.Println(numbers)

	max := func(arr []int) int {
		idx := 0
		for i := range arr {
			if arr[i] > arr[idx] {
				idx = i
			}
		}
		return arr[idx]
	}

	//min := func(arr []int) (int, int) {
	min := func(arr []int) int {
		idx := 0
		for i := range arr {
			if arr[i] < arr[idx] {
				idx = i
			}
		}
		return arr[idx]
	}

	sum := func(arr []int) int {
		s := 0
		for _, v := range arr {
			s += v
		}
		return s
	}

	hist := make([][][]int, len(report[1:]))
	m, idx := 0, 0
	for i, board := range report[1:] {
		b := generateBoard(board)
		seqB := convertBoard(Cmap, b)
		hist[i] = seqB
		val := boardEaliest(max, min, seqB)
		if val < m {
			idx = i
			m = val
		}
	}

	part1 := Cmap[idx] * sum(numbers[idx+1:])

	fmt.Println(Cmap[idx], numbers[Cmap[idx]+1:])
	fmt.Println("part 1 :", part1)
}

type mm func([]int) int

func boardEaliest(max mm, min mm, b [][]int) int {
	//max := func(arr []int) (int, int) {

	smallest := make([]int, len(b))
	for i := 0; i < len(b); i++ {
		smallest[i] = max(b[i])
	}
	v := min(smallest)

	return v
}

func generateBoard(bs string) [][]int {
	b := make([][]int, 0)

	rows := strings.Split(bs, "\n")
	for _, row := range rows {
		r := strings.Split(row, " ")
		br := make([]int, len(r))
		for j := range r {
			v, _ := strconv.Atoi(r[j])
			br[j] = v
		}
		b = append(b, br)
	}
	return b
}

func convertBoard(m []int, b [][]int) [][]int {
	seqB := make([][]int, len(b))

	for i := 0; i < len(b); i++ {
		r := make([]int, len(b[0]))
		for j := 0; j < len(b[0]); j++ {
			r[j] = m[b[i][j]]
		}
		seqB[i] = r
	}
	return seqB
}
