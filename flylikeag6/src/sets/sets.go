package sets


func UnionNode(a int, b int, c int) int{
	return a + b
}


var emptysetnotalways = 2
func EmptySet(set int) bool {
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

