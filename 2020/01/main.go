package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {

	fN := flag.String("file", "input", "File name")
	flag.Parse()

	file, err := os.Open(*fN)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	report := make([]int, 0)
	expMap := make(map[int]bool)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(scanner.Text()) // Println will add back the final '\n'
		line := scanner.Text()
		expense, _ := strconv.Atoi(line)
		report = append(report, expense)
		expMap[expense] = true
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	sort.Ints(report)
	//fmt.Println(expMap)
	for l := 0; l < len(report)-3; l++ {
		i, j := l+1, len(report)-1
		for i < j {
			sum := report[i] + report[j]
			switch {
			case sum == 2020-report[l]:
				fmt.Println(report[i], report[l], report[j], report[i]*report[l]*report[j])

				i++
				j--
			case sum < 2020-report[l]:
				i++
			case sum > 2020-report[l]:
				j--
			}
		}
	}

}
