package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

	//fmt.Println(report)
	count := 0
	for _, v := range report {
		if validate(v) {
			count++
		}
	}

	fmt.Println(count)
}

func validate(line string) bool {
	res := strings.Split(line, ":")
	if len(res) < 2 {
		return false
	}
	rule, pass := res[0], res[1]
	return checkRule(rule, pass)
}

func checkRule(rule string, pass string) bool {
	res := strings.Fields(rule)
	scope, letter := res[0], res[1]

	ss := strings.Split(scope, "-")
	l, _ := strconv.Atoi(ss[0])
	r, _ := strconv.Atoi(ss[1])
	target := []byte(letter)[0]
	p := []byte(pass)
	return (p[l] == target || p[r] == target) && !(p[l] == target && p[r] == target)
}
