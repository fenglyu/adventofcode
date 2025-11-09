package main

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
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

func (m *card) flip() *card {
	var flipc card

	return &flipc
}

// rotate
func (m *card) rotate() *card {
	var flipc card

	return &flipc
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
	return fmt.Sprintf("Title: %d, data: %s, edges: %v", m.title, util.Matrix2Str(m.data), m.edges)
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
		fmt.Printf("key: %d, val: \n", k)
		// util.PrintMatrix(v.data)
		for e := v.Front(); e != nil; e = e.Next() {
			fmt.Printf("%v \t", e.Value)
		}

		fmt.Println("\n")
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

func (ma *graph) pair() {
	for tite, c := range ma.index {
		card := c
		card.setEdges()
		ma.index[tite] = card
	}

	memo := make(map[int]set.Set[int])

	for _, v := range ma.index {
		// fmt.Println("k: ", k, "v: ", v)
		for _, e := range v.edges {
			if _, Ok := memo[e]; !Ok {
				memo[e] = set.NewSet[int]()
			}
			memo[e].Add(v.title)
		}
	}

	fmt.Println("memo: ")
	spew.Dump(memo)

	// ma.Adj
	for _, v := range memo {
		if v.Cardinality() < 2 {
			// skip card's edge that has no neighbours
			continue
		}

		for _, q := range v.ToSlice() {
			for _, p := range v.ToSlice() {
				if q == p {
					continue
				}

				if _, Ok := ma.Adj[q]; !Ok {
					ma.Adj[q] = list.New()
				}

				if !eleInList(p, ma.Adj[q]) {
					ma.Adj[q].PushBack(p)
				}
			}
		}
	}

	return
}

// eleInList return true if v already exists in given linked list
func eleInList(v any, l *list.List) bool {
	for e := l.Front(); e != nil; e = e.Next() {
		val := v.(int)
		if val == e.Value.(int) {
			return true
		}
	}

	return false
}

func newGraph(raw []string) *graph {
	if len(raw) == 0 {
		return nil
	}
	// res := make([]card, len(raw))
	adj := make(map[int]*list.List)
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
		if card == nil {
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
	// graph.print()
	graph.pair()
	spew.Dump(graph.index)
	graph.print()
}

/*
Understanding the current state
pair() in 2020/20/02.go is already turning each tile into a set of matching neighbors by comparing every edge (both orientations) and populating graph.Adj[title] with IDs of tiles that can touch it.

Each card now carries all of its edges, forward and reversed, courtesy of setEdges, so you have the raw data needed to reason about rotations and flips later.

Next steps for assembling the image
Classify tiles first (corners, edges, interiors).
With the adjacency map, count how many neighbors every tile has: 2 means “corner”, 3 means “edge”, and 4 means “interior”. That lets you pick a corner tile to anchor the layout and narrows search when you fill borders, because only edge tiles can go there.

Represent board positions explicitly.
Treat the final picture as an N x N grid of tile IDs. You can compute N = sqrt(len(graph.index)), then run a backtracking search that tries to place tiles row by row. Track which IDs you have already used and consult adjacency sets so you only consider candidates that actually share an edge with already-placed neighbors.

Store oriented edges, not just IDs.
To know whether a candidate tile fits to the right/below another tile, you need the literal edge strings (or their integer encodings). Build a helper that, given a tile ID and a rotation/flip choice, returns the four oriented edges. You already have the base edges; adding rotation and mirroring helpers on top of card.data is the next milestone.

Backtracking search with pruning.
For each position (r, c) in the grid:

derive the required top edge (from (r-1, c)) and left edge (from (r, c-1)) if they exist;

iterate neighbors from graph.Adj that are still unused;

for each, iterate the eight orientations (4 rotations × 2 flips) and check whether the oriented edges align.
Once all constraints match, place the tile and recurse to the next position. Backtrack if you hit a dead end.

Optional: use BFS/DFS only for connectedness, not placement.
Graph traversal algorithms (BFS/DFS) are good for confirming the graph is a single component, but they can’t on their own assign absolute positions because rotations matter. A constraint-satisfaction/backtracking solver is the usual approach here.

On the adjacency list structure
map[int]*list.List works, but it forces linear scans (eleInList) to keep entries unique.
Consider storing map[int]set.Set[int] or simply map[int][]int and rely on Go’s slices plus a seen map during generation. A slice is more idiomatic and simpler, given that neighbor counts are tiny (≤4).

Regardless of container, keep the adjacency lookup O(1) per tile ID, because you will query it constantly during the search.

Putting this together, your pipeline can look like:

Build the adjacency map you already have.

Split tiles into corner/edge/interior buckets.

Drive a recursive backtracking solver that fills an N x N grid, testing all orientations.

Once the grid is filled, strip borders (stripBoarder) and stitch the tile cores together to form the large image, then continue with sea-monster detection.

Following that structure will get you from the neighbor graph to a full puzzle assembly.
*/
