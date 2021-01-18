package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/fenglyu/adventofcode/util"
)

func main() {

	report := util.ParseBasedOnEachLine()
	earlistTimeStamp, _ := strconv.Atoi(report[0])
	buses := make([]int, 0)
	part2 := make([]int, 0)
	for _, v := range strings.Split(report[1], ",") {
		if v != "x" {
			busID, _ := strconv.Atoi(v)
			buses = append(buses, busID)
			part2 = append(part2, busID)
		} else {
			part2 = append(part2, 0)
		}
	}
	i, idx := part1(earlistTimeStamp, buses)
	fmt.Println("part1: ", (i-earlistTimeStamp)*buses[idx])

	p2 := partTwo(part2)
	fmt.Println("part2: ", p2)
}

func part1(start int, buses []int) (int, int) {
	i := start
	for i < start*2 {
		for j, v := range buses {
			if i%v == 0 {
				return i, j
			}
		}
		i++
	}
	return 0, 0
}

func partTwo(buses []int) *big.Int {
	mods := make([]int64, 0)
	remainders := make([]int64, 0)
	for i, v := range buses {
		if v == 0 {
			continue
		}
		mods = append(mods, int64(v))
		remainders = append(remainders, int64((v-i)%v))
	}
	res := chineseRemainderTheorem(mods, remainders)
	return res
}

// Fermat's little theorem
func modInvBig(a *big.Int, m *big.Int) *big.Int {
	var limit, mid, z big.Int
	limit.Exp(a, mid.Sub(m, big.NewInt(2)), nil)
	var mod big.Int
	_, n := z.DivMod(&limit, m, &mod)
	return n
}

func chineseRemainderTheorem(mods []int64, remainders []int64) *big.Int {
	N := big.NewInt(1)
	for _, v := range mods {
		N.Mul(N, big.NewInt(v))
	}

	sum := big.NewInt(0)
	for i, v := range mods {
		var m big.Int
		m.Div(N, big.NewInt(v))
		r := big.NewInt(remainders[i])
		r = r.Mul(r, &m)
		r = r.Mul(r, modInvBig(&m, big.NewInt(v)))
		sum.Add(sum, r)
	}

	var res, z big.Int
	_, n := z.DivMod(sum, N, &res)
	return n
}
