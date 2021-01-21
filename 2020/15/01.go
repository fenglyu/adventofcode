package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fenglyu/adventofcode/util"
)

type circleList struct {
	head, tail int
	data       []int
	capacity   int
	count      int
}

func newCL(capacity int) *circleList {
	c := circleList{head: 0, tail: 0, capacity: capacity, count: 0}
	c.data = make([]int, capacity)
	return &c
}

func (c *circleList) add(ele int) {
	if c.count == 0 {
		c.data[c.tail] = ele
	} else {
		c.tail = (c.tail + 1) % c.capacity
		if c.tail == c.head {
			c.head = mod(c.head-1, c.capacity)
		}
		c.data[c.tail] = ele

	}

	if c.count < c.capacity {
		c.count++
	}
}

// return the next `count` elements start from tail in reverse order
func (c *circleList) next(count int) []int {
	if count > c.capacity {
		count = c.capacity
	}

	res := make([]int, count)
	cnt := 0
	l := c.tail
	for cnt < count {
		res[cnt] = c.data[l]
		l = mod(l-1, c.capacity)
		cnt++
	}
	return res
}

// python like mod, -1 % 5 == 4, in go -1 % 5 == -1
func mod(a, b int) int {
	return (a%b + b) % b
}

func main() {

	report := util.ParseBasedOnEachLine()
	turns := make([]int, 0)

	for _, v := range strings.Split(report[0], ",") {
		val, _ := strconv.Atoi(v)
		turns = append(turns, val)
	}
	part2Turns := make([]int, len(turns))
	copy(part2Turns, turns)

	limit := 2020
	turns = startingGame(turns, limit)
	fmt.Println("part 1: ", turns[len(turns)-2])

	limit = 30000000
	part2Turns = startingGame(part2Turns, limit)
	fmt.Println("part 2: ", part2Turns[len(part2Turns)-2])

	/*
		cl := newCL(5)
		a := []int{1, 2, 3, 4, 5, 6, 7, 8}
		for _, v := range a {
			cl.add(v)
		}

		fmt.Println(cl.data)
		fmt.Println(cl.next(2))
		fmt.Println(cl.next(3))
	*/

}

func startingGame(turns []int, limit int) []int {
	turn := 0
	freq := make(map[int]*circleList)
	for turn < limit {
		spokenNum := turns[turn]
		length := len(turns) - 1
		if _, Ok := freq[spokenNum]; Ok {
			cl := freq[spokenNum]
			cl.add(turn)
			if turn+1 > length {
				// prepare the next turn spoken number here
				pop := cl.next(2)
				// next spoken number difference of last two turns
				turns = append(turns, pop[0]-pop[1])
			}
		} else {
			cl := newCL(5)
			cl.add(turn)
			freq[spokenNum] = cl
			if turn+1 > length {
				turns = append(turns, 0)
			}
		}
		turn++
	}
	return turns
}
