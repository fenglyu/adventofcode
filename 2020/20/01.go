package main

import (
	"fmt"
	"strconv"
	"strings"

	set "github.com/deckarep/golang-set/v2"
	"github.com/fenglyu/adventofcode/util"
)

//	type key []uint8

type card struct {
	title int
	data  [][]uint8
	edges []int
}

func (m *card) cols(idx int) []uint8 {
	if idx < 0 || idx >= len(m.data) {
		return nil
	}

	res := make([]uint8, len(m.data))
	for i := 0; i < len(m.data); i++ {
		res[i] = m.data[i][idx]
	}
	return res
}

func (m *card) val(ar []uint8) int {
	var cval int = 0
	for i := 0; i < len(ar); i++ {
		if ar[i] == '#' {
			cval = cval<<1 + 1
		}
	}
	return cval
}

func (m *card) setEdges() [][]uint8 {
	res := make([][]uint8, 4)
	res[0], res[2] = m.data[0], m.data[9]
	res[1], res[3] = m.cols(0), m.cols(9)

	edgeInts := make([]int, 4)
	for _, v := range res {
		edgeInts = append(edgeInts, m.val(v))
	}
	m.edges = edgeInts
	return res
}

func (m *card) String() string {
	return fmt.Sprintf("Title: %d, %s", m.title, util.Matrix2Str(m.data))
}

func newMtx(raw string) *card {
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

	return &card{
		title: title,
		data:  data,
	}
}

type matrix struct {
	data  []card
	index map[int]card
}

func (ma *matrix) print() {
	for _, v := range ma.data {
		fmt.Println("tile: ", v.title)
		util.PrintMatrix(v.data)
	}
}

func (ma *matrix) pair() {
	for i, card := range ma.data {
		card.setEdges()
		ma.data[i] = card
	}
	memo := make(map[int]set.Set[int])

	for _, card := range ma.data {
		for _, e := range card.edges {
			if v, ok := memo[e]; ok {
				tileSet := v
				tileSet.Add(card.title)
				memo[e] = tileSet
			} else {
				tileSet := set.NewSet[int]()
				tileSet.Add(card.title)
				memo[e] = tileSet
			}
		}
	}
	fmt.Println(memo)
}

func newMatric(raw []string) *matrix {
	if len(raw) == 0 {
		return nil
	}
	res := make([]card, len(raw))
	index := make(map[int]card, len(raw))
	for i, v := range raw {
		res[i] = *newMtx(v)
		index[res[i].title] = res[i]
	}
	return &matrix{data: res, index: index}
}

func main() {

	report := util.ParseBasedOnEmptyLine()
	//fmt.Println(len(report))
	if len(report) == 0 {
		return
	}

	matrix := newMatric(report)
	//matrix.print()
	//fmt.Println(matrix)
	matrix.pair()

}
