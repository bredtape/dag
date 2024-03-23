package dag

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopologicalSort(t *testing.T) {
	tcs := []struct {
		G        map[string][]string
		Expected []string
		IsDAG    bool
	}{
		{
			G:        map[string][]string{"a": {"b", "c"}},
			Expected: []string{"a", "c", "b"},
			IsDAG:    true,
		},
		{
			G:        map[string][]string{"b": {"a", "c"}},
			Expected: []string{"b", "c", "a"},
			IsDAG:    true,
		},
		{
			G: map[string][]string{
				"a": nil,
				"b": nil},
			Expected: []string{"b", "a"},
			IsDAG:    true,
		},
		{
			G: map[string][]string{
				"a": {},
				"b": {"a"}},
			Expected: []string{"b", "a"},
			IsDAG:    true,
		},
		{
			G: map[string][]string{
				"a": {"b", "c"},
				"b": {"d"},
				"c": {"d"},
				"d": {"e"},
			},
			Expected: []string{"a", "c", "b", "d", "e"},
			IsDAG:    true,
		},
		// with cycles
		{
			// simple cycle
			G: map[string][]string{
				"a": {"b"},
				"b": {"c"},
				"c": {"a"},
			},
			IsDAG: false,
		},
		{
			// simple cycle
			G: map[string][]string{
				"a": {"a"},
			},
			IsDAG: false,
		},
	}

	for idx, tc := range tcs {
		t.Run(fmt.Sprintf("test case index %d", idx), func(t *testing.T) {
			dg := NewSimple(tc.G)
			actual, err := TopologicalSort(dg)

			if err != nil && tc.IsDAG {
				t.Error("graph have cycles, but expected not to have")
			} else if err == nil && !tc.IsDAG {
				t.Error("expected graph to have cycles, but it does not")
			} else {
				assert.EqualValues(t, tc.Expected, actual)
			}
		})
	}
}
