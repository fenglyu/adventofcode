package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/fenglyu/adventofcode/util"
)

func main() {
	raw := util.ParseBasedOnEachLine()

	match := make(map[int]int)
	sum := 0
	for _, v := range raw {
		winning, having := mapset.NewSet[int](), mapset.NewSet[int]()

		card := strings.Split(v, ": ")
		cstr := strings.Fields(card[0])
		cNo, _ := strconv.Atoi(cstr[1])
		nums := strings.Split(strings.TrimSpace(card[1]), " | ")
		//winning = mapset.NewSet[]()
		for _, numStr := range strings.Fields(strings.TrimSpace(nums[0])) {
			var num int
			fmt.Sscanf(numStr, "%d", &num)
			winning.Add(num)
		}
		for _, numStr := range strings.Fields(strings.TrimSpace(nums[1])) {
			var num int
			fmt.Sscanf(numStr, "%d", &num)
			having.Add(num)
		}

		//fmt.Println("Winning Set:", strings.TrimSpace(nums[0]), winning)
		//fmt.Println("Having Set:", having)
		points := winning.Intersect(having).Cardinality()
		match[cNo] = points
		sum += int(math.Pow(2, float64(points-1)))
	}
	fmt.Println("Problem 1: ", sum)
	//fmt.Println(matchingNo(0, match))
	dp := make([]int, len(raw)+1)

	for i := 1; i < len(dp); i++ {
		dp[i] = 1 //+ match[i]
	}

	dp[len(dp)-1] = 1
	for i := len(dp) - 1; i > 0; i-- {
		m := match[i]
		if m == 0 {
			dp[i] = 1
		} else {
			for j := i + 1; j < len(dp) && j <= i+m; j++ {
				dp[i] += dp[j]
			}
		}
	}
	res := 0
	for _, v := range dp {
		res += v
	}
	fmt.Println("Problem 2:", res)
}

func matchingNo(n int, match map[int]int) int {
	// itself
	count := 1
	if v, Ok := match[n]; Ok && v == 0 {
		return count
	} else {
		for i := n; i <= n+v; i++ {
			count += matchingNo(i, match)
		}
	}
	return count
}
