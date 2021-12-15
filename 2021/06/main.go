package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/fenglyu/adventofcode/util"
)

func main() {
	report := util.ParseBasedOnEachLine()
	l := strings.Split(report[0], ",")
	//fmt.Println(l)
	fishList := make([]int, len(l))
	for i := range fishList {
		v, _ := strconv.Atoi(l[i])
		fishList[i] = v
	}

	fishList = lanternfish(fishList, len(fishList), 0, 80)

	fmt.Println("part 1:", len(fishList))

	fish := make(map[int]int)
	for i := range l {
		v, _ := strconv.Atoi(l[i])
		fish[v]++
	}

	part2(fish, 0, 256)
	sum := 0
	fmt.Println(fish)
	for _, v := range fish {
		sum += v
	}
	fmt.Println("part 2:", sum)
}

func lanternfish(state []int, len int, day int, limit int) []int {

	if day >= limit {
		return state
	}

	count := 0
	for i, v := range state {
		if v == 0 {
			state[i] = 6
			count++
		} else {
			state[i]--
		}
	}

	for i := 0; i < count; i++ {
		state = append(state, 6+2)
	}

	return lanternfish(state, len, day+1, limit)
}

func part2(smap map[int]int, day int, limit int) {

	for day <= limit {
		fmt.Println("day: ", day)
		length := len(smap)
		count := 0
		seq := make([]int, length)
		j := 0
		for k, _ := range smap {
			seq[j] = k
			j++
		}
		sort.Ints(seq)

		for i, v := range seq {
			if v == 0 {
				smap[9] = smap[0]
				smap[8]++
				delete(smap, 0)
				count++
			} else {
				smap[v-1] = smap[v]
			}

			if i == len(seq)-1 {
				delete(smap, v)
			}
		}

		if count > 0 {
			smap[8] += count
		}

		if len(smap) >= 7 {
			smap[6] = smap[6] + smap[9]
		}

		delete(smap, 9)

		fmt.Println("smap: ", smap)
		day++
	}
}
