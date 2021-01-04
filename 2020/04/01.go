package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/fenglyu/adventofcode/util"
)

func main() {

	report := util.ParseBasedOnEmptyLine()
	//fmt.Println(report)

	count := 0
	for _, v := range report {
		if valid(v) {
			count++
		}
	}
	fmt.Println(count)

	c := 0
	nums := make([]int, 0)
	for i, v := range report {
		if valid2(v) {
			nums = append(nums, i)
			c++
		}
	}
	fmt.Println(c)
	//fmt.Println(nums)
	//fmt.Println(report[75])

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

func checkValue(s string, len int, min int, max int) bool {
	return regexp.MustCompile(fmt.Sprintf(`[\d]{%d}`, len)).MatchString(s) && numRange(s, min, max)
}

func numRange(s string, min int, max int) bool {
	v, _ := strconv.Atoi(s)
	return v >= min && v <= max
}

func valid2(passport string) bool {

	if !valid(passport) {
		return false
	}

	ff := strings.Fields(passport)

	if len(ff) < 7 || len(ff) > 8 {
		return false
	}

	for _, f := range ff {
		ts := strings.Split(f, ":")
		if ts[0] != "cid" && (ts[0] == "" || ts[1] == "") {
			return false
		}

		switch ts[0] {
		case "byr":
			if !checkValue(ts[1], 4, 1920, 2002) {
				return false
			}
		case "iyr":
			if !checkValue(ts[1], 4, 2010, 2020) {
				return false
			}
		case "eyr":
			if !checkValue(ts[1], 4, 2020, 2030) {
				return false
			}
		case "hgt":
			if !regexp.MustCompile(`(cm|in)$`).MatchString(ts[1]) {
				return false
			}
			if strings.Contains(ts[1], "cm") {
				ll := strings.TrimSuffix(ts[1], "cm")
				if !numRange(ll, 150, 193) {
					return false
				}
			} else if strings.Contains(ts[1], "in") {
				ll := strings.TrimSuffix(ts[1], "in")
				if !numRange(ll, 59, 76) {
					return false
				}
			}
		case "hcl":
			if !regexp.MustCompile(`^#[0-9a-f]{6}$`).MatchString(ts[1]) {
				return false
			}
		case "ecl":
			test := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
			f := false
			for _, l := range test {
				if strings.EqualFold(ts[1], l) {
					f = true
				}
			}
			if !f {
				return false
			}
		case "pid":
			if !regexp.MustCompile(`^[0-9]{9}$`).MatchString(ts[1]) {
				return false
			}

		default:
			//fmt.Println("")
		}
	}

	return true
}
