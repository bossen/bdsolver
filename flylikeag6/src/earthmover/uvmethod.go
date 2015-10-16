package earthmover

import (
	"coupling"
)

func findMinimum(tableu *[][]coupling.Edge, u []float64, v []float64) (float64, int, int) {
	cols := (*tableu)[0]

	min := (*tableu)[0][0].Prob
	current := min
	iresult := 0
	jresult := 0

	for i := range *tableu {
		for j := range cols {
			current = (*tableu)[i][j].Prob - u[i] - v[j]
			if current < min {
				min = current
				iresult = i
				jresult = j
			}
		}
	}
	return min, iresult, jresult
}

func calculateuv(tableu *[][]coupling.Edge, u *[]float64, v *[]float64, udefined *[]bool, vdefined *[]bool) {
	cols := (*tableu)[0]

	first := true
	for i := range *tableu {
		for j := range cols {
			if !(*tableu)[i][j].IsBasic {
				continue
			}

			if first {
				(*u)[i] = 0
				(*udefined)[i] = true
				first = false
			}

			if (*udefined)[i] {
				(*v)[j] = (*tableu)[i][j].Prob - (*u)[i]
				(*vdefined)[j] = true

			} else if (*vdefined)[j] {
				(*u)[i] = (*tableu)[i][j].Prob - (*v)[j]
				(*vdefined)[j] = true
			}
		}
	}
}

func Uvmethod(node *coupling.Node) (float64, int, int) {
	if node.Adj == nil {
		panic("Empty node!")
	}

	tableu := node.Adj
	cols := (*tableu)[0]

	ulen := len(*tableu)
	u := make([]float64, ulen, ulen)
	udefined := make([]bool, ulen, ulen)

	vlen := len(cols)
	v := make([]float64, vlen, vlen)
	vdefined := make([]bool, ulen, ulen)

	calculateuv(tableu, &u, &v, &udefined, &vdefined)

	min, iindex, jindex := findMinimum(tableu, u, v)
	return min, iindex, jindex
}
