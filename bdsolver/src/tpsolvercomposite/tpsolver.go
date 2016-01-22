package tpsolvercomposite

import (
    "markov"
    "coupling"
    "tpsolverdefault"
    "tpsolvercplex"
    "log"
)


var rundefault bool

func Solve(m markov.MarkovChain, n *coupling.Node, d [][]float64, min float64, i, j int) {
    if rundefault {
        log.Println("running default")
        tpsolverdefault.Solve(m, n, d, min, i, j)
        rundefault = false
    } else {
        log.Println("running tpsolvercplex")
        tpsolvercplex.Solve(m, n, d, min, i, j)
        rundefault = true
    }
}
