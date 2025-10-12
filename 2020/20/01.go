package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fenglyu/adventofcode/util"
)

//	type key []uint8

type mtx struct {
	title int
	data  [][]uint8
}

func (m *mtx) String() string {
	return fmt.Sprintf("Title: %d, %s", m.title, util.Matrix2Str(m.data))
}

func newMtx(raw string) *mtx {
	var title int
	res := strings.Split(raw, ":")
	if len(res) < 2 {
		return nil
	}

	tileArr := strings.Fields(res[0])
	if len(tileArr) < 2 {
		return nil
	}

	title, err := strconv.Atoi(tileArr[1])
	if err != nil {
		fmt.Println(err)
		return nil
	}

	data := make([][]uint8, 0)
	for _, v := range strings.Split(res[1], "\n") {
		if len(v) == 0 {
			continue
		}

		data = append(data, []uint8(v))
	}

	return &mtx{
		title: title,
		data:  data,
	}
}

type matrix struct {
	data []mtx
}

func (ma *matrix) print() {
	for _, v := range ma.data {
		fmt.Println("tile: ", v.title)
		util.PrintMatrix(v.data)
	}
}

func newMatric(raw []string) *matrix {
	if len(raw) == 0 {
		return nil
	}
	res := make([]mtx, len(raw))

	for i, v := range raw {
		res[i] = *newMtx(v)
	}
	return &matrix{data: res}
}

func main() {

	report := util.ParseBasedOnEmptyLine()
	//fmt.Println(len(report))
	if len(report) == 0 {
		return
	}

	matrix := newMatric(report)
	matrix.print()
	//fmt.Println(matrix)

}
