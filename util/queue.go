package util

import (
	"container/list"
	"fmt"
)

type Queue interface {
	Enqueue(x any)
	Dequeue() any
	//	Len() int
}

type LIFO struct {
	list.List
}

func NewLIFO() *LIFO {
	f := &LIFO{*list.New()}
	f.Init()
	return f
}

func (f *LIFO) Enqueue(x any) {
	f.PushFront(x)
}

func (f *LIFO) Dequeue() any {
	e := f.Front()
	return f.Remove(e)
}

func (f *LIFO) Reverse() *LIFO {
	r := NewLIFO()
	for e := f.Front(); e != nil; e = e.Next() {
		r.PushFront(e.Value)
	}
	return r
}

type FIFO struct {
	list.List
}

func NewFIFO() *FIFO {
	f := &FIFO{*list.New()}
	f.Init()
	return f
}

func (f *FIFO) Enqueue(x any) *list.Element {
	return f.PushBack(x)
}

func DisPlayFIFO(l *FIFO) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%c ", e.Value)
	}
	fmt.Println("")
}

func DisPlayLIFO(l *LIFO) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%c ", e.Value)
	}
	fmt.Println("")
}

func (f *FIFO) Dequeue() any {
	e := f.Front()
	if e != nil {
		return f.Remove(e)
	}
	return nil
}

func (f *FIFO) Reverse() *FIFO {
	r := NewFIFO()
	for e := f.Front(); e != nil; e = e.Next() {
		r.PushFront(e.Value)
	}
	return r
}
