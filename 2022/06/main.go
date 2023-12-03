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

	part2(fish, 1, 256)
	sum := 0
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
		length := len(smap)
		count := 0
		seq := make([]int, length)
		j := 0
		for k, _ := range smap {
			seq[j] = k
			j++
		}
		sort.Ints(seq)

		// cond 2 day == 4
		for i, v := range seq {
			if v == 0 {
				smap[10] = smap[0]
				count += smap[0]
			} else {
				smap[v-1] = smap[v]
			}

			if _, Ok := smap[v+1]; !Ok {
				delete(smap, v)
			}

			if i == len(seq)-1 {
				delete(smap, v)
			}
		}

		// new fish, and renew
		if v, Ok := smap[10]; Ok && v > 0 {
			smap[6] = smap[6] + smap[10]
			smap[8] += smap[10]
		}

		delete(smap, 10)

		day++
	}
}
