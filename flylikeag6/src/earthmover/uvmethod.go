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
			
			log.Printf("node.S = %v, node.T = %v and d[%v][%v] = %v", node.S, node.T, node.S, node.T, d[node.S][node.T])
			
			if first {
				u[i] = 0
				udefined[i] = true
				first = false
			}
			
			log.Printf("udefined[%v] = %v, vdefined[%v] = %v", i, udefined[i], j, vdefined[j])
			if udefined[i] && !vdefined[j] {
				v[j] = d[node.S][node.T] - u[i]
				vdefined[j] = true
				log.Printf("d[%v][%v] - u[%v]: %v", node.S, node.T, i, d[node.S][node.T] - u[i])
			} else if vdefined[j] && !udefined[i] {
				u[i] = d[node.S][node.T] - v[j]
				udefined[i] = true
				log.Printf("UV d[%v][%v] - v[%v]: %v", node.S, node.T, j, d[node.S][node.T] - v[j])
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
			log.Printf("index %v %v with value %v", i, j, d[node.S][node.T] - u[i] - v[j])
			
			if tableu[i][j].Basic {
				continue
			}
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

	calculateuv2(tableu, u, v, udefined, vdefined, d)

	min, iindex, jindex := findMinimum(tableu, u, v, d)
	log.Printf("with index %v %v node (%v,%v) was chosen by UVMethod with value %v", iindex, jindex, node.Adj[iindex][jindex].To.S, node.Adj[iindex][jindex].To.T, min)
	return min, iindex, jindex
}

func calculateuv2(tableu [][]*coupling.Edge, u []float64, v []float64, udefined []bool, vdefined []bool, d [][]float64) {
	var unfinished []IntPair
	
	for i, row := range tableu {
		for j, edge := range row {
			if !edge.Basic {
				continue
			}
			
			node := edge.To
			
			if !udefined[i] && !vdefined[j] {
				unfinished = append(unfinished, IntPair{i, j})
				continue
			}
			
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
	
	unfinishcounter := len(unfinished)
	var finished bool
	
	for !finished {
		startunfinished := unfinishcounter
		finished = true
		
		for _, cell := range unfinished {
			i, j := cell.I, cell.J
			node := tableu[i][j].To
			
			if !udefined[i] && !vdefined[j] {
				finished = false
				continue
			}

			if udefined[i] && !vdefined[j] {
				v[j] = d[node.S][node.T] - u[i]
				vdefined[j] = true
				log.Printf("v[%v] = %v - u[%v] = %v", j, d[node.S][node.T], i, d[node.S][node.T] - u[i])
				unfinishcounter = unfinishcounter - 1
			} else if vdefined[j] && !udefined[i] {
				u[i] = d[node.S][node.T] - v[j]
				udefined[i] = true
				log.Printf("u[%v] = %v - v[%v] = %v", i, d[node.S][node.T], j, d[node.S][node.T] - v[j])
				unfinishcounter = unfinishcounter - 1
			}
		}
		
		if unfinishcounter > 0 && startunfinished == unfinishcounter {
			log.Println("have to recover")
			var i, j int
			
			for n, cell := range unfinished {
				if udefined[cell.I] || vdefined[cell.J] {
					continue
				}
				i, j = unfinished[n].I, unfinished[n].J
			}
			
			rowlen := len(tableu[0])
			for n := range tableu[0] {
				if !tableu[i][(j+n) % rowlen].Basic {
					log.Printf("setting (%v,%v) to Basic", i, (j+n) % rowlen)
					tableu[i][(j+n) % rowlen].Basic = true
					unfinished = append(unfinished, IntPair{i, (j+n) % rowlen})
					break
				}
			}
		}
	}
	
	return
}
