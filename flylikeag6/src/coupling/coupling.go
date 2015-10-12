package coupling

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
	IsBasic bool
}

func New() Coupling {
    c := Coupling{}
    c.Matchings = make(map[StatePair][]CouplingEdge)
    return c
}
