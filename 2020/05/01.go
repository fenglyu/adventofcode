package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"

	_set "github.com/deckarep/golang-set"
)

func main() {

	fN := flag.String("file", "input", "File name")
	flag.Parse()

	file, err := os.Open(*fN)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	report := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(scanner.Text()) // Println will add back the final '\n'
		line := scanner.Text()
		report = append(report, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

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
	fmt.Println(array)
	fmt.Println(lar, len(array))
	//fmt.Println(set)
	sort.Ints(array)
	scanSet := makeUnsafeSet(array)
	allSet := _set.NewSet()
	for i := array[0]; i <= lar; i++ {
		allSet.Add(i)
	}
	//fmt.Println(allSet, len(allSet.ToSlice()))
	fmt.Println(scanSet, len(scanSet.ToSlice()))

	fmt.Println(allSet.Difference(scanSet))

}

func makeUnsafeSet(ints []int) _set.Set {
	set := _set.NewThreadUnsafeSet()
	for _, i := range ints {
		set.Add(i)
	}
	return set
}

/*
func search(array []int, l int, r int) {
*/

func decode(seat string) (row int, col int) {
	f, l := []byte(seat)[0:6], []byte(seat)[7:]

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

	//fmt.Println(i, j)
	return row, col
}
