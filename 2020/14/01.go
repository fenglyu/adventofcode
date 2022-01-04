package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fenglyu/adventofcode/util"
)

var product map[int]any

func main() {

	report := util.ParseBasedOnEachLine()
	//fmt.Println(len(report))

	sumMap := make(map[int]uint64)

	var mask string
	for _, v := range report {
		res := strings.Split(v, " = ")
		if strings.HasPrefix(res[0], "mask") {
			mask = res[1]
			continue
		}
		idx, _ := strconv.Atoi(res[0][4 : len(res[0])-1])
		val, _ := strconv.Atoi(res[1])
		ov := uint64(val)
		sumMap[idx] = overWrite(mask, ov)
	}

	var sum uint64 = 0
	for _, v := range sumMap {
		sum += v
	}

	fmt.Println("Part 1 ", sum)

	//fmt.Println("max count x:", max)
	product = make(map[int]any)

	for i := 1; i < 10; i++ {
		a := []int{0, 1}
		npt := nextProduct(a, i)
		array := make([][]int, 0)
		for {
			np := npt()
			if len(np) == 0 {
				break
			}
			//fmt.Println(np)
			nn := make([]int, len(np))
			copy(nn, np)
			array = append(array, nn)
		}
		product[i] = array
	}

	//fmt.Println(product)
	memMap := make(map[uint64]uint64)
	mask = ""
	for _, v := range report {
		res := strings.Split(v, " = ")
		if strings.HasPrefix(res[0], "mask") {
			mask = res[1]
			continue
		}
		idx, _ := strconv.Atoi(res[0][4 : len(res[0])-1])
		val, _ := strconv.Atoi(res[1])
		index := uint64(idx)
		addr := memAddressDecoder(mask, index)
		//fmt.Println("address: ", addr)
		for _, v := range addr {
			memMap[v] = uint64(val)
		}
	}

	var memSum uint64 = 0
	for _, v := range memMap {
		memSum += v
	}

	fmt.Println("Part 2  ", memSum)
}

func overWrite(mask string, value uint64) uint64 {
	ow := value
	// overwrite the value with all the 1 bit in mask
	var (
		maskOr  uint64 = 0
		maskAnd uint64 = 0
	)
	ar := []byte(mask)
	for i := len(ar) - 1; i >= 0; i-- {
		switch ar[i] {
		case 'X':
			continue
		case '1':
			// bit OR operation for all the 1 bit in mask with value will enforce those position to be 1
			//maskOr += uint64(math.Pow(2, float64(len(ar)-1-i)))
			maskOr += 1 << (len(ar) - 1 - i)
		case '0':
			// mark all the 0 bit to be 1, then do Bitwise NOT operation to revert 1 back to 0 in postion
			//maskAnd += uint64(math.Pow(2, float64(len(ar)-1-i)))
			maskAnd += 1 << (len(ar) - 1 - i)
		}
	}

	maskAnd = ^maskAnd
	// mark '1' bit to be '1'
	ow = ow | maskOr
	// mark '0' bit to be '0'
	ow = ow & maskAnd
	return ow
}

// return a list of memory addresses
func memAddressDecoder(mask string, idx uint64) []uint64 {
	ar := []byte(mask)
	xcount := strings.Count(mask, "X")
	prod := product[xcount].([][]int)
	result := make([]uint64, 0)
	for _, v := range prod {
		ow := idx
		var maskOr uint64 = 0
		j, xarray := len(v)-1, v
		for i := len(ar) - 1; i >= 0; i-- {
			switch ar[i] {
			case '1':
				maskOr += 1 << (len(ar) - 1 - i)
			case '0':
				break
			case 'X':
				switch xarray[j] {
				case 1:
					ow = ow | 1<<(len(ar)-1-i)
				case 0:
					ow = ow & ^(1 << (len(ar) - 1 - i))
				}
				j--
			}
		}
		ow = ow | maskOr
		result = append(result, ow)
	}

	return result
}

func nextProduct(a []int, r int) func() []int {
	p := make([]int, r)
	x := make([]int, len(p))
	return func() []int {
		p := p[:len(x)]
		for i, xi := range x {
			p[i] = a[xi]
		}
		for i := len(x) - 1; i >= 0; i-- {
			x[i]++
			if x[i] < len(a) {
				break
			}
			x[i] = 0
			if i <= 0 {
				x = x[0:0]
				break
			}
		}
		return p
	}
}

/*
Operation	Result	Description
0011 & 0101	0001	Bitwise AND
0011 | 0101	0111	Bitwise OR
0011 ^ 0101	0110	Bitwise XOR
^0101	1010	Bitwise NOT (same as 1111 ^ 0101)
0011 &^ 0101	0010	Bitclear (AND NOT)
00110101<<2	11010100	Left shift
00110101<<100	00000000	No upper limit on shift count
00110101>>2	00001101	Right shift

	//int64(math.Pow(2, 36)) == 68719476736
		// taking the number in the hexadecimal
		// form it is 15 i.e. 00001111 in 8-bit form
		var bitwisenot uint64 = 0xFFFFFFFFFFFFFFFF
		//var bitwisenot uint64 = 0x0

		// printing the number in 8-Bit
		fmt.Printf("%64b\n", bitwisenot)

		// using the ^M = 1 ^ M
		fmt.Printf("%64b\n", ^bitwisenot)
*/
