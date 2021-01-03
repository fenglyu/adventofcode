package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/golang-collections/collections/stack"
)

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

var steps map[int]bool

func main() {

	report := parseBasedOnEachLine()
	fmt.Println("len(report) = ", len(report))

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
			fmt.Printf("line index %d again, acc %d\n", i, acc)
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
				fmt.Printf("final line %d, acc %d\n", i, acc)
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
						fmt.Printf("[jmp] %d [%d] next line is %d again, acc %d\n", j, num, i, acc)
					}
					flag = false
				}

			case "nop":
				j := i
				i++
				if flag {

					if v, Ok := steps[i]; Ok && v {
						i = j + num
						fmt.Printf("[nop] %d next line is %d again, acc %d\n", j, i, acc)
					}
					flag = false
				}
			}
		}
	}

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
