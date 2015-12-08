package ouroptimal

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"coupling"
)

func preparenode() coupling.Node {
	var n coupling.Node
	
	e1 := coupling.Edge{&coupling.Node{S: 0, T: 0}, 0.1, true}
	e2 := coupling.Edge{&coupling.Node{S: 0, T: 1}, 0.5, true}
	e3 := coupling.Edge{&coupling.Node{S: 1, T: 0}, 0, false}
	e4 := coupling.Edge{&coupling.Node{S: 1, T: 1}, 0.4, true}
	
	n.Adj = [][]*coupling.Edge{
    []*coupling.Edge{&e1, &e2},
    []*coupling.Edge{&e3, &e4}}
	
	return n
}

func preparenode2() coupling.Node {
	var n coupling.Node
	
	e1 := coupling.Edge{&coupling.Node{S: 0, T: 0}, 0, false}
	e2 := coupling.Edge{&coupling.Node{S: 0, T: 1}, 0.1, true}
	e3 := coupling.Edge{&coupling.Node{S: 1, T: 0}, 0.4, true}
	e4 := coupling.Edge{&coupling.Node{S: 1, T: 1}, 0.5, true}
	
	n.Adj = [][]*coupling.Edge{
    []*coupling.Edge{&e1, &e2},
    []*coupling.Edge{&e3, &e4}}
	
	return n
}

func prepared() [][]float64 {
	d := make([][]float64, 2, 2)
	d[0] = make([]float64, 2, 2)
	d[1] = make([]float64, 2, 2)
	d[0][0] = 0.5
	d[0][1] = 0.2
	d[1][0] = 0.0
	d[1][1] = 0.9
	
	return d
}

func prepared2() [][]float64 {
	d := make([][]float64, 2, 2)
	d[0] = make([]float64, 2, 2)
	d[1] = make([]float64, 2, 2)
	d[0][0] = 0.0
	d[0][1] = 0.5
	d[1][0] = 0.9
	d[1][1] = 0.2
	
	return d
}

func TestFindMinimum(t *testing.T) {
	node := preparenode()
	d := prepared()

	tableu := node.Adj
	u := []float64{0.0, 0.7}
	v := []float64{0.5, 0.2}
	min, iindex, jindex := findMinimum(tableu, u, v, d)
	assert.Equal(t, -1.2, min, "Wrong findMinimum result")
	assert.Equal(t, 1, iindex, "Wrong findMinimum i index")
	assert.Equal(t, 0, jindex, "Wrong findMinimum j index")
}

func TestCalculateUV(t *testing.T) {
	uexpected := []float64{0, 0.7}
	vexpected := []float64{0.5, 0.2}
	
	node := preparenode()
	d := prepared()
	
	u := make([]float64, 2, 2)
	v := make([]float64, 2, 2)
	udefined := make([]bool, 2, 2)
	vdefined := make([]bool, 2, 2)
	udefined[0] = true
	
	calculateuv(node.Adj, u, v, udefined, vdefined, d)
	
	assert.Equal(t, uexpected[0], u[0])
	assert.Equal(t, uexpected[1], u[1])
	assert.Equal(t, vexpected[0], v[0])
	assert.Equal(t, vexpected[1], v[1])
}

func TestCalculateUVLoop(t *testing.T) {
	uexpected := []float64{0, -0.3}
	vexpected := []float64{1.2, 0.5}
	
	node := preparenode2()
	d := prepared2()
	
	u := make([]float64, 2, 2)
	v := make([]float64, 2, 2)
	udefined := make([]bool, 2, 2)
	vdefined := make([]bool, 2, 2)
	udefined[0] = true
	
	calculateuv(node.Adj, u, v, udefined, vdefined, d)
	
	assert.Equal(t, uexpected[0], u[0])
	assert.Equal(t, uexpected[1], u[1])
	assert.Equal(t, vexpected[0], v[0])
	assert.Equal(t, vexpected[1], v[1])
}
