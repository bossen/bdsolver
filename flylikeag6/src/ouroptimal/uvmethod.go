package ouroptimal

import (
	"coupling"
	"log"
)

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
			log.Printf("d[%v][%v] - u[%v] - v[%v] = %v - %v - %v = %v", i, j, i, j, d[node.S][node.T], u[i], v[j], d[node.S][node.T] - u[i] - v[j])
			
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
	log.Printf("calculate UV for node (%v,%v) with matching:", node.S, node.T)
	for i, row := range node.Adj {
		for j, edge := range row {
			log.Printf(" - index %v %v: Prob: %v, Basic: %v, Dist: %v", i, j, edge.Prob, edge.Basic, d[edge.To.S][edge.To.T])
		}
	}

	tableu := node.Adj

	ulen := len(tableu)
	u := make([]float64, ulen, ulen)
	udefined := make([]bool, ulen, ulen)

	vlen := len(tableu[0])
	v := make([]float64, vlen, vlen)
	vdefined := make([]bool, vlen, vlen)
	
	u[0] = 0
	udefined[0] = true

	calculateuv(tableu, u, v, udefined, vdefined, d)

	min, iindex, jindex := findMinimum(tableu, u, v, d)
	log.Printf("with index %v %v node (%v,%v) was chosen by UVMethod with value %v", iindex, jindex, node.Adj[iindex][jindex].To.S, node.Adj[iindex][jindex].To.T, min)
	return min, iindex, jindex
}

func calculateuv(tableu [][]*coupling.Edge, u []float64, v []float64, udefined []bool, vdefined []bool, d [][]float64) {
	var unfinished []IntPair
	finished := true
	
	// here we calculate the u-v modifiers possible in the first iteration
	for i, row := range tableu {
		for j, edge := range row {
			if !edge.Basic {
				continue
			}
			
			node := edge.To
			
			// if neither the u or v modifier has already been definied for the row/column pair, we add the index pair (i,j) so we can complete it later
			if !udefined[i] && !vdefined[j] {
				unfinished = append(unfinished, IntPair{i, j})
				finished = false
				continue
			}
			
			// if the u modifier can been defined, we can calculate the v modifier, and vice versa
			if udefined[i] {
				v[j] = d[node.S][node.T] - u[i]
				vdefined[j] = true
				log.Printf("v[%v] = %v - u[%v] = %v", j, d[node.S][node.T], i, d[node.S][node.T] - u[i])
			} else if vdefined[j] {
				u[i] = d[node.S][node.T] - v[j]
				udefined[i] = true
				log.Printf("u[%v] = %v - v[%v] = %v", i, d[node.S][node.T], j, d[node.S][node.T] - v[j])
			}
		}
	}
	
	// here we calculate the remaining u-v modifers, until none are left
	for !finished {
		finished = true
		
		for _, cell := range unfinished {
			i, j := cell.I, cell.J
			node := tableu[i][j].To
			
			// both modifiers are still undefinied, so we have to iterate again
			if !udefined[i] && !vdefined[j] {
				finished = false
				continue
			}
			
			// if the v modifier has not been defined, we can calculate it using the u modifier, and vice versa
			if !vdefined[j] {
				v[j] = d[node.S][node.T] - u[i]
				vdefined[j] = true
				log.Printf("v[%v] = %v - u[%v] = %v", j, d[node.S][node.T], i, d[node.S][node.T] - u[i])
			} else if !udefined[i] {
				u[i] = d[node.S][node.T] - v[j]
				udefined[i] = true
				log.Printf("u[%v] = %v - v[%v] = %v", i, d[node.S][node.T], j, d[node.S][node.T] - v[j])
			}
		}
	}
	
	return
}
