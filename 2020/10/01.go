package main

import (
	"fmt"
	"math/big"
	"sort"
	"strconv"

	"github.com/fenglyu/adventofcode/util"
)

type gap struct {
	l    int
	r    int
	diff int
}

func (g *gap) String() string {
	return fmt.Sprintf("l %d, r %d, diff %d", g.l, g.r, g.diff)
}

func main() {

	report := util.ParseBasedOnEachLine()
	length := len(report)

	//fmt.Println(length)
	temp := make([]int, length)
	jolts := make([]int, length+2)
	for i := 0; i < length; i++ {
		v, _ := strconv.Atoi(report[i])
		temp[i] = v
	}
	sort.Ints(temp)
	copy(jolts[1:length+1], temp)

	jolts[0] = 0
	jolts[length+1] = temp[length-1] + 3

	calc := make(map[int]int, 0)
	for i := 1; i <= 3; i++ {
		calc[i] = 0
	}

	gaps := make([]*gap, 0)
	i := 1
	for i < length+2 {
		diff := jolts[i] - jolts[i-1]
		calc[diff]++
		if diff == 3 {
			gaps = append(gaps, &gap{l: i - 1, r: i, diff: diff})
		}
		i++
	}

	multiplied := 1
	for _, v := range calc {
		if v != 0 {
			multiplied *= v
		}
	}

	fmt.Println("part1: ", multiplied)

	distNums := make([]int, 0)
	for j := 0; j < len(gaps); j++ {
		if j == 0 && gaps[j].l != 0 {
			nums := distinctWays(0, gaps[j].l)
			distNums = append(distNums, nums)
			continue
		} else if j == len(gaps)-1 && gaps[j].r != length+1 {
			nums := distinctWays(gaps[j].r, length+1)
			distNums = append(distNums, nums)
			break
		} else {
			nums := distinctWays(gaps[j-1].r, gaps[j].l)
			distNums = append(distNums, nums)
		}
	}

	mul := big.NewInt(1)
	for _, v := range distNums {
		bv := big.NewInt(int64(v))
		mul.Mul(bv, mul)
	}
	fmt.Printf("part2: %v\n", mul)
}

func distinctWays(l int, r int) int {
	dist := r - l
	//fmt.Printf("%d - %d = %d\n", r, l, dist)
	dic := make(map[int]int)
	dic[0] = 1
	dic[1] = 1
	dic[2] = 2
	dic[3] = 4
	dic[4] = 7
	return dic[dist]
}
