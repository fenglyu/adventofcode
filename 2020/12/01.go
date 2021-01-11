package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/fenglyu/adventofcode/util"
)

//type move func(uint8, int)

type navi struct {
	action uint8
	value  int
}

func (n *navi) String() string {
	return fmt.Sprintf("action %c, value %d", n.action, n.value)
}

type position struct {
	x         int
	y         int
	direction uint8
}

func (p *position) String() string {
	return fmt.Sprintf("[x, y] [%d, %d] in direction %c", p.x, p.y, p.direction)
}

func (p *position) ManhattanDistance() int {
	return int(math.Abs(float64(p.x))) + int(math.Abs(float64(p.y)))
}

func (p *position) turn(navi *navi) {
	dirs := []uint8{'E', 'N', 'W', 'S'}
	var curDir int
	for j, v := range dirs {
		if v == p.direction {
			curDir = j
			break
		}
	}

	for i := 0; i < navi.value/90; i++ {
		if navi.action == 'L' {
			if curDir == len(dirs)-1 {
				curDir = 0
			} else {
				curDir++
			}
		} else if navi.action == 'R' {
			if curDir == 0 {
				curDir = len(dirs) - 1
			} else {
				curDir--
			}
		}
	}

	//	fmt.Println("postion before turn: ", p)
	p.direction = dirs[curDir]
	//	fmt.Println("postion after turn: ", p)
}

func (p *position) move(navi *navi) *position {
	switch navi.action {
	case 'N':
		p.y = p.y + navi.value
	case 'S':
		p.y = p.y - navi.value
	case 'E':
		p.x = p.x + navi.value
	case 'W':
		p.x = p.x - navi.value
	case 'F':
		navi.action = p.direction
		return p.move(navi)
	case 'L', 'R':
		p.turn(navi)
	}

	return p
}

func main() {

	report := util.ParseBasedOnEachLine()
	//fmt.Println(report)

	navis := make([]*navi, 0)
	for _, v := range report {
		arr := []byte(v)
		val, _ := strconv.Atoi(string(arr[1:]))
		navis = append(navis, &navi{action: arr[0], value: val})
	}
	/*
		for i, n := range navis {
			fmt.Println(i, n)
		}
	*/
	initP := &position{x: 0, y: 0, direction: 'E'}
	for _, n := range navis {
		initP.move(n)
	}

	fmt.Println("part1: ", initP.ManhattanDistance())
}
