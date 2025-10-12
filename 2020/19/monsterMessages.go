package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/fenglyu/adventofcode/util"
)

type key struct {
	id  uint8
	pos int
}

var mapp map[uint8]any
var memo map[key][]int

func concatBytes(res string) []uint8 {
	fields := strings.Fields(res)
	u8a := make([]uint8, len(fields))
	for i, f := range fields {
		ia, err := strconv.Atoi(f)
		if err != nil {
			log.Fatalf("invalid rule reference %q: %v", f, err)
		}
		u8a[i] = uint8(ia)
	}
	return u8a
}

func main() {
	report := util.ParseBasedOnEmptyLine()
	if len(report) == 0 {
		return
	}

	mapp = make(map[uint8]any)
	parseRules(strings.TrimSpace(report[0]))
	messages := parseMessages(report)

	part1 := countMatches(messages)
	fmt.Println("part1:", part1)

	updateRulesForPart2()
	part2 := countMatches(messages)
	fmt.Println("part2:", part2)
}

func matchRule(str []byte, row uint8, pos int) []int {
	k := key{id: row, pos: pos}
	if res, Ok := memo[k]; Ok {
		return res
	}

	value, Ok := mapp[row]
	if !Ok {
		return nil
	}

	var out []int
	switch v := value.(type) {
	case uint8:
		if pos < len(str) && str[pos] == byte(v) {
			out = []int{pos + 1}
		} else {
			out = nil
		}

	case []uint8:
		positions := []int{pos}
		for _, sub := range v {
			var next []int
			for _, p := range positions {
				ends := matchRule(str, sub, p)
				if len(ends) > 0 {
					next = append(next, ends...)
				}
			}
			positions = next
			if len(positions) == 0 {
				break
			}
		}
		out = positions

	case [][]uint8:
		var results []int
		for _, seq := range v {
			// run the same sequence logic as above
			positions := []int{pos}
			for _, sub := range seq {
				var next []int
				for _, p := range positions {
					ends := matchRule(str, sub, p)
					if len(ends) > 0 {
						next = append(next, ends...)
					}
				}
				positions = next
				if len(positions) == 0 {
					break
				}
			}
			results = append(results, positions...)
		}
		out = results

	default:
		out = nil
	}

	memo[k] = out
	return out
}

func parseRules(block string) {
	if block == "" {
		return
	}
	for _, line := range strings.Split(block, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			log.Fatalf("invalid rule line: %q", line)
		}
		ruleID, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("invalid rule id %q: %v", parts[0], err)
		}
		body := parts[1]
		switch {
		case strings.Contains(body, "\""):
			value := []byte(strings.Trim(body, "\""))
			if len(value) != 1 {
				log.Fatalf("invalid literal rule: %q", line)
			}
			mapp[uint8(ruleID)] = uint8(value[0])
		case strings.Contains(body, "|"):
			alternatives := strings.Split(body, "|")
			seqs := make([][]uint8, 0, len(alternatives))
			for _, alt := range alternatives {
				seq := concatBytes(strings.TrimSpace(alt))
				seqs = append(seqs, seq)
			}
			mapp[uint8(ruleID)] = seqs
		default:
			mapp[uint8(ruleID)] = concatBytes(body)
		}
	}
}

func parseMessages(report []string) []string {
	if len(report) < 2 {
		return nil
	}
	block := strings.TrimSpace(report[1])
	if block == "" {
		return nil
	}
	lines := strings.Split(block, "\n")
	msgs := make([]string, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		msgs = append(msgs, line)
	}
	return msgs
}

func countMatches(messages []string) int {
	count := 0
	for _, msg := range messages {
		memo = make(map[key][]int)
		ends := matchRule([]byte(msg), 0, 0)
		for _, end := range ends {
			if end == len(msg) {
				count++
				break
			}
		}
	}
	return count
}

func updateRulesForPart2() {
	if _, ok := mapp[8]; ok {
		mapp[8] = [][]uint8{{42}, {42, 8}}
	}
	if _, ok := mapp[11]; ok {
		mapp[11] = [][]uint8{{42, 31}, {42, 11, 31}}
	}
}
