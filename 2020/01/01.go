package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	//"strconv"
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

	//fmt.Println(report)
	for _, v := range report {
		if _, ok := expMap[2020-v]; ok {
			fmt.Println(v, 2020-v, v*(2020-v))
		}
	}
}
