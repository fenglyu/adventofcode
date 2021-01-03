package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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
			return len(data), data[0:], nil
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
	//fN := flag.String("file", "test_input", "File name")
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

var calc map[int]bool

type Record struct {
	lineno int
	key    string
	nums   []int
	value  []string
	tag    bool
}

func (r Record) String() string {
	return fmt.Sprintf("lineno %d, key %s, nums %v, value %v, tag %v", r.lineno, r.key, r.nums, r.value, r.tag)
}

var count map[int]bool
var count2 map[int]bool

func main() {

	report := parseBasedOnEachLine()
	recordStat := make(map[string]*Record, 0)

	for i, v := range report {
		//fmt.Printf("[%d]: %s\n", i, v)
		q := regexp.MustCompile(`([0-9]|bags|bag|,)`).Split(v, -1)
		//var result []string
		result := make([]string, 0)
		nums := make([]int, 0)
		if len(q) > 0 {
			//fmt.Printf("[%d]: %q\n", i, q)
			//idx := 0
			for j, color := range q {
				if j == 0 || colorIn(color, []string{"", "contain", ".", "contain no other"}) {
					continue
				}
				trimStr := strings.Trim(q[j], " ")
				result = append(result, trimStr)
			}

			//var re = regexp.MustCompile(`(?:[a-z\ ]*) bag[s].*(\d{1,}) (?:[a-z\ ]*) bag[s]?(?:\, (\d{1,}) (?:[a-z\ ]*) bag[s])?.`)
			re := regexp.MustCompile(`\d{1,}`)
			m := re.FindAllString(v, -1)
			for _, v := range m {
				if n, err := strconv.Atoi(v); err == nil {
					nums = append(nums, n)
				}
			}

			recordStat[strings.Trim(q[0], " ")] = &Record{
				lineno: i,
				key:    strings.Trim(q[0], " "),
				tag:    false,
				nums:   nums,
				value:  result,
			}
		}
	}

	//fmt.Println(recordStat)
	/*
		for k, v := range recordStat {
			fmt.Println(k, *v)
	}*/
	count = make(map[int]bool, 0)
	part1 := countBags("shiny gold", recordStat)
	fmt.Println("part1 >", part1)

	count2 = make(map[int]bool, 0)
	part2 := countBags2("shiny gold", recordStat)
	fmt.Println("part2 >", part2)
}

func colorIn(color string, exa []string) bool {
	for _, v := range exa {
		if strings.EqualFold(strings.Trim(color, " "), strings.Trim(v, " ")) {
			return true
		}
	}
	return false
}

func countLine(color string, r *Record) bool {
	if _, Ok := count[r.lineno]; !Ok {
		if colorIn(color, r.value) {
			count[r.lineno] = true
			return true
		}
	}
	return false
}

func countBags(bag string, stats map[string]*Record) int {
	sum := 0
	for k, v := range stats {
		if countLine(bag, v) {
			sum++
			sum += countBags(k, stats)
		}
	}

	return sum
}

func countBags2(bag string, stats map[string]*Record) int {
	sum := 0

	v := stats[bag]
	if len(v.nums) == 0 {
		return 1
	}

	for j, q := range v.nums {
		b := countBags2(v.value[j], stats)
		a := q * b
		// when the color has more child
		if b != 1 {
			sum += q
		}
		sum += a
	}

	return sum
}
