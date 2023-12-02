package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/fenglyu/adventofcode/util"
)

type color struct {
	Blue  int
	Green int
	Red   int
}

func (c color) String() string {
	return fmt.Sprintf("%d red %d green %d blue", c.Red, c.Green, c.Blue)
}

func newColor(s string) color {
	co := color{}
	for _, c := range strings.Split(s, ", ") {
		attr := strings.Split(strings.TrimSpace(c), " ")
		num, _ := strconv.Atoi(attr[0])
		switch attr[1] {
		case "green":
			co.Green = num
		case "red":
			co.Red = num
		case "blue":
			co.Blue = num
		}
	}
	return co
}

func (c color) possible() bool {
	return c.Red <= 12 && c.Green <= 13 && c.Blue <= 14
}

func (c color) impossible() bool {
	return !c.possible()
}

func parseGame(s string) (int, []color) {
	// Game 1: 1 green, 6 red, 4 blue; 2 blue, 6 green, 7 red; 3 red, 4 blue, 6 green; 3 green; 3 blue, 2 green, 1 red
	gs := strings.Split(s, ":")
	ids := strings.Split(strings.TrimSpace(gs[0]), " ")
	id, _ := strconv.Atoi(ids[1])

	rounds := strings.Split(strings.TrimSpace(gs[1]), ";")
	colors := make([]color, len(rounds))

	for i, r := range rounds {
		colors[i] = newColor(r)
	}
	return id, colors
}

/*
func min(a, b, c int) int {
	return int(math.Min(float64(a), math.Min(float64(b), float64(c))))
}
*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	rawData := util.ParseBasedOnEachLine()
	//games := make([][]color, len(rawData))

	sum := 0
	res := 0
	for _, v := range rawData {
		id, colors := parseGame(v)
		flag := true
		for _, c := range colors {
			if c.impossible() {
				flag = false
			}
		}
		if flag {
			sum += id
		}

		// problem 2
		minR, minG, minB := math.MinInt, math.MinInt, math.MinInt
		for _, c := range colors {
			minR = max(c.Red, minR)
			minG = max(c.Green, minG)
			minB = max(c.Blue, minB)
		}
		power := minR * minG * minB
		res += power
	}
	fmt.Println("Problem 1:", sum)
	fmt.Println("Problem 2:", res)
}
