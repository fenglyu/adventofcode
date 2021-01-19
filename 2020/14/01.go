package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fenglyu/adventofcode/util"
)

func main() {

	report := util.ParseBasedOnEachLine()
	fmt.Println(len(report))
	//int64(math.Pow(2, 36)) == 68719476736
	/*
		// taking the number in the hexadecimal
		// form it is 15 i.e. 00001111 in 8-bit form
		var bitwisenot uint64 = 0xFFFFFFFFFFFFFFFF
		//var bitwisenot uint64 = 0x0

		// printing the number in 8-Bit
		fmt.Printf("%64b\n", bitwisenot)

		// using the ^M = 1 ^ M
		fmt.Printf("%64b\n", ^bitwisenot)
	*/

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

func memAddressDecoder(mask string, value uint64) uint64 {
	return 0
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
*/
