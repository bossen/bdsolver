package lp

import (
	"markov"
	"coupling"
	"strconv"
	"utils"
	"log"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
	"regexp"
)

//https://stackoverflow.com/questions/20437336/how-to-execute-system-command-in-golang-with-unknown-arguments
func exeCmd(cmd string, wg *sync.WaitGroup) string {
  // splitting head => g++ parts => rest of the command
  parts := strings.Fields(cmd)
  head := parts[0]
  parts = parts[1:len(parts)]

  out, err := exec.Command(head,parts...).Output()
  if err != nil {
    log.Printf("%s", err)
  }
  wg.Done() // Need to signal to waitgroup that this goroutine is done
	return fmt.Sprintf("%s", out)
}

func findConstraints(buffer *bytes.Buffer, transitions []float64) (int, []int) {
	size := 0
	n := len(transitions)
	used := make([]int, n, n)
	for i, constraint := range transitions {
		if (!utils.ApproxEqual(constraint, 0.0)) {
			(*buffer).WriteString(strconv.FormatFloat(constraint, 'f', -1, 64))
			(*buffer).WriteString(" ")
			used[size] = i
			size++
		}
	}
	return size, used[:size]
}

//Retrieves d values into the buffer given, based on indexes from rowused and columnused
func retrieveDValues(buffer *bytes.Buffer, d [][]float64, rowused, columnused []int) {
	for _, i := range rowused {
		for _, j := range columnused {
			u, v := utils.GetMinMax(i, j)
			log.Printf("d[%v][%v] = %v", u, v, d[u][v])
			(*buffer).WriteString(strconv.FormatFloat(d[u][v], 'f', -1, 64))
			(*buffer).WriteString(" ")
		}
	}
}

//https://stackoverflow.com/questions/27875479/how-to-convert-string-to-float64-in-golang
func stringArrayToFloat(input []string) []float64 {
	n := len(input)
	numbers := make([]float64, n, n)

	for i, value := range input {
		number := strings.TrimSpace(value)
		if n, err := strconv.ParseFloat(number, 64); err == nil {
			numbers[i] = n
		} else {
			panic(err)
		}
	}
	return numbers
}

func cplexOutputToArray(input string) []string {
	re := regexp.MustCompile(`(\d+(\.\d+)?,\s)*(\d+(\.\d+)?)`)
	match := re.FindStringSubmatch(input)
	if (len(match) == 0) {
		fmt.Println("Error in cplex output")
		os.Exit(1)
	}

	output := strings.Split(match[0], ",")
	return output
}

func cplexOptimize(dbuffer, constraints string, rowcount, columncount int) []float64 {
	log.Printf("Running program.out %v %v %v%v", rowcount, columncount, dbuffer, constraints)
	wg := new(sync.WaitGroup)
	wg.Add(1)

	command := fmt.Sprintf("./../../../cplex/program.out %v %v %v%v", rowcount, columncount, dbuffer, constraints)
	output := exeCmd(command, wg)
	wg.Wait()
	err, _ := regexp.MatchString("Failed", output)
	if err {
		fmt.Println("Cplex failed")
		os.Exit(1)
	}
	outputArray := cplexOutputToArray(output)
	floatArray := stringArrayToFloat(outputArray)
	return floatArray
}

func updateNode(node *coupling.Node, newValues []float64, rowcount, columncount int) {
	k := 0
	for i := range (*node).Adj {
		for _, j := range (*node).Adj[i] {
			//TODO increment or decrement basiccount
			prob := newValues[k]
			if utils.ApproxEqual(prob, 0.0) {
				j.Basic = false
			} else {
				j.Basic = true
			}
			j.Prob = prob
			k++
		}
	}
}

func optimize(markov markov.MarkovChain, node *coupling.Node, d [][]float64, min float64, i, j int) {
	var dbuffer, constraints bytes.Buffer
	rowcount, rowused := findConstraints(&constraints, markov.Transitions[node.S])
	columncount, columnused := findConstraints(&constraints, markov.Transitions[node.T])

	log.Printf("%v %v", rowused, rowcount)
	log.Printf("%v %v", columnused, columncount)
	retrieveDValues(&dbuffer, d, rowused, columnused)

	newValues := cplexOptimize(dbuffer.String(), constraints.String(), rowcount, columncount)
	log.Printf("%v", newValues)
	updateNode(node, newValues, rowcount, columncount)
	//TODO recoverBasicNodes(node)
}
