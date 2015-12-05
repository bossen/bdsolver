package matching

import (
	"github.com/stretchr/testify/assert"
	"testing"
    "utils"
    "log"
    "coupling"
)

func TestCorrectMatchingFound(t *testing.T) {
	expected := [][]float64{
		[]float64{0.33, 0.0, 0.0},
		[]float64{0.0, 0.33, 0.0},
		[]float64{0.0, 0.0, 0.17},
		[]float64{0.0, 0.0, 0.17}}
	
	c := coupling.New()
	m := coupling.SetUpMarkov()

	w := FindFeasibleMatching(m, 0, 3, &c)

	for i := 0; i < len(expected); i++ {
		for j := 0; j < len(expected[0]); j++ {
			assert.True(t, utils.ApproxEqual(expected[i][j], w.Adj[i][j].Prob), "the correct probability were not inserted")
		}
	}
}

func TestCorrectBasicFound(t *testing.T) {
	expected := [][]bool{
		[]bool{true, true, false},
		[]bool{false, true, true},
		[]bool{false, false, true},
		[]bool{false, false, true}}
	
	c := coupling.New()
	m := coupling.SetUpMarkov()

	w := FindFeasibleMatching(m, 0, 3, &c)

	for i := 0; i < len(expected); i++ {
		for j := 0; j < len(expected[0]); j++ {
			assert.Equal(t, expected[i][j], w.Adj[i][j].Basic, "the cell were not correctly set to either basic or non-basic")
		}
	}
	
	assert.Equal(t, len(w.Adj) + (len(w.Adj[0]) - 1), w.BasicCount, "the number of basic cell is not correct")
}

func TestCorrectSuccessorFound(t *testing.T) {
	c := coupling.New()
	m := coupling.SetUpMarkov()

	w := FindFeasibleMatching(m, 0, 3, &c)
	log.Println(w.Adj[2][2].To)
	assert.True(t, coupling.IsNodeInSlice(w, w.Adj[0][0].To.Succ), "node (0,3) did not become a successor for (0,1)")
	assert.True(t, coupling.IsNodeInSlice(w, w.Adj[1][1].To.Succ), "node (0,3) did not become a successor for (1,2)")
	assert.True(t, coupling.IsNodeInSlice(w, w.Adj[2][2].To.Succ), "node (0,3) did not become a successor for (2,3)")
	assert.True(t, coupling.IsNodeInSlice(w, w.Adj[3][2].To.Succ), "node (0,3) did not become a successor for (2,5)")
	assert.False(t, coupling.IsNodeInSlice(w, w.Adj[0][1].To.Succ), "node (0,3) become a successor for (1,1)")
	assert.False(t, coupling.IsNodeInSlice(w, w.Adj[1][0].To.Succ), "node (0,3) become a successor for (0,2)")
}
