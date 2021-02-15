package main

import (
	"fmt"

	"github.com/fenglyu/adventofcode/util"
)

var (
	grid [][][]uint8
)

type point struct {
	x, y, z int
}

func (p *point) String() string {
	return fmt.Sprintf("[%d, %d, %d]", p.x, p.y, p.z)
}

// https://github.com/ryanbressler/numpy2go
func main() {

	report := util.ParseIntobyteArray("input")

	grid = make([][][]uint8, 0)
	grid = append(grid, report)

	vec := util.ProductGen([]int{-1, 0, 1}, 3, true)

	part1 := conwayCubesSolution1(grid, vec, 6)
	//printCubes(grid)
	fmt.Println("part1: ", part1)

	hypeVec := util.ProductGen([]int{-1, 0, 1}, 4, true)
	//fmt.Println("hypeVec: ", len(hypeVec), hypeVec)

	hypeCube := make([][]*square, 0)
	g := make([]*square, 0)
	s := &square{sqz: report}
	g = append(g, s)
	hypeCube = append(hypeCube, g)

	hc := newHypeCube(hypeCube)
	//fmt.Println((hc.c[0][0]).sqz)
	fmt.Println("conwayCubes: ", hc.conwayCubes(hypeVec))
}

func conwayCubesSolution1(g [][][]uint8, vector [][]int, cycle int) int {
	c := 0
	for c < cycle {
		grid = expandCubes(g)
		other := expandCubes(g)
		z1, x1, y1 := len(grid), len(grid[0]), len(grid[0][0])
		for i := 0; i < z1; i++ {
			for j := 0; j < x1; j++ {
				for k := 0; k < y1; k++ {
					grid = rule(vector, other, i, j, k, grid)
				}
			}
		}
		g = grid
		c++
	}

	return countActive(grid)
}

