package main

import (
	"container/list"
	//"fmt"
	"fmt"
)

type GraphInterface interface {
	addNode()
	//remove_nodes() A proper algorithm for removal is required
	addEdge(interface{}, interface{})
	String()
}

type Graph struct {
	nodes int
	nodeMap map[*Node] int
	indexMap map[int] *Node
	adj []*list.List
}

func(g *Graph) init() {
	g.nodeMap = make(map[*Node] int)
	g.indexMap = make(map[int] *Node)
}

func(g *Graph) addNode(node *Node) (int, *list.Element) {
	g.nodeMap[node] = g.nodes
	g.indexMap[g.nodes] = node
	l := list.New()
	elm := l.PushBack(g.nodes)
	g.adj = append(g.adj, l)
	g.nodes++
	return g.nodes - 1, elm // return the current index at which Node is inserted
}

func(g *Graph) addEdge(from int, to int) {
	g.adj[from].PushBack(to)
}

func(g *Graph) String() string {
	var bfsGraph string
	for i :=0; i < g.nodes; i++ {
		traversed, _ := g.BFS(i)
		bfsGraph += fmt.Sprint(traversed) + "\n"
	}
	return bfsGraph
}

//func(g *Graph) removeEdge(from *Node, to *Node) {
//	fromIndex := g.nodeMap[from]
//	toIndex := g.nodeMap[to]
//	g.adj[fromIndex].Remove(toIndex)
//}

