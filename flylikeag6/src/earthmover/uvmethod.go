package earthmover

import (
	"coupling"
	"log"
)

func calculateuv(tableu [][]*coupling.Edge, u []float64, v []float64, udefined []bool, vdefined []bool, d [][]float64) {
	collength := tableu[0]

	first := true
	for i := range tableu {
		for j := range collength {
			if !tableu[i][j].Basic {
				continue
			}

			node := tableu[i][j].To
			log.Printf("Calculate UV node.S = %v, node.T = %v and d[%v][%v] = %v", node.S, node.T, node.S, node.T, d[node.S][node.T])

			if first {
				u[i] = 0
				udefined[i] = true
				first = false
			}

			log.Printf("Calculate UV udefined[%v] = %v, vdefined[%v] = %v", i, udefined[i], j, vdefined[j])

			if udefined[i] {
				v[j] = d[node.S][node.T] - u[i]
				vdefined[j] = true
				log.Printf("Calculate UV d[%v][%v] - u[%v]: %v", node.S, node.T, i, d[node.S][node.T] - u[i])

			} else if vdefined[j] {
				u[i] = d[node.S][node.T] - v[j]
				udefined[i] = true
				log.Printf("Calculate UV d[%v][%v] - v[%v]: %v", node.S, node.T, j, d[node.S][node.T] - v[j])
			}
		}
	}
}

func findMinimum(tableu [][]*coupling.Edge, u []float64, v []float64, d [][]float64) (float64, int, int) {
	cols := tableu[0]
	
	node := tableu[0][0].To
	min := d[node.S][node.T]
	current := min
	iresult := 0
	jresult := 0

	for i := range tableu {
		for j := range cols {
			node = tableu[i][j].To
			current = d[node.S][node.T] - u[i] - v[j]
			if current < min {
				min = current
				iresult = i
				jresult = j
			}
		}
	}
	return min, iresult, jresult
}

func Uvmethod(node *coupling.Node, d [][]float64) (float64, int, int) {
	if node.Adj == nil {
		panic("Empty node!")
	}

	tableu := node.Adj
	collength := tableu[0]

	ulen := len(tableu)
	u := make([]float64, ulen, ulen)
	udefined := make([]bool, ulen, ulen)

	vlen := len(collength)
	v := make([]float64, vlen, vlen)
	vdefined := make([]bool, ulen, ulen)

	calculateuv(tableu, u, v, udefined, vdefined, d)

	min, iindex, jindex := findMinimum(tableu, u, v, d)
	return min, iindex, jindex
}
