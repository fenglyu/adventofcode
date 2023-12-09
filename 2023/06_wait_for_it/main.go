package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/fenglyu/adventofcode/util"
)

func str2Arr(s string) []int {
	categories := strings.Split(s, ":")
	sArr := strings.Fields(strings.TrimSpace(strings.TrimSpace(categories[1])))
	arr := make([]int, len(sArr))
	for i, v := range sArr {
		iv, _ := strconv.Atoi(v)
		arr[i] = iv
	}
	return arr
}

func concatNum(s string) int {
	categories := strings.Split(s, ":")
	sArr := strings.Fields(strings.TrimSpace(strings.TrimSpace(categories[1])))
	var res strings.Builder
	for _, v := range sArr {
		fmt.Fprintf(&res, "%s", v)
	}
	r, _ := strconv.Atoi(res.String())
	return r
}

func solveQuadratic(a, b, c float64) (float64, float64) {
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return math.NaN(), math.NaN()
	}

	sqrtDisc := math.Sqrt(discriminant)
	root1 := (-b + sqrtDisc) / (a * 2)
	root2 := (-b - sqrtDisc) / (a * 2)
	if root1 > root2 {
		return root2, root1
	}
	return root1, root2
}
func countWays(a, b, c float64) int {
	root1, root2 := solveQuadratic(a, b, c)
	start := int(math.Floor(root1))
	end := int(math.Ceil(root2))
	return end - start - 1
}

func main() {
	startTime := time.Now()
	raw := util.ParseBasedOnEachLine()

	td := make([][]int, len(raw))
	p2 := make([]int, 0)
	for i, v := range raw {
		td[i] = str2Arr(v)
		p2 = append(p2, concatNum(v))
	}

	power := 1
	for i := 0; i < len(td[0]); i++ {
		a, b, c := float64(1), -float64(td[0][i]), float64(td[1][i])
		numIntegers := countWays(a, b, c)
		//fmt.Println(start, end, root1, root2, numIntegers)
		power = power * numIntegers
	}

	fmt.Println("Problem 1: ", power)
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Println("Execution time: ", duration)
	a, b, c := float64(1), -float64(p2[0]), float64(p2[1])
	fmt.Println("Problem 2: ", countWays(a, b, c))
}
