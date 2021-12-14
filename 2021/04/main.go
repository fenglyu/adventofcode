package main

import (
	"fmt"
	"regexp"
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
	for i := range Cmap {
		Cmap[i] = -1
	}

	for i, v := range numbers {
		Cmap[v] = i
	}

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

	/*
		sum := func(arr []int) int {
			s := 0
			for _, v := range arr {
				s += v
			}
			return s
		}
	*/

	hist := make([][][]int, len(report[1:]))
	m, idx := 50, 50
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

	s := sumRestElem(hist[idx], m, numbers)
	//fmt.Println(idx, "board ", numbers[m], numbers[m+1:], s)
	fmt.Println("part 1 :", numbers[m]*s)
}

// choose the winner board, start count from index based on numbers list
func sumRestElem(b [][]int, idx int, nums []int) int {
	sum := 0
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[0]); j++ {
			if b[i][j] > idx {
				sum += nums[b[i][j]]
			}
		}
	}
	return sum
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

	r := regexp.MustCompile("[^\\s]+")
	rows := strings.Split(strings.TrimSpace(bs), "\n")
	for _, row := range rows {
		ro := r.FindAllString(strings.TrimSpace(row), -1)
		br := make([]int, len(ro))
		for j := range ro {
			v, _ := strconv.Atoi(ro[j])
			br[j] = v
		}
		b = append(b, br)
	}

	//printM(b)
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
	//printM(seqB)
	return seqB
}

func printM(b [][]int) {
	fmt.Println("------->")
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[0]); j++ {
			fmt.Printf("%d ", b[i][j])
		}
		fmt.Println("")
	}
}
