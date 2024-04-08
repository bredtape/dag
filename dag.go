package dag

import (
	"errors"
	"slices"

	"github.com/bredtape/set"
)

var (
	ErrGraphHaveCycles = errors.New("the DAG have cycles")
)

type DirectedGraph[T comparable] interface {
	// get all nodes present
	Nodes() []T

	// get a list of nodes that the argument node depends on
	DependsOn(T) []T
}

// sort the nodes is topological order
// returns ErrGraphHaveCycles if the graph have cycles
func TopologicalSort[T comparable](g DirectedGraph[T]) ([]T, error) {
	nodes := g.Nodes()
	visited := set.New[T](len(nodes))
	order := make([]T, 0, len(nodes))

	for _, n := range nodes {
		xs, err := DFS(g, n)
		if err != nil {
			return nil, err
		}

		for _, x := range xs {
			if !visited.Contains(x) {
				order = append(order, x)
			}
			visited.Add(x)
		}

		if visited.Count() == len(nodes) {
			break
		}
	}

	slices.Reverse(order)
	return order, nil
}

// depth first search with cycle detection
func DFS[T comparable](g DirectedGraph[T], startingNode T) ([]T, error) {
	visited := set.New[T](4)
	marked := set.New[T](4)
	return dfs(g, visited, marked, startingNode)
}

func dfs[T comparable](g DirectedGraph[T], visited, marked set.Set[T], node T) ([]T, error) {
	order := make([]T, 0)
	visited.Add(node)
	marked.Add(node)
	for _, n := range g.DependsOn(node) {
		// we came back to a predecessor node already visited
		if marked.Contains(n) {
			return nil, ErrGraphHaveCycles
		}

		if visited.Contains(n) {
			continue
		}

		xs, err := dfs(g, visited, marked, n)
		if err != nil {
			return nil, err
		}
		order = append(order, xs...)
	}

	marked.Remove(node)
	order = append(order, node)
	return order, nil
}
