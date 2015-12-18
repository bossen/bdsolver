package markov

func New() MarkovChain {
	n := 2

	mc := MarkovChain{}
	mc.Labels = make([]int, n)
	mc.Transitions = make([][]float64, n, n)

	// Init transitions
	for i := range mc.Transitions {
		mc.Transitions[i] = make([]float64, n)
	}
	mc.Transitions[0][0] = 0.7
	mc.Transitions[1][0] = 0.5
	mc.Transitions[0][1] = 0.3
	mc.Transitions[1][1] = 0.5

	mc.Labels[0] = 1
	mc.Labels[1] = 1

	return mc
}

/* type Node struct {
	Id, Label in
 }
*/

type MarkovChain struct {
	Labels      []int
	Transitions [][]float64
}
