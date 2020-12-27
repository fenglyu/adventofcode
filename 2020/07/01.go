package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func parseBasedOnEmptyLine() []string {
	fN := flag.String("file", "input", "File name")
	flag.Parse()

	file, err := os.Open(*fN)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	report := make([]string, 0)

	scanner := bufio.NewScanner(file)
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {

		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		for i := 0; i < len(data)-1; i++ {
			if data[i] == '\n' && data[i+1] == '\n' {
				return i + 1, data[0:i], nil
			}
		}

		if atEOF {
			return len(data), data[0:len(data)], nil
		}
		// Request more data
		return 0, nil, nil
	}

	scanner.Split(onComma)

	for scanner.Scan() {
		//fmt.Println(scanner.Text()) // Println will add back the final '\n'
		line := scanner.Text()
		report = append(report, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return report
}

func parseBasedOnEachLine() []string {
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
	return report
}

func main() {

	report := parseBasedOnEachLine()

	//	sum := 0
	//initColors := make([]string, 0)
	stats := make(map[string][]string, 0)
	for _, v := range report {
		/*
			if regexp.MustCompile(`^(.*) bags contain.*[1-9] (shiny gold) bag[s]?`).MatchString(v) {
				initColors = append(initColors)
			}
		*/
		//fmt.Printf("[%d]: %s\n", i, v)
		q := regexp.MustCompile(`([0-9]|bags|bag|,)`).Split(v, -1)
		if len(q) > 0 {
			//fmt.Printf("[%d]: %q\n", i, q)
			result := make([]string, 0)
			for j, color := range q {
				if j == 0 || colorIn(color, []string{"", "contain", ".", "contain no other"}) {
					continue
				}
				result = append(result, color)
			}
			//fmt.Println(result)
			stats[strings.Trim(q[0], " ")] = result
		}
		//}
	}

	//	fmt.Println(stats)
	resultBags := bagBelongTo("shiny gold", stats)
	fmt.Println("sum: ")

	for i, r := range resultBags {
		fmt.Printf("[%d] %s\n", i, r)
	}

	result := countBags("shiny gold", stats, 0)
	fmt.Println("result: ", result)
}

func colorIn(color string, exa []string) bool {
	for _, v := range exa {
		if strings.EqualFold(strings.Trim(color, " "), strings.Trim(v, " ")) {
			return true
		}
	}
	return false
}

func bagBelongTo(bags string, stats map[string][]string) []string {
	outerBag := make([]string, 0)
	for k, v := range stats {
		if colorIn(bags, v) {
			outerBag = append(outerBag, k)
		}
	}
	return outerBag
}

func countBags(bags string, stats map[string][]string, lvl int) int {
	sum := 0
	pBags := make([]string, 0)
	for k, v := range stats {
		if colorIn(bags, v) {
			sum++
			pBags = append(pBags, k)
		}
	}
	fmt.Println("level >", lvl, "bags >", bags, " pBags >", pBags)
	lvl++
	for _, v := range pBags {
		sum += countBags(v, stats, lvl)
	}

	return sum
}
