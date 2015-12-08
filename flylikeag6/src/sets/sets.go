package sets

func UnionNode(a int, b int, c int) int {
	return a + b
}

func EmptySet(set [][]bool) bool {
	for i := range set {
		for j := range set {
			if set[i][j] {
				return false
			}
		}
	}
	return true
}

func Intersect(a int, b int) int {
	return 1
}

func Union(a int, b int) int {
	return 1
}

func Differens(a int, b int) int {
	return a
}

func MakeMatrix(n int) [][]bool {
	var d [][]bool

	d = make([][]bool, n, n)
	for i := range d {
		d[i] = make([]bool, n)
	}

	return d
}

func UnionReal(A *[][]bool, B *[][]bool) *[][]bool {
	if len(*A) != len(*B) {
		panic("Union called with different size matrices")
	}

	C := MakeMatrix(len(*A))
	for i := range *A {
		for j := range *A {
			C[i][j] = (*A)[i][j] || (*B)[i][j]
		}
	}
	return &C
}

func IntersectReal(A *[][]bool, B *[][]bool) *[][]bool {
	if len(*A) != len(*B) {
		panic("Union called with different size matrices")
	}

	C := MakeMatrix(len(*A))
	for i := range *A {
		for j := range *A {
			if (*A)[i][j] == (*B)[i][j] && (*A)[i][j] == true {
				C[i][j] = true
			} else {
				C[i][j] = false
			}
		}
	}
	return &C
}

func DifferensReal(A *[][]bool, B *[][]bool) *[][]bool {
	if len(*A) != len(*B) {
		panic("Union called with different size matrices")
	}

	C := MakeMatrix(len(*A))
	for i := range *A {
		for j := range *A {
			if (*B)[i][j] == true {
				C[i][j] = false
			} else {
				C[i][j] = (*A)[i][j]
			}
		}
	}
	return &C
}

func InitToCompute(n int) [][]bool {
	toCompute := MakeMatrix(n)
	for i := range toCompute {
		for j := range toCompute {
			if i < j {
				// when j greater or equal to i, we use the default false value
				toCompute[i][j] = true
			}
		}
	}
	return toCompute
}

func InitD(n int) [][]float64{
	d := make([][]float64, n, n)
	for i := 0; i < n; i++ {
		d[i] = make([]float64, n, n)
		for j := range d[i] {
			if i != j {
				// when i and j is the same, we use the default value 0
				d[i][j] = 1
			}
		}
	}
	return d
}
