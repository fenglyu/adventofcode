package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fenglyu/adventofcode/util"
)

func main() {

	report := util.ParseBasedOnEachLine()
	//fmt.Println(report)
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

	fmt.Println("part2: ", part2)
	//partTwo(part2)
	test(part2)
}
func test(buses []int) {
	for i := 1; i < 20; i++ {
		fmt.Printf("> %d ", buses[0]*i)
	}

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

func partTwo(buses []int) {
	start := 100000000000000
	i := start
	for i < start*2 {
		if i%buses[0] == 0 {
			break
		}
		i++
	}
	fmt.Println("i: ", i)
	start = i

	for start < i*2 {

		flag := 1

		for j, v := range buses {
			if v == 0 {
				continue
			}
			r := 0
			if (start+j)%v == 0 {
				r = 1
			}
			flag = flag & r
		}

		if flag == 1 {
			fmt.Println("start :", start)
			break
		}
		//fmt.Println("flag :", flag, start)
		start = start + buses[0]
	}
}
