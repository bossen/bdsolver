package earthmover

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"utils"
)

func TestSettingZeroDistanceAndExact(t *testing.T) {
	expected := [][]bool{
		[]bool{false, true, false, false, false, false, false},
		[]bool{true, true, true, false, false, false, false},
		[]bool{false, true, true, false, false, true, false},
		[]bool{false, false, false, false, false, false, false},
		[]bool{false, false, false, false, false, false, false},
		[]bool{false, false, true, false, false, false, false},
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
