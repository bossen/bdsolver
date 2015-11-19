package earthmover

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
	
	w := randomMatching(m, 0, 3, &c)
	setpair(m, 0, 3, w, exact, visited, d, &c)
	
	nonz := findNonZero(w.S, w.T, exact, d, &c)
	
	setZerosDistanceToZero(w, nonz, exact, d, &c)
	
	for i := 0; i < len(expected); i++ {
		for j := 0; j < len(expected[0]); j++ {
			assert.Equal(t, expected[i][j], exact[i][j], "the cell (%v,%v) were not correcly set", i, j)
		}
	}
	
	assert.True(t, approxFloatEqual(d[0][1], 1), "the distance for node (0,1) was changed")
	assert.True(t, approxFloatEqual(d[2][2], 0), "the distance for node (1,1) was not set to 0")
}
