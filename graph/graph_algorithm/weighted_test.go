package graph_algorithm_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/elecbug/go-dspkg/graph"
	ga "github.com/elecbug/go-dspkg/graph/graph_algorithm"
	"github.com/elecbug/go-dspkg/graph/graph_type"
)

func TestDirectedWeighted(t *testing.T) {
	cap := 100
	g := graph.NewGraph(graph_type.DIRECTED_WEIGHTED, cap)

	for i := 0; i < cap; i++ {
		g.AddNode()
	}

	for i := 0; i < cap*3; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano() + int64(i)))
		from := graph.NodeID(r.Intn(g.NodeCount()))

		r = rand.New(rand.NewSource(time.Now().UnixNano() + int64(i*i)))
		to := graph.NodeID(r.Intn(g.NodeCount()))

		r = rand.New(rand.NewSource(time.Now().UnixNano() + int64(i*i)))
		g.AddWeightEdge(from, to, graph.Distance(r.Intn(20)))
	}

	t.Log("ready")
	// t.Logf("\n%s\n", g.String())

	{
		u := ga.NewUnit(g)
		s := time.Now()

		t.Logf("\nAverageShortestPathLength: %v\n", spew.Sdump(u.AverageShortestPathLength()))
		t.Logf("\nShortestPath: %v\n", spew.Sdump(u.ShortestPath(0, graph.NodeID(cap-1))))
		t.Logf("\nShortestPath: %v\n", spew.Sdump(u.ShortestPath(graph.NodeID(cap-1), 0)))

		duration := time.Since(s)
		t.Logf("execution time: %s", duration)
	}
	{
		u := ga.NewParallelUnit(g, 5)
		s := time.Now()

		t.Logf("\nAverageShortestPathLength: %v\n", spew.Sdump(u.AverageShortestPathLength()))
		t.Logf("\nShortestPath: %v\n", spew.Sdump(u.ShortestPath(0, graph.NodeID(cap-1))))
		t.Logf("\nShortestPath: %v\n", spew.Sdump(u.ShortestPath(graph.NodeID(cap-1), 0)))

		duration := time.Since(s)
		t.Logf("execution time: %s", duration)
	}
}

func TestUndirectedWeighted(t *testing.T) {
	cap := 100
	g := graph.NewGraph(graph_type.UNDIRECTED_WEIGHTED, cap)

	for i := 0; i < cap; i++ {
		g.AddNode()
	}

	for i := 0; i < cap*3; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano() + int64(i)))
		from := graph.NodeID(r.Intn(g.NodeCount()))

		r = rand.New(rand.NewSource(time.Now().UnixNano() + int64(i*i)))
		to := graph.NodeID(r.Intn(g.NodeCount()))

		r = rand.New(rand.NewSource(time.Now().UnixNano() + int64(i*i)))
		g.AddWeightEdge(from, to, graph.Distance(r.Intn(20)))
	}

	t.Log("ready")
	// t.Logf("\n%s\n", g.String())

	{
		u := ga.NewUnit(g)
		s := time.Now()

		t.Logf("\nAverageShortestPathLength: %v\n", spew.Sdump(u.AverageShortestPathLength()))
		t.Logf("\nShortestPath: %v\n", spew.Sdump(u.ShortestPath(0, graph.NodeID(cap-1))))
		t.Logf("\nShortestPath: %v\n", spew.Sdump(u.ShortestPath(graph.NodeID(cap-1), 0)))

		duration := time.Since(s)
		t.Logf("execution time: %s", duration)
	}
	{
		u := ga.NewParallelUnit(g, 5)
		s := time.Now()

		t.Logf("\nAverageShortestPathLength: %v\n", spew.Sdump(u.AverageShortestPathLength()))
		t.Logf("\nShortestPath: %v\n", spew.Sdump(u.ShortestPath(0, graph.NodeID(cap-1))))
		t.Logf("\nShortestPath: %v\n", spew.Sdump(u.ShortestPath(graph.NodeID(cap-1), 0)))

		duration := time.Since(s)
		t.Logf("execution time: %s", duration)
	}
}
