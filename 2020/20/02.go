package main

import (
	"container/list"
	"fmt"
	"math"
	"sort"
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

func (m *card) flip() *card {
	var flipc card

	return &flipc
}

// rotate
func (m *card) rotate() *card {
	var flipc card

	return &flipc
}

type Transform func(r, c, n int) (int, int)

var RotRight90 Transform = func(r, c, n int) (int, int) {
	return c, n - 1 - r
}

var RotLeft90 Transform = func(r, c, n int) (int, int) {
	return n - 1 - c, r
}

var Rot180 Transform = func(r, c, n int) (int, int) {
	return n - 1 - r, n - 1 - c
}

var Rotations = []Transform{RotRight90, Rot180, RotLeft90}

func rotateN(r, c, n, times int) (int, int) {
	t := ((times % 4) + 4) % 4
	for i := 0; i < t; i++ {
		r, c = Rotations[0](r, c, n) // right rotation
	}
	return r, c
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
	// Grid
	Grid [][]*card
	// Adj []*list.List
	Adj map[int]*list.List
	// data  []card
	index map[int]*card
}

type ByAdjBorderNum struct {
	Adj   map[int]*list.List
	Order []int
}

// Return the Graph's Adj to order by tile's boarders number
func NewByAdjBorderNum(Adj map[int]*list.List) *ByAdjBorderNum {
	orders := make([]int, 0, len(Adj))
	for k := range Adj {
		orders = append(orders, k)
	}
	return &ByAdjBorderNum{Adj: Adj, Order: orders}
}

func (b *ByAdjBorderNum) Less(i, j int) bool {
	ai, bj := b.Order[i], b.Order[j]
	// order in desc
	return b.Adj[ai].Len() > b.Adj[bj].Len()
}

func (b *ByAdjBorderNum) Swap(i, j int) {
	// ai, bj := b.Order[i], b.Order[j]
	b.Order[i], b.Order[j] = b.Order[j], b.Order[i]
}

func (b *ByAdjBorderNum) Len() int {
	return len(b.Adj)
}

func (ma *graph) print() {
	for k, v := range ma.Adj {
		fmt.Printf("key: %d, val: ", k)
		// util.PrintMatrix(v.data)
		for e := v.Front(); e != nil; e = e.Next() {
			fmt.Printf("%v\t", e.Value)
		}

		fmt.Println("\n")
	}
}

func (ma *graph) pair() {
	for tite, c := range ma.index {
		card := c
		card.setEdges()
		ma.index[tite] = card
	}

	// memo := make(map[int]set.Set[int])
	// how many card share which edge
	// edge -> {card0, card1, ...}
	memo := make(map[int]map[int]bool) // directly use {'a': true, 'b': true} as set{'a', 'b'}

	for _, v := range ma.index {
		// fmt.Println("k: ", k, "v: ", v)
		for _, e := range v.edges {
			if _, Ok := memo[e]; !Ok {
				// memo[e] = set.NewSet[int]()
				memo[e] = make(map[int]bool)
			}
			// memo[e].Add(v.title)
			memo[e][v.title] = true
		}
	}

	// fmt.Println("memo: ")
	// spew.Dump(memo)

	// ma.Adj
	for _, v := range memo {
		// if v.Cardinality() < 2 {
		if len(v) < 2 {
			// skip card's edge that has no neighbours
			continue
		}

		for q := range v {
			for p := range v {
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
	gridSize := int(math.Sqrt(float64(len(raw))))
	grid := make([][]*card, gridSize)
	adj := make(map[int]*list.List)
	index := make(map[int]*card, len(raw))
	for _, v := range raw {
		card := newMtx(v)
		if card == nil {
			continue
		}
		index[card.title] = card
		adj[card.title] = list.New()

	}
	return &graph{Adj: adj, index: index, Grid: grid}
}

// placement will ajust the card's 2d placement in the grid
func (ma *graph) PlaceMent() {
	// > https://www.geeksforgeeks.org/artificial-intelligence/explain-the-concept-of-backtracking-search-and-its-role-in-finding-solutions-to-csps/
	// contraint satification problem(CSP)
	// Variables: a set of variables, X1.. Xn
	// Domains: Each variable Xi has a domain Di of possible values
	// Constraits: A set of contraints that specify allowable combinations of values for subsets of values

	// The goal in a CSP is to assign values to all variables from their respective domains such that all constraints are satisfied.
	/*
		Sudoku: Filling a 9x9 grid with digits so that each row, column, and 3x3 subgrid contains all digits from 1 to 9 without repetition.
		Map Coloring: Coloring a map with a limited number of colors so that no adjacent regions share the same color.
		N-Queens: Placing N queens on an N×N chessboard so that no two queens threaten each other.
	*/

	// 1. init the board
	for i := range ma.Grid {
		ma.Grid[i] = make([]*card, len(ma.Grid))
	}

	// 2. Solve the problem
	N := 0
	res := Solve(ma.Grid, 0, N)
	if res {
		PrintBoard(ma.Grid, N)
	} else {
		fmt.Println("no solution exists")
	}

	// 3.
}

// isRightPlace define is exsting card are in right order
func IsRightPlace(board [][]*card, row int, col int, N int) bool {
	return true
}

func Solve(board [][]*card, tile int, N int) bool {
	// if the card(with tile) can't be placed in any places, return false

	return false
}

func PrintBoard(board [][]*card, N int) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("[%d, %d]-> card: %v", i, j, board[i][j])
		}
		fmt.Println("")
	}
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
	// spew.Dump(graph.index)
	// graph.print()
	// problem 1
	prod := 1

	for t, v := range graph.Adj {
		if v.Len() == 2 {
			prod *= t
		}
	}

	fmt.Printf("problem 1: %d\n", prod)

	// problem 2
	bdNum := NewByAdjBorderNum(graph.Adj)
	sort.Sort(bdNum)

	fmt.Println("boarder number: ", bdNum)
	for _, tile := range bdNum.Order {
		fmt.Printf("tile : %d\t", tile)
		fmt.Printf("Adj: ")
		adj := graph.Adj[tile]
		for e := adj.Front(); e != nil; e = e.Next() {
			fmt.Printf("%v\t", e.Value)
		}
		fmt.Println("")
	}

	// backtracking
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
