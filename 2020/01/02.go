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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(scanner.Text()) // Println will add back the final '\n'
		line := scanner.Text()
		expense, _ := strconv.Atoi(line)
		report = append(report, expense)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	sort.Ints(report)
	//fmt.Println(report)
	i, j := 0, len(report)-1
	for i < j {
		sum := report[i] + report[j]
		switch {
		case sum == 2020:
			fmt.Println(report[i], report[j])
			i++
			j--
		case sum < 2020:
			i++
		case sum > 2020:
			j--
		}
	}
}
