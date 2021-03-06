package lazy_prime_mst

import (
	"github.com/dczombera/data-structures-and-algorithms-in-go/algorithms/graph/minimum_spanning_tree/priority_queue"
	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/edge_weighted_graph"
	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/edge_weighted_graph/edge"
)

// LazyPrimeMST is a data type for computing the minimum spanning tree/forest in an edge weighted undirected graph using a lazy version of Prim’s algorithm with a binary heap
type LazyPrimeMST struct {
	mst    []edge.Edge
	weight float64
	pq     priority_queue.MinPriorityQueue
	marked []bool
}

var initCap = 8

func NewLazyPrimeMST(g *edge_weighted_graph.EdgeWeightedGraph) *LazyPrimeMST {
	mst := &LazyPrimeMST{make([]edge.Edge, 0, initCap), 0.0, priority_queue.NewMinPriorityQueue(), make([]bool, g.VerticesCount())}
	for v := 0; v < g.VerticesCount(); v++ {
		if !mst.marked[v] {
			// Run from each vertex to find minimum spanning forest
			mst.prime(g, v)
		}
	}

	return mst
}

func (mst *LazyPrimeMST) prime(g *edge_weighted_graph.EdgeWeightedGraph, s int) {
	mst.scan(g, s)
	for !mst.pq.IsEmpty() {
		min, err := mst.pq.DelMin()
		if err != nil {
			panic(err)
		}

		v := min.Either()
		w := min.Other(v)
		// ignore edge if both vertices are already in the tree
		if mst.marked[v] && mst.marked[w] {
			continue
		}
		mst.mst = append(mst.mst, min)
		mst.weight += min.Weight()

		if !mst.marked[v] {
			mst.scan(g, v)
		}
		if !mst.marked[w] {
			mst.scan(g, w)
		}
	}
}

func (mst *LazyPrimeMST) scan(g *edge_weighted_graph.EdgeWeightedGraph, v int) {
	mst.marked[v] = true
	for _, e := range g.AdjacencyList(v) {
		if !mst.marked[e.Other(v)] {
			mst.pq.Insert(e)
		}
	}
}

func (mst *LazyPrimeMST) Weight() float64 {
	return mst.weight
}

func (mst *LazyPrimeMST) Edges() []edge.Edge {
	return mst.mst
}
