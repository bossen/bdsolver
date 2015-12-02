package lp

import (
		"testing"
		"github.com/stretchr/testify/assert"
		"coupling"
		"markov"
		"earthmover"
		"utils"
		"sync"
		"strings"
		"bytes"
		"log"
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
	var expectedBasicValues = []bool{false, true, false, false, false, true, true, false, false, true, false, false}

func TestOptimize(t *testing.T) {
	c, m, d := setUpTest()
	node := earthmover.FindFeasibleMatching(m, 0, 3, &c)
	optimize(m, node, d, 0.0, 0, 0)
	k := 0
	for i := range node.Adj {
		for _, j := range node.Adj[i] {
			assert.True(t, utils.ApproxEqual(j.Prob, expectedProbValues[k]), "optimize does not work for %v, probability", j)
			assert.Equal(t, expectedBasicValues[k], j.Basic, "optimize does not work for %v, basic value", j)
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

}

func TestCplexOptimize(t *testing.T) {
	var expectedProbValues = []float64{0.0, 0.333333, 0.0, 0.0, 0.0, 0.333333, 0.166667, 0.0, 0.0, 0.166667, 0.0, 0.0}

	newValues := cplexOptimize("1 0 1 1 1 0 1 1 1 1 1 1", " 0.3333333333333333 0.3333333333333333 0.1666666666666667 0.1666666666666667 0.3333333333333333 0.3333333333333333 0.3333333333333333", 4, 3)
	for i, value := range newValues {
		assert.Equal(t, expectedProbValues[i], value, "Item number %v had a wrong value", i)
	}
}

func TestCplexOutputToArray(t *testing.T) {
	var expectedProbValues = []string{"0", "0.333333", "0", "0", "0", "0.333333", "0.166667", "0", "0", "0.166667", "0", "0"}

	valueString := "Values = [0, 0.333333, 0, 0, 0, 0.333333, 0.166667, 0, 0, 0.166667,\n0, 0]"
	newValues := cplexOutputToArray(valueString)

	for i, value := range newValues {
		assert.Equal(t, expectedProbValues[i], strings.TrimSpace(value), "Item number %v had a wrong value", i)
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

func TestAppendDValues(t *testing.T) {
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
