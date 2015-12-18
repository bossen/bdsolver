package tpsolvercplex

import (
		"sync"
		"strings"
		"bytes"
		"log"
		"testing"

		"coupling"
		"markov"
		"matching"
		"utils"

		"github.com/stretchr/testify/assert"
)

func setUpCouplingMatching() coupling.Coupling {
	return coupling.New()
}

func setUpMarkov() markov.MarkovChain {
	return markov.MarkovChain{
		Labels: []int{0, 1, 0, 0, 0, 1, 0},
		Transitions: [][]float64{
			[]float64{0.0, 0.33333333337, 0.33333333337, 0.1666666666666667, 0.0, 0.1666666666666667, 0.0},
			[]float64{0.0, 0.0, 0.4, 0.4, 0.0, 0.2, 0.0},
			[]float64{0.0, 0.5, 0.5, 0.0, 0.0, 0.0, 0.0},
			[]float64{0.33333333337, 0.33333333337, 0.33333333337, 0.0, 0.0, 0.0, 0.0},
			[]float64{0.4, 0.4, 0.2, 0.0, 0.0, 0.0, 0.0},
			[]float64{0.0, 0.1, 0.0, 0.2, 0.5, 0.2, 0.0},
			[]float64{0.0, 0.2, 0.33, 0.0, 0.1, 0.2, 0.17}}}
}

func initializeD(n int) [][]float64{
	d := make([][]float64, n, n)
	for i := 0; i < n; i++ {
		d[i] = make([]float64, n, n)
		for j := range d[i] {
			if i != j {		//when i equal j, it will use the default value of 0
				d[i][j] = 1
			}
		}
	}
	return d
}

func setUpTest() (coupling.Coupling, markov.MarkovChain, [][]float64) {
	c := setUpCouplingMatching()
	m := setUpMarkov()
	n := len(m.Transitions)
	d := initializeD(n)

	return c, m, d
}

	var expectedProbValues = []float64{0.0, 0.333333, 0.0, 0.0, 0.0, 0.333333, 0.166667, 0.0, 0.0, 0.166667, 0.0, 0.0}
	var expectedBasicValues = []bool{false, true, true, true, false, true, true, false, false, true, false, false}

func TestCplexOptimize(t *testing.T) {
	c, m, d := setUpTest()
	node := matching.FindFeasibleMatching(m, 0, 3, &c)
    Solve(m, node, d, 0.0, 0, 0)
	k := 0
	for i := range node.Adj {
		for _, edge := range node.Adj[i] {
			assert.True(t, utils.ApproxEqual(edge.Prob, expectedProbValues[k]), "optimize does not work for %v, probability, iteration %v", edge, k)
			assert.Equal(t, expectedBasicValues[k], edge.Basic, "optimize does not work for %v, basic value, iteration %v", edge, k)
			k++
		}
	}
}

func TestExeCmd(t *testing.T) {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	assert.Equal(t, exeCmd("echo -n test", wg), "test", "exeCmd does not work!")
}

func TestFindConstraints(t *testing.T) {
	var expectedUsedValues = []int{1, 3}
	var buffer bytes.Buffer
	var transitions = []float64{0.0, 0.33333333337, 0.0, 0.1666666666666667}

	size, used := findConstraints(&buffer, transitions)

	assert.Equal(t, "0.33333333337 0.1666666666666667 ", buffer.String(), "buffer has the wrong value!")
	assert.Equal(t, 2, size, "size has the wrong value!")
	for i, u := range used {
		assert.Equal(t, expectedUsedValues[i], u, "used[%v] has the wrong value!", i)
	}
}

func TestFindConstraintsEmpty(t *testing.T) {
	var buffer bytes.Buffer

	size, used := findConstraints(&buffer, []float64{})

	assert.Equal(t, "", buffer.String(), "buffer has the wrong value!")
	assert.Equal(t, 0, size, "size has the wrong value!")
	assert.Equal(t, 0, len(used), "used is not empty!")
}

