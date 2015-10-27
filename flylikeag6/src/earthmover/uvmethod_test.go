package earthmover

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "coupling"
	"log"
)

func preparecoupling() coupling.Node {
  c := coupling.New()
  n1 := coupling.Node{S: 0, T: 0}
  n2 := coupling.Node{S: 0, T: 1}
  n3 := coupling.Node{S: 1, T: 0}
  n4 := coupling.Node{S: 1, T: 1}

  c.Nodes = []*coupling.Node{&n1, &n2, &n3, &n4}

  e1 := coupling.Edge{&n1, 0.1, true}
  e2 := coupling.Edge{&n2, 0.5, true}
  e3 := coupling.Edge{&n3, 0, false}
  e4 := coupling.Edge{&n4, 0.4, true}

  n2.Adj = [][]*coupling.Edge{
    []*coupling.Edge{&e1, &e2},
    []*coupling.Edge{&e3, &e4},
  }

  return n2
}

func prepared() [][]float64 {
  d := make([][]float64, 2, 2)
  d[0] = make([]float64, 2, 2)
  d[1] = make([]float64, 2, 2)
  d[0][0] = 5.0
  d[0][1] = 2.0
  d[1][0] = 0.0
  d[1][1] = 3.0

  return d
}

func preparecoupling2() coupling.Node {
  c := coupling.New()
  n1 := coupling.Node{S: 2, T: 1}
  n2 := coupling.Node{S: 2, T: 2}
  n3 := coupling.Node{S: 2, T: 3}
  n4 := coupling.Node{S: 3, T: 1}
  n5 := coupling.Node{S: 3, T: 2}
  n6 := coupling.Node{S: 3, T: 3}
  n7 := coupling.Node{S: 4, T: 1}
  n8 := coupling.Node{S: 4, T: 2}
  n9 := coupling.Node{S: 4, T: 3}
  n10 := coupling.Node{S: 6, T: 1}
  n11 := coupling.Node{S: 6, T: 2}
  n12 := coupling.Node{S: 6, T: 3}

  c.Nodes = []*coupling.Node{&n1, &n2, &n3, &n4, &n5, &n6, &n7, &n8, &n9, &n10, &n11, &n12}

  e1 := coupling.Edge{&n1, 0.333, true}
  e2 := coupling.Edge{&n2, 0.0, true}
  e3 := coupling.Edge{&n3, 0.0, false}
  e4 := coupling.Edge{&n4, 0.0, false}
  e5 := coupling.Edge{&n5, 0.333, true}
  e6 := coupling.Edge{&n6, 0.0, true}
  e7 := coupling.Edge{&n7, 0.0, false}
  e8 := coupling.Edge{&n8, 0.0, false}
  e9 := coupling.Edge{&n9, 0.166, true}
  e10 := coupling.Edge{&n10, 0.0, false}
  e11 := coupling.Edge{&n11, 0.0, false}
  e12 := coupling.Edge{&n12, 0.166, true}

  n7.Adj = [][]*coupling.Edge{
    []*coupling.Edge{&e1, &e2, &e3},
    []*coupling.Edge{&e4, &e5, &e6},
    []*coupling.Edge{&e7, &e8, &e9},
    []*coupling.Edge{&e10, &e11, &e12},
  }

  return n7
}

func prepared2() [][]float64 {
  size := 7
  d := make([][]float64, size, size)
  for i := range d {
    d[i] = make([]float64, size, size)
  }

  d[2][1] = 1.0
  d[2][2] = 0.0
  d[3][2] = 1.0
  d[3][3] = 0.0
  d[4][3] = 0.5
  d[6][3] = 1.0

  return d
}

func TestUvmethod(t *testing.T) {
  node := preparecoupling()
  d := prepared()

  res1, res2, res3 := Uvmethod(&node, d)
  assert.Equal(t, -6.0, res1, "incorrect return cell value")
  assert.Equal(t, 1, res2, "incorrect i index")
  assert.Equal(t, 0, res3, "incorrect j index")

	node2 := preparecoupling2()
	d2 := prepared2()
	res21, res22, res23 := Uvmethod(&node2, d2)
	log.Println("start 2")
  assert.Equal(t, -3.0, res21, "incorrect return cell value")
  assert.Equal(t, 3, res22, "incorrect i index")
  assert.Equal(t, 0, res23, "incorrect j index")
	log.Println(res22)
	log.Println(res23)
}


func TestCalculateuvSmall(t *testing.T) {
	uExpectedResults := []float64{0.0, 1.0}
	vExpectedResults := []float64{5.0, 2.0}
  node := preparecoupling()
  d := prepared()

  tableu := node.Adj
	collength := tableu[0]

	ulen := len(tableu)
	u := make([]float64, ulen, ulen)
	udefined := make([]bool, ulen, ulen)

	vlen := len(collength)
	v := make([]float64, vlen, vlen)
	vdefined := make([]bool, ulen, ulen)

  calculateuv(tableu, u, v, udefined, vdefined, d)
	for i := range u {
		assert.Equal(t, uExpectedResults[i], u[i], "Wrong value at u[%d].", i)
	}
	for i := range v {
		assert.Equal(t, vExpectedResults[i], v[i], "Wrong value at v[%d].", i)
	}
}

func TestCalculateuvBig(t *testing.T) {
	uExpectedResults := []float64{0.0, 1.0, 1.5, 2.0}
	vExpectedResults := []float64{1.0, 0.0, -1.0}
  node := preparecoupling2()
  d := prepared2()

  tableu := node.Adj
	collength := tableu[0]

	ulen := len(tableu)
	u := make([]float64, ulen, ulen)
	udefined := make([]bool, ulen, ulen)

	vlen := len(collength)
	v := make([]float64, vlen, vlen)
	vdefined := make([]bool, ulen, ulen)

  calculateuv(tableu, u, v, udefined, vdefined, d)
	for i := range u {
		assert.Equal(t, uExpectedResults[i], u[i], "Wrong value at u[%d].", i)
	}
	for i := range v {
		assert.Equal(t, vExpectedResults[i], v[i], "Wrong value at v[%d].", i)
	}
}

func TestFindMinimumSmall(t *testing.T) {
	node := preparecoupling()
	d := prepared()

  tableu := node.Adj
	u := []float64{0.0, 1.0}
	v := []float64{5.0, 2.0}
	min, iindex, jindex := findMinimum(tableu, u, v, d)
	assert.Equal(t, min, -6.0, "Wrong findMinimum result")
	assert.Equal(t, iindex, 1, "Wrong findMinimum i index")
	assert.Equal(t, jindex, 0, "Wrong findMinimum j index")
}

func TestFindMinimumBig(t *testing.T) {
	node := preparecoupling2()
	d := prepared2()

  tableu := node.Adj
	u := []float64{0.0, 1.0, 1.5, 2.0}
	v := []float64{1.0, 0.0, -1.0}
	min, iindex, jindex := findMinimum(tableu, u, v, d)
	assert.Equal(t, min, -3.0, "Wrong findMinimum result")
	assert.Equal(t, iindex, 3, "Wrong findMinimum i index")
	assert.Equal(t, jindex, 0, "Wrong findMinimum j index")
}
