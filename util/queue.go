package util

import (
	"container/list"
	"fmt"
)

type Queue interface {
	Enqueue(x interface{})
	Dequeue() interface{}
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

func (f *LIFO) Enqueue(x interface{}) {
	f.PushFront(x)
}

func (f *LIFO) Dequeue() interface{} {
	e := f.Front()
	return f.Remove(e)
}

type FIFO struct {
	list.List
}

func NewFIFO() *FIFO {
	f := &FIFO{*list.New()}
	f.Init()
	return f
}

func (f *FIFO) Enqueue(x interface{}) *list.Element {
	return f.PushBack(x)
}

func DisPlay(l *FIFO) {
	fmt.Println("l: ", l)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func (f *FIFO) Dequeue() interface{} {
	e := f.Front()
	if e != nil {
		return f.Remove(e)
	}
	return nil
}
