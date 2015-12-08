package tpsolverdefault

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"utils"
	"coupling"
	"matching"
	"setpair"
	"sets"
    "uvmethod"
)

func TestOptimalSolutionFound(t *testing.T) {
	c, m, visited, exact, d := coupling.SetUpTest()
	d = sets.InitD(len(m.Transitions))
	
	w := matching.FindFeasibleMatching(m, 0, 3, &c)
	setpair.Setpair(m, w, exact, visited, d, &c)
	
	// the expected solution is not precice since the markov chain used do not use precise fractions
	expected := [][]float64{
		[]float64{0, 0.33, 0},
		[]float64{0, 0, 0.33},
		[]float64{0.17, 0, 0},
		[]float64{0.16, 0, 0.01}}
		
	min, i, j := uvmethod.Run(w, d)
	
	Solve(m, w, d, min, i, j)
	
	for i := range(w.Adj) {
		for j := range(w.Adj[0]) {
			assert.True(t, utils.ApproxEqual(expected[i][j], w.Adj[i][j].Prob), "the optimal probability found was not what we expected")
		}
	}
}
