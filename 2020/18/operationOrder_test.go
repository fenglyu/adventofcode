package main

import (
	"fmt"
	"log"
	"testing"
)

func TestCalc(t *testing.T) {
	expr := []byte{'4', '+', '5', '*', '3', '+', '6'} //33
	stk := newStack()
	for i := len(expr) - 1; i >= 0; i-- {
		stk.push(expr[i])
	}
	res := calc(stk)
	//fmt.Println("res: ", res)
	if res != 33 {
		t.Fatalf("expected result 33, actuall result %d\n", res)
	}
}

func TestEvaluate(t *testing.T) {
	tests := []struct {
		expression string
		result     int
	}{
		{
			expression: "2 * 3 + (4 * 5)",
			result:     26,
		},
		{
			expression: "5 + (8 * 3 + 9 + 3 * 4 * 3)",
			result:     437,
		},
		{
			expression: "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",
			result:     12240,
		},
		{
			expression: "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
			result:     13632,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d-%s", i, tt.expression), func(t *testing.T) {
			result := evaluate([]byte(tt.expression), calc)
			if result != tt.result {
				log.Fatal("Expected: ", tt.result, "\n", " Result: ", result)
			}
		})
	}
}

func TestAdvancedEvaluate(t *testing.T) {

	tests := []struct {
		expression string
		result     int
	}{
		{
			expression: "1 + (2 * 3) + (4 * (5 + 6))",
			result:     51,
		},
		{
			expression: "2 * 3 + (4 * 5)",
			result:     46,
		},
		{
			expression: "5 + (8 * 3 + 9 + 3 * 4 * 3)",
			result:     1445,
		},
		{
			expression: "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",
			result:     669060,
		},
		{
			expression: "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
			result:     23340,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d-%s", i, tt.expression), func(t *testing.T) {
			result := evaluate([]byte(tt.expression), advancedCalc)
			if result != tt.result {
				log.Fatal("Expected: ", tt.result, "\n", " Result: ", result)
			}
		})
	}
}

func TestAdvancedCalc(t *testing.T) {
	expr := []byte{'2', '*', '3', '+', '8'}
	stk := newStack()
	for i := len(expr) - 1; i >= 0; i-- {
		stk.push(expr[i])
	}
	res := advancedCalc(stk)
	if res != 22 {
		t.Fatalf("expected result 33, actuall result %d\n", res)
	}
}
