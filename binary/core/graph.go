package main

import "container/list"

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

func(g *Graph) addNode(node *Node) {
	g.nodeMap[node] = g.nodes
	g.indexMap[g.nodes] = node
	l := list.New()
	l.PushBack(g.nodes)
	g.adj = append(g.adj, l)
	g.nodes++
}

func(g *Graph) addEdge(from *Node, to *Node) {
	fromIndex := g.nodeMap[from]
	toIndex := g.nodeMap[to]
	g.adj[fromIndex].PushBack(toIndex)
}

//func(g *Graph) removeEdge(from *Node, to *Node) {
//	fromIndex := g.nodeMap[from]
//	toIndex := g.nodeMap[to]
//	g.adj[fromIndex].Remove(toIndex)
//}

func main() {

}
