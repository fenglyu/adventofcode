import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/fenglyu/adventofcode/util"
)

var Card2Val map[uint8]uint8
var Val2Card map[uint8]uint8

//type uint8 uint8

var (
	CARDS      = []uint8{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}
	CARDSValue = []uint8{0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xA, 0xB, 0xC}
)

type HandType int

const (
	High HandType = iota
	One
	Two
	Three
	Full
	Four
	Five
)

type Hand struct {
	Cards      string
	Bid        int
	Type       HandType
	Largest    uint8
	Second     uint8
	Index      []uint8
	Occurrence []uint8
	Rest       []uint8
}

/*
	func (h *Hand) String() string {
		return fmt.Sprintf("Cards: %s, Bid %d, Type %v, Rest: %v", h.Cards, h.Bid, h.Type, h.Rest)
	}
*/
func (h *Hand) String() string {
	restCards := make([]uint8, len(h.Rest))
	for i, val := range h.Rest {
		restCards[i] = Val2Card[val]
	}
	return fmt.Sprintf("%s %d", h.Cards, h.Bid)
	//return fmt.Sprintf("Cards: %s, Bid: %d, Type: %v, largest: %d, second: %v, Rest: %v", h.Cards, h.Bid, h.Type, h.Largest, h.Second, string(restCards))
}

func InitHand(s string) *Hand {
	// 32T3K 765
	sl := strings.Fields(s)
	hs, bidstr := sl[0], sl[1]
	val := make([]uint8, len(hs))
	for i, v := range []byte(hs) {
		val[i] = Card2Val[v]
	}
	bid, _ := strconv.Atoi(bidstr)
	h := Hand{Cards: hs, Bid: bid}
	h.SetType()
	return &h
}

type Uint8Slice []uint8

func (p Uint8Slice) Len() int           { return len(p) }
func (p Uint8Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Uint8Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type IndexOccurrence struct {
	Index     uint8
	Occurence uint8
}

func generateLists(bucket []uint8) ([]uint8, []uint8, []uint8) {
	var pairs []IndexOccurrence
	var singleOccurrences []uint8

	for i, v := range bucket {
		if v > 0 {
			pairs = append(pairs, IndexOccurrence{Index: uint8(i), Occurence: v})
			if v == 1 {
				singleOccurrences = append(singleOccurrences, uint8(i))
			}
		}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Occurence > pairs[j].Occurence || (pairs[i].Occurence == pairs[j].Occurence && pairs[i].Index > pairs[j].Index)
	})

	// Extract indices and occurrences
	var indices []uint8
	var occurrences []uint8
	for _, pair := range pairs {
		indices = append(indices, pair.Index)
		occurrences = append(occurrences, pair.Occurence)
	}

	return indices, occurrences, singleOccurrences
}

// func (h *Hand) SetType() (HandType, uint8, uint8, []uint8) {
func (h *Hand) SetType() {
	bucket := make([]uint8, len(CARDSValue))
	for _, v := range []byte(h.Cards) {
		bucket[Card2Val[v]]++
	}

	//lval, sval := uint8(0), uint8(0)
	//rest := make([]uint8, 0)
	indices, occurrences, rest := generateLists(bucket)
	//fmt.Println(h.Cards)
	//fmt.Println(indices)
	//fmt.Println(occurrences)
	lval, sval := uint8(0), uint8(0)
	lval = occurrences[0]
	h.Largest = indices[0]
	if len(occurrences) > 1 {
		sval = occurrences[1]
		h.Second = indices[1]
	}

	var t HandType
	switch {
	case lval == 5:
		t = Five
	case lval == 4:
		t = Four
	case lval == 3 && sval == 2:
		t = Full
	case lval == 3 && sval == 1:
		t = Three
	case lval == 2 && sval == 2:
		t = Two
	case lval == 2 && sval == 1:
		t = One
	default:
		t = High
	}

	h.Type = t
	h.Occurrence = occurrences
	h.Index = indices

	sort.Sort(sort.Reverse(Uint8Slice(rest)))
	// the indeices lval and sval corresponding to
	h.Rest = rest
}

func reverse(slice []uint8) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

type CarmelCards struct {
	hands []*Hand
}

func ArrLess(a []uint8, b []uint8) bool {
	if len(a) > len(b) {
		a, b = b, a
	}

	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			continue
		}
		return a[i] < b[i]
	}
	return false
}

func (c *CarmelCards) Less(i int, j int) bool {
	a, b := c.hands[i], c.hands[j]
	at, alval, asval, arest := a.Type, a.Largest, a.Second, a.Rest
	bt, blval, bsval, brest := b.Type, b.Largest, b.Second, b.Rest

	if at == bt {
		if at == Five {
			return alval < blval
		} else if at == Four {
			if alval == blval {
				//return ArrLess(arest, brest)
				return a.Cards < b.Cards
			}
			return alval < blval
		} else if at == Full {
			if alval == blval {
				return asval < bsval
			}
			return alval < blval
		} else if at == Three {
			if alval == blval {
				//	return ArrLess(arest, brest)
				return a.Cards < b.Cards
			}
			return alval < blval
		} else if at == Two {
			if alval == blval {
				//				return asval < bsval
				if asval == bsval {
					//	return ArrLess(arest, brest)
					return a.Cards < b.Cards
				}
				return asval < bsval
			}
			return alval < blval
		} else if at == One {
			if alval == blval {
				//return ArrLess(arest, brest)
				return a.Cards < b.Cards
			}
			return alval < blval
		} else if at == High {
			//return ArrLess(arest, brest)
			return a.Cards < b.Cards
		}
	}
	return at < bt
}

func (c *CarmelCards) Len() int {
	return len(c.hands)
}

func (c *CarmelCards) Swap(i, j int) {
	c.hands[i], c.hands[j] = c.hands[j], c.hands[i]
}

func main() {
	startTime := time.Now()
	//fmt.Println("CARDS: ", CARDS)
	//fmt.Println("CARDSValue: ", CARDSValue)
	Card2Val = make(map[uint8]uint8)
	Val2Card = make(map[uint8]uint8)
	for i := 0; i < len(CARDS); i++ {
		Card2Val[CARDS[i]] = CARDSValue[i]
		Val2Card[CARDSValue[i]] = CARDS[i]
	}
	//fmt.Println(Card2Val)
	//fmt.Println(Val2Card)
	//t := InitHand("KTJJT 220")
	//fmt.Println(t)
	raw := util.ParseBasedOnEachLine()

	var c CarmelCards
	hands := make([]*Hand, len(raw))
	for i, v := range raw {
		h := InitHand(v)
		hands[i] = h
	}
	c.hands = hands

	sort.Sort(&c)

	var power uint64 = 0
	for i := len(c.hands) - 1; i >= 0; i-- {
		//fmt.Println(c.hands[i], i+1, c.hands[i].Bid)
		fmt.Println(c.hands[i])
		power += uint64(i+1) * uint64(c.hands[i].Bid)
	}
	//fmt.Println(raw)
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Println("Problem 1: ", power)
	fmt.Println("Execution time: ", duration)
}
