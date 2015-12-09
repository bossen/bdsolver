package utils



// credits to https://gist.github.com/cevaris/bc331cbe970b03816c6b
func ApproxEqual(a, b float64) bool {
	var epsilon float64 = 0.00000000000001

	if (a-b) < epsilon && (b-a) < epsilon {
		return true
	}
	return false
}

