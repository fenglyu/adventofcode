package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/fenglyu/adventofcode/util"
)

type conditionFunc func(int) bool
type actionFunc func(int) int
type chainFunc func(int) int

type conditionActionPair struct {
	condition conditionFunc
	action    actionFunc
}

func piecewiseFunction(pairs []conditionActionPair) chainFunc {
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
		// Capture the current value of i
		i := i
		pairs[i] = conditionActionPair{
			condition: func(x int) bool {
				return x >= arr[i][1] && x < arr[i][1]+arr[i][2]
			},
			action: func(x int) int { return arr[i][0] + (x - arr[i][1]) },
		}
	}
	return pairs
}

func locatePostion(seed int, pf [][]conditionActionPair) int {
	loc := seed
	for _, v := range pf {
		pfunc := piecewiseFunction(v)
		loc = pfunc(loc)
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
	//fmt.Println("seeds: ", seeds)
	//	pwfChains := make([]chainFunc, len(raw)-1)
	capArr := make([][]conditionActionPair, 0)
	for _, v := range raw[1:] {
		categories := strings.Split(v, ":")
		_, val := strings.TrimSpace(categories[0]), strings.TrimSpace(categories[1])
		//fmt.Println(hdr)
		//fmt.Println(map2Arr(val))
		//piecewiseFunc := piecewiseFunction(x, genFunc(map2Arr(val)))
		capArr = append(capArr, genFunc(map2Arr(val)))
	}

	fmt.Println("problem 1:", lowestLocNum(seeds, capArr))

	//min := math.MaxInt32
	/*
		arr := make([]int, 0)
		for i := 0; i < len(seeds); i += 2 {
			//sArr := genArr(seeds[i], seeds[i+1])
			//n := lowestLocNum(sArr, capArr)
			arr = append(arr, seeds[i])
		}
		fmt.Println(arr)
		fmt.Println("problem 2:", lowestLocNum(arr, capArr))
	*/

	min := math.MaxInt32
	var wg sync.WaitGroup
	resCh := make(chan int, len(seeds)/2)
	startTime := time.Now()
	for i := 0; i < len(seeds); i += 2 {
		sArr := genArr(seeds[i], seeds[i+1])
		wg.Add(1)
		go func() {
			defer wg.Done()
			n := lowestLocNum(sArr, capArr)
			resCh <- n
		}()
	}
	go func() {
		wg.Wait()
		close(resCh)
	}()

	for v := range resCh {
		if v < min {
			min = v
		}
	}
	endTime := time.Now()
	duration := endTime.Sub(startTime)

	fmt.Println("problem 2: ", min)
	fmt.Println("Execution time: ", duration)
	/*
	   problem 2:  79004094
	   Execution time:  2m28.753255334s
	*/
}

func genArr(first, count int) []int {
	res := make([]int, count)
	for i := 0; i < count; i++ {
		res[i] = first + i
	}
	return res
}

func lowestLocNum(seeds []int, capArr [][]conditionActionPair) int {
	min := math.MaxInt32
	for _, s := range seeds {
		r := locatePostion(s, capArr)
		if r < min {
			min = r
		}
	}
	return min
}
