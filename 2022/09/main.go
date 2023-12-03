package main

import (
	"fmt"
	"sort"

	"github.com/fenglyu/adventofcode/util"
)

var matrix [][]uint8
var tracker [][]uint8
var FLAG uint8 = (1 << 8) - 1

func main() {
	report := util.ParseBasedOnEachLine()
	//	fmt.Println(report)
	matrix = make([][]uint8, len(report))

	for i := range matrix {
		matrix[i] = []byte(report[i])
	}

	tracker = make([][]uint8, len(matrix))
	for i := 0; i < len(matrix); i++ {
		row := make([]uint8, len(matrix[0]))
		for j := 0; j < len(matrix[0]); j++ {
			row[j] = 0
		}
		tracker[i] = row
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			lowest(i, j)
		}
	}

	sum := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if tracker[i][j] == FLAG {
				sum += int(matrix[i][j]-'0') + 1
			}
		}
	}

	fmt.Println("part 1 :", sum)

	for i := 0; i < len(matrix); i++ {
		row := make([]uint8, len(matrix[0]))
		for j := 0; j < len(matrix[0]); j++ {
			row[j] = 0
		}
		tracker[i] = row
	}

	//fmt.Println(tracker)

	res := make([]int, 0)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			c := basin(i, j)
			res = append(res, c)
		}
	}

	sort.Ints(res)

	mul := func(arr []int) int {
		sum := 1
		for i := range arr {
			sum *= arr[i]
		}
		return sum
	}
	fmt.Println("part2 :", mul(res[len(res)-3:]))
}

func basin(i, j int) int {

	//fmt.Println("basin : ", i, j)
	if matrix[i][j] >= '9' {
		return 0
	}
	if tracker[i][j] == FLAG {
		return 0
	}

	tracker[i][j] = FLAG
	count := 1

	direction := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	//	s := set.NewSet()
	for _, v := range direction {
		// board boundary
		if v[0]+i < 0 || v[1]+j < 0 || v[0]+i >= len(matrix) || v[1]+j >= len(matrix[0]) {
			continue
		}

		if tracker[v[0]+i][v[1]+j] != FLAG && matrix[v[0]+i][v[1]+j] < '9' {
			count += basin(v[0]+i, v[1]+j)
		}
	}

	return count
}

func lowest(i, j int) {
	direction := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	for _, v := range direction {
		if v[0]+i < 0 || v[1]+j < 0 || v[0]+i >= len(matrix) || v[1]+j >= len(matrix[0]) {
			continue
		}

		if matrix[i][j] > matrix[i+v[0]][j+v[1]] {
			tracker[i+v[0]][j+v[1]] = FLAG
			tracker[i][j] = 0
			lowest(i+v[0], j+v[1])
		}
	}
}
