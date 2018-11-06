package bfs

import (
	"fmt"
	"testing"

	"github.com/Rongcong/algorithms/data-structures/graph"
)

func TestBfs(t *testing.T) {
	h := graph.NewDirected()

	for i := 0; i < 10; i++ {
		v := graph.VertexId(i)
		h.AddVertex(v)
	}

	for i := 0; i < 9; i++ {
		h.AddEdge(graph.VertexId(i), graph.VertexId(i+1), 1)
	}

	bfsMap := make(map[graph.VertexId]bool)
	checkVertices := func(v graph.VertexId) {
		bfsMap[v] = true
	}

	Bfs(h, graph.VertexId(3), checkVertices)

	for i := 3; i < len(bfsMap); i++ {
		if _, ok := bfsMap[graph.VertexId(i)]; !ok {
			fmt.Println(bfsMap)
			t.Error()
		}
	}
}
