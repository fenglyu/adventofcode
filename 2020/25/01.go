package main

import (
	"container/list"
	"fmt"

	"github.com/fenglyu/adventofcode/util"
)

type dimen struct {
	//z int
	//w int
	vec    []int
	count  int
	values [][]int
}

func (d *dimen) String() string {
	return fmt.Sprintf("[vectors = %v, count = %d, values = %v]", d.vec, d.count, d.values)
}

func main() {

	report := util.ParseBasedOnEachLine()
	fmt.Println(report)
	/*
		a := []interface{}{1, 2, 3}
		b := []interface{}{"a", "b", "c"}

		c := cartesian.Iter(a, b)

		// receive products through channel
		for product := range c {
			fmt.Println(len(product))
		}


			d := []interface{}{-1, 0, 1}
			e := []interface{}{-1, 0, 1}
			//f := []interface{}{-1, 0, 1}
			//g := []interface{}{-1, 0, 1}


				h := cartesian.Iter(d, e) //, f, g)
				count := 0
				for product := range h {
					fmt.Println(product)
					count++
				}

				fmt.Println(count)
	*/

	cycle := 1
	for cycle <= 6 {
		//	arr := make([]int, 0)
		l := list.New()
		i := 0
		for i < cycle {
			switch i {
			case 0:
				l.PushFront(i)
			default:
				l.PushFront(-i)
				l.PushBack(i)
			}
			i++
		}

		//sort.Ints(arr)
		fmt.Println(l)

		arr := make([]int, 0)
		for e := l.Front(); e != nil; e = e.Next() {
			arr = append(arr, e.Value.(int))
		}

		fmt.Println(arr)

		hypeVec := util.ProductGen(arr, 2, false)
		fmt.Println(hypeVec)
		/*
			for i := 0; i < len(arr); i++ {
				arr[i] += arr[len(arr)-1]
			}
			hypeVec = util.ProductGen(arr, 2)
			fmt.Println(hypeVec)

			for _, v := range hypeVec {
				fmt.Println(v)
			}
		*/
		cycle++
	}

	hypeVec := util.ProductGen([]int{-1, 0, 1}, 4, false)
	for _, v := range hypeVec {
		fmt.Println(v)
	}
}
