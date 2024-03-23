package dag

import "slices"

// simple map implementation of the DirectedGraph interface
type Simple map[string][]string

// construct a DG from a map, making sure all nodes are represented
func NewSimple(m map[string][]string) Simple {
	for _, xs := range m {
		for _, x := range xs {
			if _, exists := m[x]; !exists {
				m[x] = nil
			}
		}
	}
	return Simple(m)
}

func (g Simple) Nodes() []string {
	keys := make([]string, 0, len(g))
	for k := range g {
		keys = append(keys, k)
	}
	// sort keys to get stable test output
	slices.Sort(keys)
	return keys
}

func (g Simple) DependsOn(node string) []string {
	return g[node]
}
