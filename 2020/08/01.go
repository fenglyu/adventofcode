package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fenglyu/adventofcode/util"
	"github.com/golang-collections/collections/stack"
)

var steps map[int]bool

func main() {

	report := util.ParseBasedOnEachLine()

	trace := stack.New()

	part1(trace, report)

	part2(trace, report)
}

type Track struct {
	idx int
	acc int
	ins string
}

func (t *Track) String() string {
	return fmt.Sprintf("idx %d, acc %d, instruction %s\n", t.idx, t.acc, t.ins)
}

func part1(trace *stack.Stack, report []string) {
	steps = make(map[int]bool, 1)
	for i := 0; i < len(report)-1; i++ {
		steps[i] = false
	}

	acc := 0
	i := 0
	for {
		if i < 0 || i > len(report)-1 {
			fmt.Println(i, acc)
			break
		}
		if v, Ok := steps[i]; Ok && v {
			//fmt.Printf("line index %d again, acc %d\n", i, acc)
			fmt.Println("part1 > ", acc)
			break
		}

		//fmt.Printf("line %d again, acc %d\n", i, acc)
		steps[i] = true

		ins, num := parseCMD(report[i])
		switch ins {
		case "acc":
			acc += num
			trace.Push(&Track{idx: i, acc: acc, ins: report[i]})
			i++

		case "jmp":
			trace.Push(&Track{idx: i, acc: acc, ins: report[i]})
			i += num
		case "nop":
			trace.Push(&Track{idx: i, acc: acc, ins: report[i]})
			i++
		}
	}
	/*
		for trace.Len() > 0 {
			s := trace.Pop().(*Track)
			fmt.Println(s)
		}
	*/
}

func part2(trace *stack.Stack, report []string) {

	for trace.Len() > 0 {
		s := trace.Pop().(*Track)

		acc := s.acc
		i := s.idx

		// don't check the first step
		steps[i] = false
		flag := true

		for {
			if i < 0 || i > len(report)-1 {
				//fmt.Printf("final line %d, acc %d\n", i, acc)
				fmt.Println("part2 > ", acc)
				goto breakHere
			}

			if v, Ok := steps[i]; Ok && v {
				//fmt.Printf("part2 line index %d again, acc %d\n", i, acc)
				break
			}

			//fmt.Printf("line %d again, acc %d\n", i, acc)
			steps[i] = true
			ins, num := parseCMD(report[i])
			switch ins {

			case "acc":
				acc += num
				i++

			case "jmp":
				j := i
				i += num
				if flag {
					if v, Ok := steps[i]; Ok && v {
						i = j + 1
						//fmt.Printf("[jmp] %d [%d] next line is %d again, acc %d\n", j, num, i, acc)
					}
					flag = false
				}

			case "nop":
				j := i
				i++
				if flag {
					if v, Ok := steps[i]; Ok && v {
						i = j + num
						//fmt.Printf("[nop] %d next line is %d again, acc %d\n", j, i, acc)
					}
					flag = false
				}
			}
		}
	}

	// important line to jump out from the inside loop
breakHere:
	fmt.Println("Done here")
}

func parseCMD(cmd string) (string, int) {
	t := strings.Split(cmd, " ")
	ins, arg := t[0], t[1]
	args := []byte(arg)
	sign := args[0]
	num, _ := strconv.Atoi(string(args[1:]))

	var res int
	switch sign {
	case '+':
		res = num
	case '-':
		res = 0 - num
	}

	return ins, res
}
