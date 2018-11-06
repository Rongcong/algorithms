package dfs

import (
	"github.com/Rongcong/algorithms/data-structures/graph"
	"github.com/Rongcong/algorithms/data-structures/stack"
)

// UndirectedDfs ...
func UndirectedDfs(g *graph.UnGraph, v graph.VertexId, fn func(graph.VertexId)) {
	s := stack.New()
	s.Push(v)
	visited := make(map[graph.VertexId]bool)

	for s.Len() > 0 {
		v := s.Pop().(graph.VertexId)

		if _, ok := visited[v]; !ok {
			visited[v] = true
			fn(v)
			neighbours := g.GetNeighbours(v).VerticesIter()
			for neighbour := range neighbours {
				s.Push(neighbour)
			}
		}
	}
}

// DirectedDfs ...
func DirectedDfs(g *graph.DirGraph, v graph.VertexId, fn func(graph.VertexId)) {
	s := stack.New()
	s.Push(v)
	visited := make(map[graph.VertexId]bool)

	for s.Len() > 0 {
		v = s.Pop().(graph.VertexId)

		if _, ok := visited[v]; !ok {
			visited[v] = true
			fn(v)
			neighbours := g.GetSuccessors(v).VerticesIter()
			for neighbour := range neighbours {
				s.Push(neighbour)
			}
		}
	}
}
