package main

import (
	"fmt"

	"github.com/fenglyu/adventofcode/util"
)

var dict [][]int = [][]int{{0, -1}, {-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}}

func main() {
	report := util.ParseIntobyteArray("")

	//Display(report)
	flash := 0
	for i := 0; i < 100; i++ {
		flash += Step(report)
		//fmt.Printf("After step %d:\n", i+1)
		//Display(report)
	}

	fmt.Println("part 1 :", flash)

	report = util.ParseIntobyteArray("")
	for i := 0; ; i++ {
		Step(report)
		if synchronizing(report) {
			fmt.Println("part 2: ", i+1)
			break
		}
	}
}

func synchronizing(grid [][]uint8) bool {

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != '0' {
				return false
			}
		}
	}

	return true
}

func Step(grid [][]uint8) int {
	flash := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		row := make([]int, len(grid[0]))
		for j := 0; j < len(grid[0]); j++ {
			row[j] = 0
		}
		flash[i] = row
	}

	return Flash(grid, flash)
}

func Flash(grid [][]uint8, flash [][]int) int {
	count := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			count += task(grid, flash, i, j)
		}
	}

	return count
}

func task(grid [][]uint8, flash [][]int, i, j int) int {
	// recrusive call with quit fast
	if flash[i][j] > 0 {
		return 0
	}
	count := 0
	grid[i][j]++

	//Display(grid)
	if grid[i][j] <= '9' {
		return 0
	}

	if grid[i][j] > '9' {
		count++
		grid[i][j] = '0'
		flash[i][j] = 1
	}

	// either when grid[i][j] == '9' or when grid[i][j] has been flashed already in this step
	if grid[i][j] > '9' || flash[i][j] == 1 {
		for _, v := range dict {
			if i+v[0] < 0 || j+v[1] < 0 || i+v[0] > len(grid)-1 || j+v[1] > len(grid[0])-1 {
				continue
			}

			if flash[i+v[0]][j+v[1]] == 1 {
				continue
			}

			count += task(grid, flash, i+v[0], j+v[1])
		}
	}
	return count
}

func Display(grid [][]uint8) {
	fmt.Println("==============>")
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			fmt.Printf("%c ", grid[i][j])
		}
		fmt.Println("")
	}
}
