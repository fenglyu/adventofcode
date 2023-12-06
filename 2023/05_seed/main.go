package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fenglyu/adventofcode/util"
)

type conditionFunc func(int) bool
type actionFunc func(int) int
type chainFunc func(int) int

type conditionActionPair struct {
	condition conditionFunc
	action    actionFunc
}

type pieceFuncChain struct {
	chains []chainFunc
}

func piecewiseFunction(x int, pairs []conditionActionPair) chainFunc {
	return func(x int) int {
		for _, pair := range pairs {
			if pair.condition(x) {
				return pair.action(x)
			}
		}
		return x
	}
}

func genFunc(arr [][]int) []conditionActionPair {
	pairs := make([]conditionActionPair, len(arr))
	for i := 0; i < len(arr); i++ {
		pairs[i] = conditionActionPair{
			condition: func(x int) bool { return x >= arr[i][1] && x < arr[i][1]+arr[i][2] },
			action:    func(x int) int { return arr[i][0] + (x - arr[i][1]) },
		}
	}
	return pairs
}

func locatePostion(seed int, pieceChains []chainFunc) int {
	loc := seed
	for _, pieceFunc := range pieceChains {
		loc = pieceFunc(loc)
	}
	return loc
}

func head2Arr(s string) []int {
	categories := strings.Split(s, ":")
	sArr := strings.Fields(strings.TrimSpace(strings.TrimSpace(categories[1])))
	arr := make([]int, len(sArr))
	for i, v := range sArr {
		iv, _ := strconv.Atoi(v)
		arr[i] = iv
	}
	return arr
}

func map2Arr(s string) [][]int {
	val := make([][]int, 0)
	rules := strings.Split(s, "\n")
	for _, r := range rules {
		sArr := strings.Fields(strings.TrimSpace(r))
		arr := make([]int, len(sArr))
		for i, v := range sArr {
			iv, _ := strconv.Atoi(v)
			arr[i] = iv
		}
		val = append(val, arr)
	}
	return val
}

func main() {
	raw := util.ParseBasedOnEmptyLine()
	//fmt.Println(len(raw), raw[0])

	seeds := head2Arr(raw[0])
	fmt.Println("seeds: ", seeds)
	pwfChains := make([]chainFunc, len(raw)-1)
	for _, v := range raw[1:] {
		categories := strings.Split(v, ":")
		hdr, val := strings.TrimSpace(categories[0]), strings.TrimSpace(categories[1])
		fmt.Println(hdr)
		fmt.Println(map2Arr(val))
		piecewiseFunc := piecewiseFunction(x, genFunc(map2Arr(val)))

	}
}