func TestOptimize(t *testing.T) {
	var expectedProbValues = []float64{0.0, 0.333333, 0.0, 0.0, 0.0, 0.333333, 0.166667, 0.0, 0.0, 0.166667, 0.0, 0.0}

	newValues := optimize("1 0 1 1 1 0 1 1 1 1 1 1", " 0.3333333333333333 0.3333333333333333 0.1666666666666667 0.1666666666666667 0.3333333333333333 0.3333333333333333 0.3333333333333333", 4, 3)
	for i, value := range newValues {
		assert.Equal(t, expectedProbValues[i], value, "Item number %v had a wrong value", i)
	}
}

func TestRetrieveDValues(t *testing.T) {
	var expectedValues = []float64{1, 1, 0.3, 1, 0, 1}
	var buffer bytes.Buffer
	rowused := []int{1, 2}
	columnused := []int{0, 2, 4}
	d := initializeD(5)
	d[1][4] = 0.3

	retrieveDValues(&buffer, d, rowused, columnused)
	log.Printf("%v", buffer.String())
	k := 0
	for _, i := range rowused {
		for _, j := range columnused {
			u, v := utils.GetMinMax(i, j)
			assert.Equal(t, expectedValues[k], d[u][v], "d[%v][%v] = %v", u, v, d[u][v])
			k++
		}
	}
}

func TestStringArrayToFloat(t *testing.T) {
	var expectedValues = []float64{4.0, 3.0, 4.0, 1.231, 0.0, -0.2, 0.2, 1}
	var stringArray = []string{"4", "3", "4", "1.231", "0", "-0.2", "   0.2", "\n1"}
	floatArray := stringArrayToFloat(stringArray)
	for i, value := range floatArray {
		assert.Equal(t, expectedValues[i], value, "Item number %v had a wrong value", i)
	}
	assert.Equal(t, 0, len(stringArrayToFloat([]string{})), "Empty array failed!")
}

func TestCplexOutputToArray(t *testing.T) {
	var expectedProbValues = []string{"0", "0.333333", "0", "0", "0", "0.333333", "0.166667", "0", "0", "0.166667", "0", "0"}

	valueString := "Values = [0, 0.333333, 0, 0, 0, 0.333333, 0.166667, 0, 0, 0.166667,\n0, 0]"
	newValues := cplexOutputToArray(valueString)

	for i, value := range newValues {
		assert.Equal(t, expectedProbValues[i], strings.TrimSpace(value), "Item number %v had a wrong value", i)
	}
}

func TestUpdateNodeWrongAmount(t *testing.T) {
	c, m, d := setUpTest()
	node := matching.FindFeasibleMatching(m, 0, 3, &c)
	_ = d

	var values = []float64{1, 2, 3, 4, 5}

	assert.Panics(t, func() {
		updateNode(node, values)
	}, "Calling updateNode with an amount that does not match the adjacency matrix, it should fail")
}

func buildSlice(n int) []float64 {
	var values []float64
	k := 0.0

	for len(values) < n {
		values = append(values, k)
		k++
	}

	return values
}

func TestUpdateNode(t *testing.T) {
	c, m, _ := setUpTest()
	node := matching.FindFeasibleMatching(m, 0, 3, &c)

	adjlen := len(node.Adj) * len(node.Adj[0])

	values := buildSlice(adjlen)

	updateNode(node, values)

	k := 0.0
	for i := range node.Adj {
		for j, edge := range node.Adj[i] {
			if (i == 0 && j == 0) {
				assert.Equal(t, k, edge.Prob, "Edge.Prob at Adj[%v][%v] had a wrong value", i, j)
				assert.False(t, edge.Basic, "Edge.Basic at Adj[%v][%v] had a wrong value", i, j)
			} else {
				assert.Equal(t, k, edge.Prob, "Edge.Prob at Adj[%v][%v] had a wrong value", i, j)
				assert.True(t, edge.Basic, "Edge.Basic at Adj[%v][%v] had a wrong value", i, j)
			}
			k++
		}
	}
	assert.Equal(t, 11, node.BasicCount, "node.BasicCount had a wrong value")
}
