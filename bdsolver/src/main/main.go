package main

import (
	"pseudometric"
	"tpsolverdefault"
	"tpsolvercplex"
    "tpsolvercomposite"
	"log"
	"fmt"
	"compiler"
	"markov"
	"os"
	"strconv"
	"io/ioutil"
	"strings"
)

var version = "No version provided"

func printdocumentation() {
    documentation := `bdsolver - Bisimilarity distance solver version '%version%'

usage: bdsolver [arguments] file        solve the specified file

Arguments:
   -l <lambda>           Defines the lambda. The lambda has to be larger than 0 up to 1.
   -tpsolver <solver>    Defines the transportation solver. Possible arguments are cplex or default.
   -v                    Running verbose logging.
   -m                    Print result as a bisimilarity distance matrix.
   -h                    Shows this description.
`
    documentation =  strings.Replace(documentation, "%version%", version, -1)

    fmt.Println(documentation)

}

func main() {
    log.SetFlags(log.Lshortfile)
	lambda := 1.0
	TPSolver := tpsolverdefault.Solve
	log.SetOutput(ioutil.Discard)
    resultPrinter := printStandard
	filename := "NOFILENAMECHOSEN"
	
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "-l" {
			arg := getOrFail(i+1, "expected a float after -l but there was nothing")
			
			temp, err := strconv.ParseFloat(arg, 64)
			lambda = temp
			
			if err != nil {
				fail(fmt.Sprintf("expected a float for lambda(-l) but got %s", arg))
			}
			if lambda < 0 || lambda > 1 {
				fail(fmt.Sprintf("invalid lambda value, has to be larger than 0 and less or equal to 1"))
			}
			
			i++
		} else if os.Args[i] == "-tpsolver" {
			arg := getOrFail(i+1, "expected cplex or default after -tpsolver but there was nothing")
			
			if arg == "cplex" {
				TPSolver = tpsolvercplex.Solve
			} else if arg == "composite" {
                TPSolver = tpsolvercomposite.Solve
			} else  if arg != "default" {
				fail(fmt.Sprintf("expected cplex or default after -tpsolver but got %s", arg))
			}
			
			i++
		} else if os.Args[i] == "-v" {
			log.SetOutput(os.Stdout)
        } else if os.Args[i] == "-h" {
            printdocumentation()
            os.Exit(0)
        } else if os.Args[i] == "-m" {
            resultPrinter = printDistanceMatrix
		} else {
		    filename = getOrFail(i, "expected a file but there was nothing")
		}
	}


    if filename == "NOFILENAMECHOSEN" {
        printdocumentation()
        os.Exit(0)
    }
    
    mc, err := compiler.Parse(filename)
			
	if err != nil {
		fail(err.Error())
	}
	
	d := pseudometric.PseudoMetric(mc, lambda, TPSolver)
    resultPrinter(d, mc)
}

func printDistanceMatrix(d [][]float64, m markov.MarkovChain) {
    // As we optimized the algorithm, we only got the top left the matrix
    // The other part is symmetric, and this is what we recreate here.
    for i := range d {
        for j := i+1; j < len(d); j++ {
            d[j][i] = d[i][j]
        }
    }
    fmt.Printf("The distance matrix for markov chain %s:", os.Args[len(os.Args)-1])
    fmt.Println()
    for i := range d {
        fmt.Printf("%v",  d[i])
        fmt.Println()
    }
}

func printStandard(d [][]float64, m markov.MarkovChain) {
    for i := range d {
        for j := i+1; j < len(d); j++ {
            if m.Labels[i] == m.Labels[j] {
                fmt.Printf("%v <-> %v %v", i+1, j+1, d[i][j])
                fmt.Println()
            }
        }
    }
}

func getOrFail(i int, message string) string {
	if i > len(os.Args)-1 {
		fail(message)
	}
	return os.Args[i]
}

func fail(message string) {
	fmt.Println(message)
    printdocumentation()
	os.Exit(1)
}
