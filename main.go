package main

import "fmt"

func main() {
	g := new(Graph)
	g.addNode(&Node{value:1})
	g.addNode(&Node{value:2})
	g.addNode(&Node{value:3})
	g.addNode(&Node{value:4})
	g.addEdge(0, 1)
	g.addEdge(2, 3)
	g.addEdge(1, 3)
	g.addEdge(2, 0)
	fmt.Print(g)
}
