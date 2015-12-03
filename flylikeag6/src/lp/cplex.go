package lp

import (
	"strconv"
	"log"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
	"regexp"

	"markov"
	"coupling"
	"utils"
)

//credits: https://stackoverflow.com/questions/20437336/how-to-execute-system-command-in-golang-with-unknown-arguments
func exeCmd(cmd string, wg *sync.WaitGroup) string {
  // splitting head => g++ parts => rest of the command
  parts := strings.Fields(cmd)
  head := parts[0]
  parts = parts[1:len(parts)]

  out, err := exec.Command(head,parts...).Output()
  if err != nil {
		log.Printf("Error in exeCmd: %s", err)
  }
  wg.Done() // Need to signal to waitgroup that this goroutine is done
	return fmt.Sprintf("%s", out)
}

//Retrieve constraints from the transitions, and adds them to the buffer given.
//Returns the number of constraints, and an array with indexes of the used elements
func findConstraints(buffer *bytes.Buffer, transitions []float64) (int, []int) {
	size := 0
	n := len(transitions)
	used := make([]int, n, n)

	for i, constraint := range transitions {
		if (constraint > 0.0) {
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
	for _, row := range rowused {
		for _, column := range columnused {
			u, v := utils.GetMinMax(row, column) //The same d value may be used two times
			log.Printf("Retrieving d[%v][%v] = %v", u, v, d[u][v])
			(*buffer).WriteString(strconv.FormatFloat(d[u][v], 'f', -1, 64))
			(*buffer).WriteString(" ")
		}
	}
}

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

func optimize(dbuffer, constraints string, rowcount, columncount int) []float64 {
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

func updateNode(node *coupling.Node, newValues []float64) {
	if (len(node.Adj) * len(node.Adj[0])) != len(newValues) {
		panic("The amount of new values does not match the adjacency matrix!")
	}

	k := 0
	for i := range (*node).Adj {
		for _, edge := range (*node).Adj[i] {
			//TODO increment or decrement basiccount
			prob := newValues[k]
			if utils.ApproxEqual(prob, 0.0) {
				edge.Basic = false
			} else {
				edge.Basic = true
			}
			edge.Prob = prob
			k++
		}
	}
}

func CplexOptimize(markov markov.MarkovChain, node *coupling.Node, d [][]float64, min float64, i, j int) {
	var dbuffer, constraints bytes.Buffer
	rowcount, rowused := findConstraints(&constraints, markov.Transitions[node.S])
	columncount, columnused := findConstraints(&constraints, markov.Transitions[node.T])

	log.Printf("Row indexes from transition matrix : %v size: %v", rowused, rowcount)
	log.Printf("Column indexes from transition matrix: %v size: %v", columnused, columncount)
	retrieveDValues(&dbuffer, d, rowused, columnused)

	newValues := optimize(dbuffer.String(), constraints.String(), rowcount, columncount)
	log.Printf("Updating node with new values: %v", newValues)
	updateNode(node, newValues)
	//TODO recoverBasicNodes(node)
}
