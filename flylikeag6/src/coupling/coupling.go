package coupling

func InitCoupling() int {
	return 1
}

type StatePair struct {
	S1, S2 int
}

type Coupling struct {
	Matchings map[StatePair][]CouplingEdge
}

type CouplingEdge struct {
	SP StatePair
	Prob float64
}
