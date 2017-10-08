package main

import (
	"errors"
)

func(g *Graph) BFS(nodeIndex int) ([]Node, error) {
	if g.nodes < nodeIndex {
		return nil, errors.New("Invalid node Index")
	}
	queue := []int{nodeIndex}
	result := []Node{}
	visited := map[int]bool{}
	visited[nodeIndex] = true
	for len(queue) > 0 {
		result = append(result, *g.indexMap[queue[0]])
		visited[queue[0]] = true
		currList := g.adj[queue[0]]
		for e := currList.Front(); e != nil; e = e.Next() {
			if !visited[e.Value.(int)] {
				queue = append(queue, e.Value.(int))
			}
		}
		queue = queue[1:]
	}
	return result, nil
}