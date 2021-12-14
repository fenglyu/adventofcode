package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/fenglyu/adventofcode/util"
)

const MIN = 0.000001

type coordinates struct {
	x, y int
}

func (o *coordinates) String() string {
	return fmt.Sprintf("[%d, %d]", o.x, o.y)
}
func NewCoordinates(s string) *coordinates {
	r := strings.Split(s, ",")
	x, _ := strconv.Atoi(r[0])
	y, _ := strconv.Atoi(r[1])
	return &coordinates{x, y}
}

type segment struct {
	f, t *coordinates
}

func (c *segment) String() string {
	return fmt.Sprintf("f: %v, t: %v", c.f, c.t)
}

func NewSegment(s string) *segment {
	r := strings.Split(s, " -> ")
	return &segment{NewCoordinates(r[0]), NewCoordinates(r[1])}
}

// only consider horizontal and vertical lines
func (c *segment) isHOrV() bool {
	return c.f.x == c.t.x || c.f.y == c.t.y
}

// the lines in your list will only ever be horizontal, vertical, or a diagonal line at exactly 45 degrees
func (c *segment) isDiagonal() bool {
	//if c.isHOrV() {
	//	return true
	//}
	// 45 degress
	k := float64(c.t.y-c.f.y) / float64(c.t.x-c.f.x)
	if IsEqual(k, 1.0) || IsEqual(k, -1.0) {
		return true
	}
	return false
}

func IsEqual(f1, f2 float64) bool {
	if f1 > f2 {
		return math.Dim(f1, f2) < MIN
	} else {
		return math.Dim(f2, f1) < MIN
	}
}

func (c *segment) coverPointsPart2(dig [][]int) {

	if !c.isHOrV() && !c.isDiagonal() {
		return
	}

	if c.isHOrV() {
		c.coverPoints(dig)
	} else if c.isDiagonal() {
		xIncr, yIncr := 1, 1
		if c.f.x > c.t.x {
			xIncr = -1
		}
		if c.f.y > c.t.y {
			yIncr = -1
		}

		i, j := c.f.x, c.f.y
		for {
			if xIncr > 0 && i > c.t.x {
				break
			}
			if xIncr < 0 && i < c.t.x {
				break
			}
			if yIncr > 0 && j > c.t.y {
				break
			}
			if yIncr < 0 && j < c.t.y {
				break
			}

			dig[i][j] += 1
			i, j = i+xIncr, j+yIncr
		}
	}
}

func (c *segment) coverPoints(dig [][]int) {
	if !c.isHOrV() {
		return
	}

	incr := 1
	if c.f.x > c.t.x || c.f.y > c.t.y {
		incr = -1
	}

	for i := c.f.x; ; i += incr {
		if incr > 0 && i > c.t.x {
			break
		}
		if incr < 0 && i < c.t.x {
			break
		}
		for j := c.f.y; ; j += incr {
			if incr > 0 && j > c.t.y {
				break
			}
			if incr < 0 && j < c.t.y {
				break
			}
			dig[i][j] += 1
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	report := util.ParseBasedOnEachLine()

	//	fmt.Println(report)

	segs := make([]*segment, len(report))
	for i, v := range report {
		segs[i] = NewSegment(v)
	}

	hSize, vSize := 0, 0
	for _, v := range segs {
		hSize = max(hSize, max(v.f.x, v.t.x))
		vSize = max(vSize, max(v.f.y, v.t.y))
	}

	//fmt.Println("h, v", hSize, vSize)

	diagram := make([][]int, hSize+1)
	for i := range diagram {
		row := make([]int, vSize+1)
		for j := range row {
			row[j] = 0
		}
		diagram[i] = row
	}

	for _, v := range segs {
		v.coverPoints(diagram)
	}

	part1 := overlap(diagram, 2)
	fmt.Println("part 1: ", part1)

	diagram2 := make([][]int, hSize+1)
	for i := range diagram2 {
		row := make([]int, vSize+1)
		for j := range row {
			row[j] = 0
		}
		diagram2[i] = row
	}

	for _, v := range segs {
		//	fmt.Println(v)
		v.coverPointsPart2(diagram2)
	}

	part2 := overlap(diagram2, 2)
	//util.PrintMatrix(diagram2)
	fmt.Println("part 2: ", part2)
}

func overlap(diag [][]int, least int) int {
	sum := 0
	for i := 0; i < len(diag); i++ {
		for j := 0; j < len(diag[0]); j++ {
			if diag[i][j] >= least {
				sum++
			}
		}
	}

	return sum
}
