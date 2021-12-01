package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestConcatBytes(t *testing.T) {
	res := "57 68 | 12 110"
	r := strings.Split(res, " | ")
	a, b := strings.Split(r[0], " "), strings.Split(r[1], " ")
	u8a := make([]uint8, 0)
	u8b := make([]uint8, 0)
	for i := 0; i < len(a); i++ {
		ia, _ := strconv.Atoi(a[i])
		ib, _ := strconv.Atoi(b[i])
		u8a = append(u8a, uint8(ia))
		u8b = append(u8b, uint8(ib))
	}
	fmt.Println("u8a", u8a)
	fmt.Println("u8b", u8b)
	fmt.Println(a[0], a[1], b[0], b[1])
	//n := [][]byte{concatBytes(a), concatBytes(b)}
	//fmt.Printf("%q\n", n)
	//fmt.Printf("%c, %c, %c,  %q\n", n[0][0], n[0][1], n[0][2], n[1])
}

/*
func TestSpreadOut(t *testing.T) {

}
*/
