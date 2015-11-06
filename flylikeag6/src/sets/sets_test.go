package sets

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func makeSlice(n int) [][]bool {
	var set [][]bool
	set = make([][]bool, n, n)
	for i := range set {
		set[i] = make([]bool, n)
	}
	return set
}

func TestMakeMatrix(t *testing.T) {
	set := *MakeMatrix(0)
	assert.True(t, EmptySet(&set), "MakeMatrix creates a wrong matrix! Slice size 0")
	assert.Equal(t, 0, len(set), "MakeMatrix creates a wrong sized matrix!")

	set = *MakeMatrix(100)
	assert.True(t, EmptySet(&set), "MakeMatrix creates a wrong matrix! Slice size 100")
	assert.Equal(t, 100, len(set), "MakeMatrix creates a wrong sized matrix!")
	for _, s := range set {
		assert.Equal(t, 100, len(s), "MakeMatrix creates a wrong sized matrix!")
	}
}

func TestEmptySets(t *testing.T) {
	set := makeSlice(0)
	assert.True(t, EmptySet(&set), "EmptySet returns wrong value! Slice size 0")
	set = makeSlice(100)
	assert.True(t, EmptySet(&set), "EmptySet returns wrong value! Slice size 100")
	set[0][0] = true
	assert.False(t, EmptySet(&set), "EmptySet returns wrong value! set[0][0] = true")
	set[0][0] = false
	set[99][99] = true
	assert.False(t, EmptySet(&set), "EmptySet returns wrong value! set[99][99] = true")
}

func TestUnionReal(t *testing.T) {
	A := makeSlice(100)
	B := makeSlice(100)
	C := UnionReal(A, B)
	res := EmptySet(C)
	assert.True(t, res, "UnionReal returns wrong value! Union two empty sets")

	A[5][5] = true
	C = UnionReal(A, B)
	assert.True(t, (*C)[5][5], "UnionReal returns wrong value! Union A[5][5] = true")

	A[5][5] = false
	B[5][5] = true
	C = UnionReal(A, B)
	assert.True(t, (*C)[5][5], "UnionReal returns wrong value! Union B[5][5] = true")
}

func TestIntersectReal(t *testing.T) {
	A := makeSlice(100)
	B := makeSlice(100)
	C := UnionReal(A, B)
	res := EmptySet(C)
	assert.True(t, res, "IntersectReal returns wrong value! Intersect two empty sets")

	A[5][5] = true
	C = IntersectReal(A, B)
	assert.False(t, (*C)[5][5], "IntersectReal returns wrong value! Intersect A[5][5] = true")

	A[5][5] = false
	B[5][5] = true
	C = IntersectReal(A, B)
	assert.False(t, (*C)[5][5], "IntersectReal returns wrong value! Intersect B[5][5] = true")

	A[5][5] = true
	B[5][5] = true
	C = IntersectReal(A, B)
	assert.True(t, (*C)[5][5], "IntersectReal returns wrong value! Intersect A[5][5] = true B[5][5] = true")
}

func TestDifferensReal(t *testing.T) {
	A := makeSlice(100)
	B := makeSlice(100)
	C := UnionReal(A, B)
	res := EmptySet(C)
	assert.True(t, res, "DifferensReal returns wrong value! Differens two empty sets")

	A[5][5] = true
	C = DifferensReal(A, B)
	assert.True(t, (*C)[5][5], "DifferensReal returns wrong value! Differens A[5][5] = true two empty sets")

	A[5][5] = false
	B[5][5] = true
	C = DifferensReal(A, B)
	assert.False(t, (*C)[5][5], "DifferensReal returns wrong value! Differens B[5][5] = true")
}
