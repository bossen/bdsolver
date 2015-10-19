package earthmover

import (
	"coupling"
  "log"
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

func calculateuv(tableu *[][]coupling.Edge, u *[]float64, v *[]float64, udefined *[]bool, vdefined *[]bool, d *[256][256]float64) {
	collength := (*tableu)[0]

	first := true
	for i := range *tableu {
		for j := range collength {
			if !(*tableu)[i][j].IsBasic {
				continue
			}

      node := (*tableu)[i][j].To
      log.Println(node)

			if first {
				(*u)[i] = 0
				(*udefined)[i] = true
				first = false
			}

			if (*udefined)[i] {
				(*v)[j] = (*d)[node.S][node.T] - (*u)[i]
				(*vdefined)[j] = true

			} else if (*vdefined)[j] {
				(*u)[i] = (*d)[node.S][node.T] - (*v)[j]
				(*vdefined)[j] = true
			}
		}
	}
}

func Uvmethod(node *coupling.Node, d *[256][256]float64) (float64, int, int) {
	if node.Adj == nil {
		panic("Empty node!")
	}

	tableu := node.Adj
	collength := (*tableu)[0]

	ulen := len(*tableu)
	u := make([]float64, ulen, ulen)
	udefined := make([]bool, ulen, ulen)

	vlen := len(collength)
	v := make([]float64, vlen, vlen)
	vdefined := make([]bool, ulen, ulen)

	calculateuv(tableu, &u, &v, &udefined, &vdefined, d)

	min, iindex, jindex := findMinimum(tableu, u, v)
	return min, iindex, jindex
}
