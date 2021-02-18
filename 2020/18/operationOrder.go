package main

import (
	"container/list"
	"fmt"

	"github.com/fenglyu/adventofcode/util"
)

type calculate func(*stack) int

type stack struct {
	l *list.List
}

func (s *stack) String() string {
	return fmt.Sprintf("%v", s.l)
}

func newStack() *stack {
	return &stack{l: list.New()}
}

func (s *stack) push(e interface{}) {
	s.l.PushBack(e)
}

func (s *stack) pop() interface{} {
	e := s.l.Back()
	if e != nil {
		return s.l.Remove(e)
	}
	return nil
}

func (s *stack) reverse() *stack {
	rL := list.New()
	for e := s.l.Back(); e != nil; e = e.Prev() {
		rL.PushBack(e.Value)
	}
	return &stack{l: rL}
}

func (s *stack) Len() int {
	return s.l.Len()
}

func (s *stack) print() {
	for e := s.l.Front(); e != nil; e = e.Next() {
		switch e.Value.(type) {
		case uint8:
			fmt.Printf("%c->", e.Value.(uint8))
		case int:
			fmt.Printf("%d->", e.Value.(int))
		}
	}
	fmt.Println("")
}

func main() {

	report := util.ParseBasedOnEachLine()
	//evaluate([]byte(report[0]))
	//res := evaluate([]byte("1 + (2 * 3) + (4 * (5 + 6) + (2 + 3))"))
	sum := 0
	for _, v := range report {
		sum += evaluate([]byte(v), calc)
	}

	fmt.Println("part1: ", sum)

	adSum := 0
	for _, v := range report {
		adSum += evaluate([]byte(v), advancedCalc)
	}
	fmt.Println("part2: ", adSum)
}

func trimSpace(arr []uint8) []uint8 {
	res := make([]uint8, 0)
	for _, v := range arr {
		if v == ' ' {
			continue
		}
		res = append(res, v)
	}
	return res
}

func evaluate(homework []uint8, cal calculate) int {
	filo := newStack()
	var res int
	pure := trimSpace(homework)
	for _, v := range pure {
		if v == ' ' {
			continue
		}
		filo.push(v)
		if v == ')' {
			calcStack := newStack()
			for {
				e := filo.pop()
				if s, ok := e.(uint8); ok {
					if s == '(' {
						break
					} else if s == ')' {
						continue
					}
				}
				calcStack.push(e)
			}
			// calculate the expression within parenthes
			res = cal(calcStack)
			filo.push(res)
		}
	}
	// reverse the stack, cause the calc function process expression stored reversely in stack
	res = cal(filo.reverse())
	return res
}

func calc(s *stack) int {
	//s.print()
	numArr, signArr := make([]int, 0), make([]uint8, 0)
	for {
		e := s.pop()
		if e == nil {
			break
		}
		switch e.(type) {
		case uint8:
			switch e.(uint8) {
			case '+', '*':
				signArr = append(signArr, e.(uint8))
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				numArr = append(numArr, int(e.(uint8)-'0'))
			}
		case int:
			numArr = append(numArr, e.(int))
		}
	}
	if len(numArr) == 0 {
		return 0
	}

	/*
		fmt.Println("numArr: ", numArr)
		fmt.Println("signArr: ", signArr)
	*/
	init := numArr[0]
	for i := 0; i < len(signArr); i++ {
		j := i + 1
		switch signArr[i] {
		case '+':
			init += numArr[j]
		case '*':
			init *= numArr[j]
		}
	}
	return init
}

func advancedCalc(s *stack) int {
	s.reverse()
	for e := s.l.Front(); e != nil; e = e.Next() {
		if z, ok := e.Value.(uint8); ok {
			if z == '+' {
				p, n := e.Prev(), e.Next()
				sum := valueSum(p.Value, n.Value)
				newEl := s.l.InsertAfter(sum, n)
				s.l.Remove(p)
				s.l.Remove(e)
				s.l.Remove(n)
				e = newEl
			}
		}
	}
	res := calc(&stack{l: s.l})
	return res
}

func valueSum(a interface{}, b interface{}) int {
	return valueToInt(a) + valueToInt(b)
}

func valueToInt(a interface{}) int {
	var b int
	switch a.(type) {
	case uint8:
		b = int(a.(uint8) - '0')
	case int:
		b = a.(int)
	}
	return b
}
