package main

import (
	"fmt"
	"strconv"
	"strings"

	set "github.com/deckarep/golang-set/v2"
	"github.com/fenglyu/adventofcode/util"
)

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
	var cval int
	for i := 0; i < len(ar); i++ {
		cval = cval << 1
		if ar[i] == '#' {
			cval++
		}
	}
	return cval
}

// reverse returns a copy of the provided edge with its bits flipped left to
// right. The puzzle allows tiles to be rotated or mirrored, so every edge must
// be considered in both directions when checking for matches.
func reverse(line []uint8) []uint8 {
	res := make([]uint8, len(line))
	copy(res, line)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func (m *card) setEdges() [][]uint8 {
	res := make([][]uint8, 4)
	lastIdx := len(m.data) - 1
	res[0], res[2] = m.data[0], m.data[lastIdx]
	res[1], res[3] = m.cols(0), m.cols(lastIdx)

	edgeInts := make([]int, 0, len(res)*2)
	for _, v := range res {
		edgeInts = append(edgeInts, m.val(v))
		edgeInts = append(edgeInts, m.val(reverse(v)))
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

// pair builds an adjacency list keyed by tile id. Every tile id maps to the set
// of other tiles that share at least one matching edge (in either orientation).
// The adjacency map makes it straightforward for later steps of the solution to
// reason about how tiles can be stitched together into the final image.
func (ma *matrix) pair() map[int]set.Set[int] {
	for i := range ma.data {
		card := ma.data[i]
		card.setEdges()
		ma.data[i] = card
	}
	memo := make(map[int]set.Set[int])

	for _, card := range ma.data {
		for _, e := range card.edges {
			if _, ok := memo[e]; !ok {
				memo[e] = set.NewSet[int]()
			}
			memo[e].Add(card.title)
		}
	}

	adjacency := make(map[int]set.Set[int])
	for _, tiles := range memo {
		if tiles.Cardinality() < 2 {
			continue
		}

		ids := make([]int, 0, tiles.Cardinality())
		for v := range tiles.Iter() {
			ids = append(ids, v)
		}

		for i := range ids {
			for j := range ids {
				if i == j {
					continue
				}
				if _, ok := adjacency[ids[i]]; !ok {
					adjacency[ids[i]] = set.NewSet[int]()
				}
				adjacency[ids[i]].Add(ids[j])
			}
		}
	}

	return adjacency
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
	pairs := matrix.pair()
	//fmt.Println(pairs)

	part1 :=1
	for k, v := range pairs{
		//fmt.Println(k, v)
		if v.Cardinality() == 2 {
			//fmt.Println(k, v)
			part1 *= k
		}
	}
	fmt.Println("part: ", part1)
}