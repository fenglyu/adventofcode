package main

import (
	"fmt"

	//set "github.com/deckarep/golang-set"
	"github.com/fenglyu/adventofcode/util"
)

var (
	grid   [][][]uint8
	vector [][]int
)

type point struct {
	x, y, z int
}

func (p *point) String() string {
	return fmt.Sprintf("[%d, %d, %d]", p.x, p.y, p.z)
}

func main() {

	report := util.ParseIntobyteArray("test_input")

	grid = make([][][]uint8, 0)
	grid = append(grid, report)

	//printCubes(grid)

	a := []int{-1, 0, 1}
	npt := util.NextProduct(a, 3)
	vector = make([][]int, 0)
	for {
		np := npt()
		if len(np) == 0 {
			break
		}
		// skip the central cube [0,0,0]
		if np[0] == 0 && np[1] == 0 && np[2] == 0 {
			continue
		}
		nn := make([]int, len(np))
		copy(nn, np)
		vector = append(vector, nn)
	}

	cubes := expandCubes(grid)
	printCubes(cubes)
	/*
		cubes = expandCubes(cubes)
		printCubes(cubes)
		cubes = expandCubes(cubes)
		printCubes(cubes)
	*/
}

func conwayCubes(initGrid [][]int, cycle int) {

}

func expandCubes(grid [][][]uint8) [][][]uint8 {
	z1, x1, y1 := len(grid), len(grid[0]), len(grid[0][0])
	ng := make([][][]uint8, z1+2)
	ng[0], ng[len(ng)-1] = newNegSquare(x1+2, y1+2), newNegSquare(x1+2, y1+2)
	for z := 1; z < len(ng)-1; z++ {
		square := make([][]uint8, x1+2)
		square[0], square[len(square)-1] = newNegLine(y1+2), newNegLine(y1+2)
		for x := 1; x < len(square)-1; x++ {
			l := make([]uint8, y1+2)
			copy(l[1:len(l)-1], grid[z-1][x-1])
			l[0], l[len(l)-1] = '.', '.'
			square[x] = l
		}
		ng[z] = square
	}
	return ng
}

func newNegLine(length int) []uint8 {
	res := make([]uint8, length)
	for i := 0; i < len(res); i++ {
		res[i] = '.'
	}
	return res
}

func newNegSquare(x, y int) [][]uint8 {
	res := make([][]uint8, x)
	for i := 0; i < x; i++ {
		res[i] = newNegLine(y)
	}
	return res
}

func buildAllCubes(grid [][]uint8, cycle int) [][][]uint8 {
	ng := make([][][]uint8, 0)
	for z := 0; z < 1+cycle*2; z++ {
		square := make([][]uint8, 0)
		for x := 0; x < len(grid)+cycle*2; x++ {
			l := make([]uint8, 0)
			for y := 0; y < len(grid[0])+cycle*2; y++ {
				l = append(l, '.')
			}
			square = append(square, l)
		}
		ng = append(ng, square)
	}
	return ng
}

func rule(x int, y int, z int, res [][][]int) [][][]int {
	active := 0
	for _, v := range vector {
		if grid[x+v[0]][y+v[1]][z+v[2]] == '*' {
			active++
		}
	}

	if grid[x][y][z] == '*' && (active == 2 || active == 3) {

	} else if grid[x][y][z] == '.' && active == 3 {

	} else {

	}
	return nil
}

func printCubes(grid [][][]uint8) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			for k := 0; k < len(grid[0][0]); k++ {
				fmt.Printf("%c ", grid[i][j][k])
			}
			fmt.Println("")
		}
		fmt.Println("-----------------------------")
	}
}
