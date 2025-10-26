package main

import (
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"

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

func (m *card) stripBoarder() [][]uint8 {
	if len(m.data) < 2 {
		return nil
	}
	core := make([][]uint8, len(m.data)-2)
	for i := 1; i < len(m.data)-1; i++ {
		coreRow := make([]uint8, len(m.data[0])-2)
		copy(coreRow, m.data[i][1:len(m.data[0])-1])
	}
	return core
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

// Adjacency list stores the information as an array of list objects, with one list for each vertex.
// For each vertex u there is a list of Integer Pair Objects representsing the edge (u, v) of weight w.
type edge struct {
	from int
	to   int
	// weight int
}

type vertex struct {
	id int
	v  any
}

type graph struct {
	// Adj []*list.List
	Adj map[int]*list.List
	// data  []card
	index map[int]*card
}

func (ma *graph) print() {
	for k, v := range ma.Adj {
		fmt.Printf("key: %d, val: %v\n", k, v)
		// util.PrintMatrix(v.data)
		for e := v.Front(); e != nil; e = e.Next() {
			fmt.Println(e.Value)
		}
	}
}

// pair builds an adjacency list keyed by tile id. Every tile id maps to the set
// of other tiles that share at least one matching edge (in either orientation).
// The adjacency map makes it straightforward for later steps of the solution to
// reason about how tiles can be stitched together into the final image.
/*
func (ma *graph) pair() map[int]set.Set[int] {
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

	fmt.Println("memo: ", memo)
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
*/

func newGraph(raw []string) *graph {
	if len(raw) == 0 {
		return nil
	}
	//res := make([]card, len(raw))
	adj:=make(map[int]*list.List)
	index := make(map[int]*card, len(raw))
	for _, v := range raw {
		card := newMtx(v)
		/*
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.LittleEndian, res[i].title)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}
		id := xxhash.Sum64(buf.Bytes())
		index[id.(int)] = res[i]
		*/
		if card == nil{
			continue
		}
		index[card.title] = card
		adj[card.title] = list.New()

	}
	return &graph{Adj: adj, index: index}
}

func main() {
	report := util.ParseBasedOnEmptyLine()
	// fmt.Println(len(report))
	if len(report) == 0 {
		return
	}

	graph := newGraph(report)
	graph.print()
	// pairs := graph.pair()
	// fmt.Println(pairs)
	os.Exit(0)
	/*
	part1 := 1
	for k, v := range pairs {
		fmt.Println(k, v)
		if v.Cardinality() == 2 {
			// fmt.Println(k, v)
			part1 *= k
		}
	}

	   ➜ go run 01.go
	   memo:  map[9:Set{1427, 2729} 18:Set{1171, 1489} 24:Set{1171} 43:Set{1489} 66:Set{3079} 78:Set{2971} 85:Set{2971, 2729} 89:Set{2311, 3079} 96:Set{1171} 116:Set{2473, 3079} 161:Set{2971} 177:Set{1951} 183:Set{1427, 1489} 184:
	   Set{2473, 3079} 210:Set{2311, 1427} 231:Set{2311} 234:Set{1427, 2473} 264:Set{3079} 271:Set{2729} 288:Set{1489, 1171} 300:Set{1427, 2311} 318:Set{2311, 1951} 348:Set{1427, 2473} 391:Set{1171} 397:Set{1951, 2729} 399:Set{117
	   1, 2473} 456:Set{2971} 481:Set{2473} 498:Set{2311, 1951} 501:Set{3079} 532:Set{2971} 542:Set{2473} 564:Set{1951} 565:Set{1489, 2971} 576:Set{1427, 2729} 587:Set{1951} 616:Set{2311, 3079} 680:Set{2971, 2729} 689:Set{2971, 14
	   89} 702:Set{3079} 710:Set{1951, 2729} 841:Set{1951} 848:Set{1489} 902:Set{1171} 924:Set{2311} 948:Set{1427, 1489} 962:Set{2729} 966:Set{2473, 1171}]
	   1951 Set{2729, 2311}
	   2729 Set{1427, 1951, 2971}
	   1427 Set{1489, 2473, 2311, 2729}
	   1489 Set{1427, 2971, 1171}
	   2473 Set{1427, 3079, 1171}
	   2971 Set{1489, 2729}
	   3079 Set{2473, 2311}
	   2311 Set{1951, 1427, 3079}
	   1171 Set{1489, 2473}

	fmt.Println("part1: ", part1)
		*/
}
