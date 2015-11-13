package earthmover

import (
	"coupling"
	"github.com/stretchr/testify/assert"
	"markov"
	"testing"
	"log"
)

func setUpCouplingMatching() coupling.Coupling {
	return coupling.New()
}

func setUpMarkov() markov.MarkovChain {
	return markov.MarkovChain{
		Labels: []int{0, 1, 0, 0, 0, 1, 0},
		Transitions: [][]float64{
			[]float64{0.0, 0.33, 0.33, 0.17, 0.0, 0.17, 0.0},
			[]float64{0.0, 0.0, 0.4, 0.4, 0.0, 0.2, 0.0},
			[]float64{0.0, 0.5, 0.5, 0.0, 0.0, 0.0, 0.0},
			[]float64{0.33, 0.33, 0.34, 0.0, 0.0, 0.0, 0.0},
			[]float64{0.4, 0.4, 0.2, 0.0, 0.0, 0.0, 0.0},
			[]float64{0.0, 0.1, 0.0, 0.2, 0.5, 0.2, 0.0},
			[]float64{0.0, 0.2, 0.33, 0.0, 0.1, 0.2, 0.17}}}
}

func TestCorrectMatchingFound(t *testing.T) {
	expected := [][]float64{
		[]float64{0.33, 0.0, 0.0},
		[]float64{0.0, 0.33, 0.0},
		[]float64{0.0, 0.0, 0.17},
		[]float64{0.0, 0.0, 0.17}}
	
	c := setUpCouplingMatching()
	m := setUpMarkov()

	w := randomMatching(m, 0, 3, &c)

	for i := 0; i < len(expected); i++ {
		for j := 0; j < len(expected[0]); j++ {
			assert.True(t, approxFloatEqual(expected[i][j], w.Adj[i][j].Prob), "the correct probability were not inserted")
		}
	}
}

func TestCorrectBasicFound(t *testing.T) {
	expected := [][]bool{
		[]bool{true, true, false},
		[]bool{false, true, true},
		[]bool{false, false, true},
		[]bool{false, false, true}}
	
	c := setUpCouplingMatching()
	m := setUpMarkov()

	w := randomMatching(m, 0, 3, &c)

	for i := 0; i < len(expected); i++ {
		for j := 0; j < len(expected[0]); j++ {
			assert.Equal(t, expected[i][j], w.Adj[i][j].Basic, "the cell were not correctly set to either basic or non-basic")
		}
	}
}

func TestCorrectSuccessorFound(t *testing.T) {
	c := setUpCouplingMatching()
	m := setUpMarkov()

	w := randomMatching(m, 0, 3, &c)
	log.Println(w.Adj[2][2].To)
	assert.True(t, succNode(w, w.Adj[0][0].To.Succ), "node (0,3) did not become a successor for (0,1)")
	assert.True(t, succNode(w, w.Adj[1][1].To.Succ), "node (0,3) did not become a successor for (1,2)")
	assert.True(t, succNode(w, w.Adj[2][2].To.Succ), "node (0,3) did not become a successor for (2,3)")
	assert.True(t, succNode(w, w.Adj[3][2].To.Succ), "node (0,3) did not become a successor for (2,5)")
	assert.False(t, succNode(w, w.Adj[0][1].To.Succ), "node (0,3) become a successor for (1,1)")
	assert.False(t, succNode(w, w.Adj[1][0].To.Succ), "node (0,3) become a successor for (0,2)")
}
