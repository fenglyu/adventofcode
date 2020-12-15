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
		line := scanner.Text()
		report = append(report, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	//fmt.Println(report)
	count := 0
	for _, v := range report {
		if valid(v) {
			count++
		}
	}
	fmt.Println(count)

	c := 0
	for _, v := range report {
		if valid2(v) {
			c++
		}
	}
	fmt.Println(c)
}

func valid(passport string) bool {
	fileds := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} // "cid"
	for _, v := range fileds {
		if !strings.Contains(passport, v) {
			return false
		}
	}
	return true
}

func valid2(passport string) bool {
	fileds := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} // "cid"
	fmap := make(map[string]string)
	//fmt.Println(passport)

	//ff := regexp.MustCompile(`[_]+`).Split(passport, -1)
	ff := strings.Fields(passport)

	if len(ff) < 7 || len(ff) > 8 {
		return false
	}
	//fmt.Println(ff, len(ff))
	//fmt.Println("-------------")
	for _, v := range ff {
		ts := strings.Split(v, ":")
		//fmt.Println(ts)
		fmap[ts[0]] = ts[1]
	}
	//fmt.Println(fmap)

	for _, v := range fileds {
		switch v {
		case "byr":
			value, _ := strconv.Atoi(fmap[v])
			//fmt.Println("byr", value)
			return regexp.MustCompile(`[\d]{4}`).MatchString(fmap[v]) && value <= 2002 && value >= 1920
		case "iyr":
			value, _ := strconv.Atoi(fmap[v])
			//fmt.Println("iyr", value)
			return regexp.MustCompile(`[\d]{4}`).MatchString(fmap[v]) && value <= 2020 && value >= 2010
		case "eyr":
			value, _ := strconv.Atoi(fmap[v])
			//fmt.Println("eyr", value)
			return regexp.MustCompile(`[\d]{4}`).MatchString(fmap[v]) && value <= 2030 && value >= 2020
		case "hgt":
			if !regexp.MustCompile(`(cm|in)$`).MatchString(fmap[v]) {
				return false
			}
			length := strings.TrimSuffix(fmap[v], "cm")
			length = strings.TrimSuffix(fmap[v], "in")
			l, _ := strconv.Atoi(length)
			if strings.HasSuffix(fmap[v], "cm") {
				return numRange(l, 150, 193)
			}
			if strings.HasSuffix(fmap[v], "in") {
				return numRange(l, 59, 76)
			}
		case "hcl":
			return regexp.MustCompile(`^#[0-9a-f]{6}`).MatchString(fmap[v])
		case "ecl":
			test := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
			flag := false
			for _, l := range test {
				if fmap[v] == l {
					flag = true
				}
			}
			return flag
		case "pid":
			return regexp.MustCompile(`[0-9]{9}`).MatchString(fmap[v])
		}
	}

	return true
}

/*
func checkValue(len int, min int, max int) bool{
	regexp.MustCompile(`[\d]{%s}`.).MatchString(fmap[v]) && value <= 2002 && value >= 1920
}
*/

func numRange(v int, min int, max int) bool {
	if v >= min && v <= max {
		return true
	}
	return false
}
