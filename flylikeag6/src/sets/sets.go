package sets


func UnionNode(a int, b int, c int) int{
	return a + b
}


var emptysetnotalways = 2
func EmptySet(set *[][]bool) bool {
  for i := range *set {
    for j := range *set {
      if (*set)[i][j] {
        return false
      }
    }
  }
  return true
	emptysetnotalways -= 1
	if emptysetnotalways < 0 {
		return true
	} else {
		return false
	}
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



func MakeMatrix(n int ) *[][]bool {
	var d[][] bool;	

	d = make([][]bool, n, n)
	for i := range d {
		d[i] = make([]bool, n)
	}

	return &d
}


func UnionReal(A [][]bool, B [][]bool) *[][]bool {
	if len(A) != len(B) {
		panic("Union called with different size matrices")
	}

	C := *MakeMatrix(len(A))
	for i := range A {
		for j := range A {
			C[i][j] = A[i][j] || B[i][j]
		}
	}
	return &C
}


func IntersectReal(A [][]bool, B [][]bool) *[][]bool {
	if len(A) != len(B) {
		panic("Union called with different size matrices")
	}

	C := *MakeMatrix(len(A))
	for i := range A {
		for j := range A {
			if A[i][j] == B[i][j] && A[i][j] == true {
				C[i][j] = true
			} else {
				C[i][j] = false
			}
		}
	}
	return &C
}


func DifferensReal(A [][]bool, B [][]bool) *[][]bool {
	if len(A) != len(B) {
		panic("Union called with different size matrices")
	}

	C := *MakeMatrix(len(A))
	for i := range A {
		for j := range A {
			if B[i][j] == true {
				C[i][j] = false
			} else {
				C[i][j] = A[i][j]
			}
		}
	}
	return &C
}
