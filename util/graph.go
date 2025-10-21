package util

import (
	"container/list"
	"fmt"
)

/*
type Node struct {
	head *list.List
}

func newNode() *Node{
	return &Node{head: list.New()}
}
*/
// https://go.dev/play/p/n-vgwP6Xlix
type Graph struct {
	numVertices int
	adjList  []*list.List
	isDirected bool
}

func createGraph(vertices int ,isDirected bool ) *Graph{
	adjlist:=make([]*list.List,0, vertices)
	for i:=0; i<vertices; i++{
		adjlist[i] = list.New()
	}

	g := Graph{
		numVertices: vertices,
		adjList: adjlist,
		isDirected: isDirected,
	}
	return &g
}
  

func (g *Graph) addEdge(dst any, src any){
	// add edge from src to dst

}




func main(){

	vertices := 5
	isDirected :=false
	g := createGraph(vertices, isDirected)

	fmt.Println(g)
}