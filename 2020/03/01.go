package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	fN := flag.String("file", "input", "File name")
	flag.Parse()

	file, err := os.Open(*fN)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//report := make([]int, 0)
	report := [][]uint8{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(scanner.Text()) // Println will add back the final '\n'
		line := scanner.Text()
		report = append(report, []byte(line))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	fmt.Println(len(report), len(report[0]), report[321][12])
	fmt.Println(walk(report, 3, 1))

	fmt.Println(walk(report, 1, 1) * walk(report, 3, 1) * walk(report, 5, 1) * walk(report, 7, 1) * walk(report, 1, 2))
}

func walk(grid [][]uint8, right int, down int) int {
	i, j := 0, 0
	os, tree := 0, 0
	uknown := 0
	for i < len(grid) {
		switch {
		case grid[i][j] == '.':
			os++
		case grid[i][j] == '#':
			tree++
		default:
			uknown++
		}
		j = (j + right) % len(grid[0])
		i = i + down
	}

	//return os, tree
	return tree
}