func expandCubes(g [][][]uint8) [][][]uint8 {
	z1, x1, y1 := len(g), len(g[0]), len(g[0][0])
	ng := make([][][]uint8, z1+2)
	ng[0], ng[len(ng)-1] = newNegSquare(x1+2, y1+2), newNegSquare(x1+2, y1+2)
	for z := 1; z < len(ng)-1; z++ {
		square := make([][]uint8, x1+2)
		square[0], square[len(square)-1] = newNegLine(y1+2), newNegLine(y1+2)
		for x := 1; x < len(square)-1; x++ {
			l := make([]uint8, y1+2)
			copy(l[1:len(l)-1], g[z-1][x-1])
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

func rule(vector [][]int, g [][][]uint8, z int, x int, y int, res [][][]uint8) [][][]uint8 {
	active := 0
	for _, v := range vector {
		z1, x1, y1 := z+v[0], x+v[1], y+v[2]
		if z1 < 0 || z1 >= len(grid) || x1 < 0 || x1 >= len(grid[0]) || y1 < 0 || y1 >= len(grid[0][0]) {
			continue
		}
		if g[z1][x1][y1] == '#' {
			active++
		}
	}

	if g[z][x][y] == '#' && (active == 2 || active == 3) {
		res[z][x][y] = '#'
	} else if g[z][x][y] == '.' && active == 3 {
		res[z][x][y] = '#'
	} else {
		res[z][x][y] = '.'
	}

	return res
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

func countActive(grid [][][]uint8) int {
	active := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			for k := 0; k < len(grid[0][0]); k++ {
				if grid[i][j][k] == '#' {
					active++
				}
			}
		}
	}
	return active
}

type square struct {
	sqz [][]uint8
}

// 4D cube
type hypeCube struct {
	c [][]*square
}

func (sq *square) Stirng() string {
	return fmt.Sprintf("%v", sq.sqz)
}

func (h *hypeCube) Stirng() string {
	return fmt.Sprintf("%v", h.c)
}

func newSquare(x, y int) *square {
	res := make([][]uint8, x)
	for i := 0; i < x; i++ {
		res[i] = newNegLine(y)
	}
	return &square{sqz: res}
}

func newHypeCube(hc [][]*square) *hypeCube {
	return &hypeCube{c: hc}
}

func (sq *square) expandSquare() *square {
	z1, x1 := len(sq.sqz), len(sq.sqz[0])
	s := newSquare(z1+2, x1+2)
	s.sqz[0], s.sqz[len(s.sqz)-1] = newNegLine(x1+2), newNegLine(x1+2)
	for x := 1; x < len(s.sqz)-1; x++ {
		l := make([]uint8, x1+2)
		copy(l[1:len(l)-1], sq.sqz[x-1])
		l[0], l[len(l)-1] = '.', '.'
		s.sqz[x] = l
	}
	return s
}

func newLineSquare(length, x, y int) []*square {
	res := make([]*square, length)
	for i := 0; i < len(res); i++ {
		res[i] = newSquare(x, y)
	}
	return res
}

func (h *hypeCube) expandHypeCube() *hypeCube {
	z1, w1 := len(h.c), len(h.c[0])
	t := h.c[0][0]
	x1, y1 := len(t.sqz), len(t.sqz[0])
	ng := make([][]*square, z1+2)
	s := newHypeCube(ng)
	s.c[0], s.c[len(s.c)-1] = newLineSquare(w1+2, x1+2, y1+2), newLineSquare(w1+2, x1+2, y1+2)
	for x := 1; x < len(s.c)-1; x++ {
		l := make([]*square, w1+2)
		//copy(l[1:len(l)-1], h.c[x-1])
		for j := 0; j < len(l); j++ {
			if j == 0 || j == len(l)-1 {
				l[j] = newSquare(x1+2, y1+2)
			} else {
				t := h.c[x-1][j-1]
				l[j] = t.expandSquare()
			}
		}
		s.c[x] = l
	}
	return s
}

func (h *hypeCube) rule(vector [][]int, z int, w int, x int, y int, o *hypeCube) *hypeCube {
	active := 0
	for _, v := range vector {
		z1, w1 := z+v[0], w+v[1]
		if z1 < 0 || z1 >= len(h.c) || w1 < 0 || w1 >= len(h.c[0]) {
			continue
		}
		x1, y1 := x+v[2], y+v[3]
		t := h.c[z1][w1]
		if x1 < 0 || x1 >= len(t.sqz) || y1 < 0 || y1 >= len(t.sqz[0]) {
			continue
		}

		if t.sqz[x1][y1] == '#' {
			active++
		}
	}

	sq, oq := h.c[z][w], o.c[z][w]
	if sq.sqz[x][y] == '#' && (active == 2 || active == 3) {
		oq.sqz[x][y] = '#'
	} else if sq.sqz[x][y] == '.' && active == 3 {
		oq.sqz[x][y] = '#'
	} else {
		oq.sqz[x][y] = '.'
	}
	o.c[z][w] = oq
	return o
}

func (h *hypeCube) conwayCubes(vec [][]int) int {
	//fmt.Println("h: ", (h.c[0][0]).sqz)
	cycle := 1
	for cycle <= 6 {
		g, o := h.expandHypeCube(), h.expandHypeCube()
		z1, x1 := len(g.c), len(g.c[0])
		for i := 0; i < z1; i++ {
			for j := 0; j < x1; j++ {
				// looking into the square
				t := g.c[i][j]
				for m := 0; m < len(t.sqz); m++ {
					for n := 0; n < len(t.sqz[0]); n++ {
						g.rule(vec, i, j, m, n, o)
					}
				}
			}
		}
		h = o
		//o.printCubes()
		cycle++
	}

	return h.countActive()
}

func (h *hypeCube) countActive() int {
	active := 0

	z1, x1 := len(h.c), len(h.c[0])
	for i := 0; i < z1; i++ {
		for j := 0; j < x1; j++ {
			// looking into the square
			t := h.c[i][j]
			for m := 0; m < len(t.sqz); m++ {
				for n := 0; n < len(t.sqz[0]); n++ {
					if t.sqz[m][n] == '#' {
						active++
					}
				}
			}
		}
	}
	return active
}

func (h *hypeCube) printCubes() {
	z1, w1 := len(h.c), len(h.c[0])
	for i := 0; i < z1; i++ {
		for j := 0; j < w1; j++ {
			fmt.Printf("z=%d, w=%d\n", j-w1/2, i-z1/2)
			t := h.c[i][j]
			x1, y1 := len(t.sqz), len(t.sqz[0])
			for m := 0; m < x1; m++ {
				for n := 0; n < y1; n++ {
					fmt.Printf("%c ", t.sqz[m][n])
				}
				fmt.Println("")
			}
		}
	}
}
