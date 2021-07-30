package main

import (
	"fmt"
	"strings"

	"github.com/fenglyu/adventofcode/util"
	//"github.com/hashicorp/terraform/dag"
)

//var mapp map[int]*rule
var mapp map[string]interface{}

func main() {

	report := util.ParseBasedOnEmptyLine()
	//fmt.Println(report, len(report))

	mapp = make(map[string]interface{})
	for _, v := range strings.Split(report[0], "\n") {
		r := strings.Split(v, ": ")
		idx, res := r[0], r[1]
		if strings.Contains(res, "|") {
			mapp[idx] = newRule(res)
		} else if strings.Contains(res, "\"") {
			mapp[idx] = strings.Trim(res, "\"")
		} else {

		}
	}

	fmt.Println(mapp)

}

type rule struct {
	a, b []string
}

func (r *rule) String() string {
	return fmt.Sprintf("%v | %v", r.a, r.b)
}

func newRule(res string) *rule {
	r := strings.Split(res, " | ")
	a, b := strings.Split(r[0], " "), strings.Split(r[1], " ")
	return &rule{a: a, b: b}
}

func validate(target []byte, idx int) bool {

	return false
}
