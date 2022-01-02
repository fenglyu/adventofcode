package util

import (
	"math/rand"
	"testing"
)

func TestLIFOEnqueue(t *testing.T) {
	f := NewLIFO()
	arr := []int{5, 4, 3}
	for _, val := range arr {
		f.Enqueue(val)
	}
	for i := range arr {
		v := f.Dequeue()
		if v != arr[len(arr)-1-i] {
			t.Fatalf("expected %v, found %v\n", arr[len(arr)-1-i], v)
		}
	}
}

//  go test -timeout 30s -run ^TestFIFODequeue$ github.com/fenglyu/adventofcode/util -v
func TestFIFODequeue(t *testing.T) {
	f := NewFIFO()
	//f.Init()
	arr := []int{5, 4, 3}
	for _, val := range arr {
		f.Enqueue(val)
	}
	for i := range arr {
		v := f.Dequeue()
		if v != arr[i] {
			t.Fatalf("expected %v, found %v\n", arr[i], v)
		}
	}
}

func TestLIFOCorrectNess(t *testing.T) {
	f := NewLIFO()
	var MAX int = 10000
	arr := make([]int, MAX)
	for i := 0; i < MAX; i++ {
		v := rand.Intn(MAX)
		f.Enqueue(v)
		arr[i] = v
	}

	for i := 0; i < MAX; i++ {
		v := f.Dequeue()
		if v != arr[len(arr)-1-i] {
			t.Fatalf("expected %v, found %v\n", arr[len(arr)-1-i], v)
		}
	}

}

func TestFIFOCorrectNess(t *testing.T) {
	f := NewFIFO()
	var MAX int = 10000
	arr := make([]int, MAX)
	for i := 0; i < MAX; i++ {
		v := rand.Intn(MAX)
		f.Enqueue(v)
		arr[i] = v
	}

	for i := 0; i < MAX; i++ {
		v := f.Dequeue()
		if v != arr[i] {
			t.Fatalf("expected %v, found %v\n", arr[i], v)
		}
	}
}
