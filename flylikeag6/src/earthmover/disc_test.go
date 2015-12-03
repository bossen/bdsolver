package earthmover

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"utils"
)

func TestSettingZeroDistanceAndExact(t *testing.T) {
	expected := [][]bool{
		[]bool{false, true, false, false, false, false, false},
		[]bool{false, true, true, false, false, false, false},
		[]bool{false, false, true, false, false, true, false},
		[]bool{false, false, false, false, false, false, false},
		[]bool{false, false, false, false, false, false, false},
		[]bool{false, false, false, false, false, false, false},
		[]bool{false, false, false, false, false, false, false}}
	
	c, m, visited, exact, d := setUpTest()
	
	w := FindFeasibleMatching(m, 0, 3, &c)
	setpair(m, w, exact, visited, d, &c)
	
	nonz := findNonZero(w, exact, d, &c)
	
	setZerosDistanceToZero(w, nonz, exact, d, &c)
	
	for i := 0; i < len(expected); i++ {
		for j := 0; j < len(expected[0]); j++ {
			assert.Equal(t, expected[i][j], exact[i][j], "the cell (%v,%v) were not correcly set", i, j)
		}
	}
	
	assert.True(t, utils.ApproxEqual(d[0][1], 1), "the distance for node (0,1) was changed")
	assert.True(t, utils.ApproxEqual(d[2][2], 0), "the distance for node (1,1) was not set to 0")
}

func TestGuassian(t *testing.T) {
	a := [][]float64{[]float64{1.0, -(1.0/6.0)}, []float64{0.0, 1.0}}
	b := []float64{5.0/6.0, 1.0/2.0}
	x, err := GaussPartial(a, b)
	
	assert.Equal(t, nil, err, "the linear equations was not calculate correctly")
	assert.True(t, utils.ApproxEqual(x[0], 11.0/12.0), "the found x value was not 11/12")
	assert.True(t, utils.ApproxEqual(x[1], 0.5), "the found x value was not 1/2")
}

func TestDisc(t *testing.T) {
	c, m, visited, exact, d := setUpTest()
	
	w := findFeasibleMatching(m, 0, 3, &c)
	setpair(m, w, exact, visited, d, &c)
	
	disc(1.0, w, exact, d, &c)
	
	assert.True(t, utils.ApproxEqual(d[0][3], 0.9133), "the distance was not correctly set")
	assert.True(t, utils.ApproxEqual(d[2][3], 0.49), "the distance was not correctly set")
}

func TestDiscResetsVisited(t *testing.T) {
	c, m, visited, exact, d := setUpTest()
	
	w := findFeasibleMatching(m, 0, 3, &c)
	setpair(m, w, exact, visited, d, &c)
	
	disc(1, w, exact, d, &c)
	
	for _, node := range c.Nodes {
		assert.False(t, node.Visited, "visited for node (%v,%v) was true", node.S, node.T)
	}
}
