package main

import (
	"fmt"
	"sort"

	_set "github.com/deckarep/golang-set"
	"github.com/fenglyu/adventofcode/util"
)

func main() {

	report := util.ParseBasedOnEachLine()

	//	fmt.Println(report)
	//fmt.Println(decode("FBFBBFFRLR"))
	array := make([]int, 0)
	//scanSet := _set.NewSet()

	lar := 0
	for _, v := range report {
		r, c := decode(v)
		sum := r*8 + c
		if r*8+c > lar {
			lar = sum
		}
		array = append(array, sum)
		//scanSet.Add(sum)
	}
	sort.Ints(array)

	fmt.Println("Part 1: ", lar)

	scanSet := makeSet(array)

	allSet := _set.NewSet()
	for i := array[0]; i <= lar; i++ {
		allSet.Add(i)
	}

	// solution 1
	fmt.Println("scanSet: ", len(scanSet.ToSlice()))
	fmt.Println("Part 2: ", allSet.Difference(scanSet))

	// solution 2, simple search
	idx := search(array, 0, len(array)-2)
	fmt.Println("Part 2: ", array[idx]+1)
}

func makeSet(ints []int) _set.Set {
	set := _set.NewSet()
	for _, i := range ints {
		set.Add(i)
	}
	return set
}

/**/
func search(ar []int, l int, r int) int {
	for i := l; i <= r; i++ {
		if ar[i]+1 == ar[i+1] {
			continue
		} else if ar[i]+2 == ar[i+1] {
			return i
		} else {
			continue
		}
	}
	return 0
}

func decode(seat string) (row int, col int) {
	f, l := []byte(seat)[0:7], []byte(seat)[7:]

	i, j := 0, 127
	for idx, v := range f {

		if v == 'F' {
			if idx == len(f)-1 {
				row = i
			}
			j = (i + j) / 2
		} else if v == 'B' {
			if idx == len(f)-1 {
				row = j
			}
			i = (i+j)/2 + 1
		}
	}

	m, n := 0, 7
	for idx, v := range l {
		if v == 'L' {
			if idx == len(l)-1 {
				col = m
			}
			n = (m + n) / 2
		} else if v == 'R' {
			if idx == len(l)-1 {
				col = n
			}
			m = (m+n)/2 + 1
		}
	}

	return row, col
}
