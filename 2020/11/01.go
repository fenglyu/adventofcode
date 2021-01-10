package main

import (
	"fmt"

	"github.com/fenglyu/adventofcode/util"
)

type rule func([][]uint8, int, int) bool

type round func([][]uint8, rule, uint8) int

func main() {
	part1 := util.ParseIntobyteArray("")
	part(part1, 1, roundN, rule1, rule2)
	part2 := util.ParseIntobyteArray("")
	part(part2, 2, roundN, rule3, rule4)
}

func part(grid [][]uint8, part int, ro round, ru1 rule, ru2 rule) {

	i := 1
	for i <= 1000 {
		changed := 0
		if i%2 != 0 {
			changed = ro(grid, ru1, '#')
		} else if i%2 == 0 {
			changed = ro(grid, ru2, 'L')
		}
		//fmt.Printf("round %d, Changed %d\n", i, changed)
		//displayGrid(grid)
		if changed == 0 {
			fmt.Printf("part %d break at round %d, Changed %d\n", part, i, changed)
			break
		}
		i++
	}

	occupied := countOccupied(grid)
	fmt.Printf("part %d: %d\n", part, occupied)
}

func displayGrid(grid [][]uint8) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			fmt.Printf("%c", grid[i][j])
		}
		fmt.Println("")
	}
}

func countOccupied(grid [][]uint8) int {
	sum := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '#' {
				sum++
			}
		}
	}
	return sum
}

func adjacentSeats(grid [][]uint8, i int, j int) int {
	sum := 0

	vector := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	length := len(vector)
	for k := 0; k < length; k++ {
		row := i + vector[k][0]
		col := j + vector[k][1]
		if row < 0 || col < 0 || row > len(grid)-1 || col > len(grid[0])-1 {
			continue
		}
		//fmt.Println(row, col)
		r := grid[row][col]
		if r == '#' {
			sum++
		}
	}
	return sum
}

func rule1(grid [][]uint8, i int, j int) bool {
	if grid[i][j] == 'L' && adjacentSeats(grid, i, j) == 0 {
		//grid[i][j] = '#'
		return true
	}
	return false
}

func rule2(grid [][]uint8, i int, j int) bool {
	if grid[i][j] == '#' && adjacentSeats(grid, i, j) >= 4 {
		//grid[i][j] = 'L'
		return true
	}
	return false
}

func roundN(grid [][]uint8, fn rule, target uint8) int {
	changed := 0
	record := make([][]int, 0)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			res := fn(grid, i, j)
			if res {
				changed++
				record = append(record, []int{i, j})
			}
			//fmt.Println(res1, res2, changed)
		}
	}

	for i := 0; i < len(record); i++ {
		grid[record[i][0]][record[i][1]] = target
	}
	return changed
}

func rule3(grid [][]uint8, i int, j int) bool {
	if grid[i][j] == 'L' && seeOccupied(grid, i, j) == 0 {
		return true
	}
	return false
}

func rule4(grid [][]uint8, i int, j int) bool {
	if grid[i][j] == '#' && seeOccupied(grid, i, j) >= 5 {
		return true
	}
	return false
}

func seeOccupied(grid [][]uint8, i int, j int) int {
	sum := 0
	// the vector used to find the next seat's postion for checking
	vector := [][]int{{0, -1}, {-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}}
	length := len(vector)
	for k := 0; k < length; k++ {
		q, p := i, j
		for {
			row := q + vector[k][0]
			col := p + vector[k][1]
			if row < 0 || col < 0 || row > len(grid)-1 || col > len(grid[0])-1 {
				break
			}
			if grid[row][col] == 'L' {
				break
			} else if grid[row][col] == '#' {
				sum++
				//q, p = row, col
				// we find a set in this direction, move on to another direction
				break
			} else if grid[row][col] == '.' {
				q, p = row, col
			}
		}
	}

	return sum
}
