package earthmover


import (
	"fmt"
    "sets"
)

func initexact() int {
	return 1
}

func initcoupling() int {
	return 1
}

//Computes all the possible combinations of the different nodes. This could be optimized, by setting everything below the i == j diagonal to false.
func initToCompute(n int) *[][]bool {
  toCompute := *sets.MakeMatrix(n)
  for i := range toCompute {
    for j := range toCompute {
      if i == j {
        toCompute[i][j] = false
      } else {
        toCompute[i][j] = true
      }
    }
  }
  return &toCompute
}

func extractrandomfromset(tocompute *[][]bool) (int, int) {
  for i := range *tocompute {
    for j:= range *tocompute {
      if (*tocompute)[i][j] == true {
        return i, j
      }
    }
  }
  panic("Tried to extract random element from empty set!")
}

func label(node int) string {
	return "red"
}


func removeedgesfromnodes(coupling *int, exact *int) int {
	return 1
}


func getoptimalschedule(d [256][256] int, u int, v int) int {
	return 1
}

func isOptimal() bool {
  return true
}



func BipseudoMetric(n int) {
	var d [256][256]int
	visited := *sets.MakeMatrix(n)
	exact := initexact()
	coupling := initcoupling()
	tocompute := initToCompute(n)
	lambda := 1
	
	for !sets.EmptySet(tocompute) {
		s, t := extractrandomfromset(tocompute)
    fmt.Println(s)
    fmt.Println(t)
		if label(s) != label(t) {
			d[s][t] = 1
			exact = sets.UnionNode(exact, s, t)
			visited[s][t] = true
		} else if s == t {
			d[s][t] = 0
			exact = sets.UnionNode(exact, s, t)
			visited[s][t] = true

		} else {
			// if s,t not in visited ...

            disc(lambda, s, t, &exact, &coupling)

			for !isOptimal()  { // TODO: add u, v in for loop. While loop check for NOT optimal matching
				// w := getoptimalschedule(d, u, v)
				w := 1
				setpair(1, s, t, w, &exact, &visited, &coupling)
				disc(lambda, s, t, &exact, &coupling)
			}
			exact = sets.Union(exact, reachable(s, t, coupling))
			removeedgesfromnodes(&coupling, &exact)
		}
    
		tocompute = sets.IntersectReal(*tocompute, *tocompute) //TODO THIS IS WRONG, use exact as second parameter, instead of tocompute twice

    break; //TODO remove this. This is for ending the code
	}
	setpair(1, 1, 1, 1, &exact, &visited, &coupling)
	disc(1, 1, 1, &exact, &coupling)
	fmt.Println("hello world!")
	// return what?
}
