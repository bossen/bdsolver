package coupling

import (
    "log"
)

func InitCoupling() int {
	return 1
}

type StatePair struct {
	S, T int
}

type Coupling struct {
	Matchings map[StatePair][]CouplingEdge
}

type CouplingEdge struct {
	S, T int
	Prob float64
    Color int
}

func New() Coupling {
    c := Coupling{}
    c.Matchings = make(map[StatePair][]CouplingEdge)
    return c
}


func Reachable(u, v int, c Coupling) []StatePair {
    // Using slices might be slow. If we got performance problems we might
    // implement using lists instead.
    var reachables []StatePair

    // Adding itself to reachables
    for _, ce := range c.Matchings[StatePair{u, v}] {
        if ce.S == u && ce.T == v {
            ce.Color = 1
        }

    }
    reachables = append(reachables, StatePair{u, v})

    // Find all reachables from the  u,v node
    reachables = visit(u, v, c, reachables)

    log.Println("reachables:")
    for _, t := range reachables {
        log.Println(t)

    }
  return reachables
}


func visit(u, v int, c Coupling, results []StatePair)  []StatePair {
    for _, ce := range c.Matchings[StatePair{u, v}] {
        if ce.Color == 0 {
            ce.Color = 1
            results = append(results, StatePair{ce.S, ce.T})
            results = visit(ce.S, ce.T, c, results)
        }
    }
    return results
}
