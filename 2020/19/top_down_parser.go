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

type rule interface {
	match(m *matcher, s []byte, pos int) []int
}

type matcher struct {
	rules map[uint8]rule
	memo  map[key][]int
}

type literalRule struct {
	char byte
}

type sequenceRule struct {
	subs []uint8
}

type alternationRule struct {
	options [][]uint8
}

func main() {
	report := util.ParseBasedOnEmptyLine()
	if len(report) == 0 {
		return
	}

	rules := parseRules(strings.TrimSpace(report[0]))
	m := newMatcher(rules)
	messages := parseMessages(report)

	part1 := m.countMatches(messages)
	fmt.Println("part1:", part1)

	m.updateRulesForPart2()
	part2 := m.countMatches(messages)
	fmt.Println("part2:", part2)
}

func newMatcher(rules map[uint8]rule) *matcher {
	return &matcher{rules: rules, memo: make(map[key][]int)}
}

func (m *matcher) countMatches(messages []string) int {
	count := 0
	for _, msg := range messages {
		m.memo = make(map[key][]int)
		ends := m.matchRule(0, []byte(msg), 0)
		for _, end := range ends {
			if end == len(msg) {
				count++
				break
			}
		}
	}
	return count
}

func (m *matcher) matchRule(id uint8, s []byte, pos int) []int {
	k := key{id: id, pos: pos}
	if res, ok := m.memo[k]; ok {
		return res
	}

	r, ok := m.rules[id]
	if !ok {
		return nil
	}

	out := r.match(m, s, pos)
	m.memo[k] = out
	return out
}

func (m *matcher) updateRulesForPart2() {
	if _, ok := m.rules[8]; ok {
		m.rules[8] = alternationRule{options: [][]uint8{{42}, {42, 8}}}
	}
	if _, ok := m.rules[11]; ok {
		m.rules[11] = alternationRule{options: [][]uint8{{42, 31}, {42, 11, 31}}}
	}
}

func (r literalRule) match(_ *matcher, s []byte, pos int) []int {
	if pos < len(s) && s[pos] == r.char {
		return []int{pos + 1}
	}
	return nil
}

func (r sequenceRule) match(m *matcher, s []byte, pos int) []int {
	positions := []int{pos}
	for _, sub := range r.subs {
		var next []int
		for _, p := range positions {
			ends := m.matchRule(sub, s, p)
			if len(ends) > 0 {
				next = append(next, ends...)
			}
		}
		if len(next) == 0 {
			return nil
		}
		positions = next
	}
	return positions
}

func (r alternationRule) match(m *matcher, s []byte, pos int) []int {
	var results []int
	for _, seq := range r.options {
		positions := []int{pos}
		for _, sub := range seq {
			var next []int
			for _, p := range positions {
				ends := m.matchRule(sub, s, p)
				if len(ends) > 0 {
					next = append(next, ends...)
				}
			}
			if len(next) == 0 {
				positions = nil
				break
			}
			positions = next
		}
		if len(positions) > 0 {
			results = append(results, positions...)
		}
	}
	return results
}

func parseRules(block string) map[uint8]rule {
	rules := make(map[uint8]rule)
	if block == "" {
		return rules
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
			rules[uint8(ruleID)] = literalRule{char: value[0]}
		case strings.Contains(body, "|"):
			alternatives := strings.Split(body, "|")
			seqs := make([][]uint8, 0, len(alternatives))
			for _, alt := range alternatives {
				seq := parseSequence(strings.TrimSpace(alt))
				seqs = append(seqs, seq)
			}
			rules[uint8(ruleID)] = alternationRule{options: seqs}
		default:
			rules[uint8(ruleID)] = sequenceRule{subs: parseSequence(body)}
		}
	}
	return rules
}

func parseSequence(res string) []uint8 {
	if strings.TrimSpace(res) == "" {
		return nil
	}
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
