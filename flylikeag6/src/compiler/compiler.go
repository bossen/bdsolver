package compiler

import (
    "log"
    "scanner"
    "utils"
    "strconv"
    "markov"
)

type token struct {
    tokentype string
    value string
}

func scan(c* scanner.Scanner) token {
    c.EatWhitespaceAndComments()
    if c.Peek() == '/' {
        c.ReadChar()
        return token{"Divide", ""}
    } else if c.EndOfFile() {
        log.Println("Found EOF")
        return token{"EOF", ""}
    } else if c.Peek() == '-' {
        c.ReadChar()
        chr := c.ReadChar()
        if chr != '>' {
            c.Fail("Found -, and expected a \"to\" statement(->) but found ", chr)
        }
        return token{"To", ""}
    } else if utils.IsAlphabetic(c.Peek()) {
        word := c.ReadWord()
        if word == "States" {
            return token{"States", ""}
        } else if word == "Edges" {
            return token{"Edges", ""}
        }
        return token{"Word", word}

    } else if utils.IsNumeric(c.Peek()) {
        integer := strconv.Itoa(c.ReadNumber())
        return token{"Integer", integer}
    } else {
        c.Fail("Unexpected keyword ", c.Peek())
    }
    return token{}
}

func getExpectedToken(c* scanner.Scanner, expected string) string {
    token := scan(c)
    expectToken(token, expected)
    return token.value
}

func expectToken(found token, expected string) {
    if found.tokentype != expected {
        log.Fatal("Expected ", expected, " found ", found.tokentype, " with value '", found.value, "'")
    }

}

func IndexInSlice(slice []string, value string) int {
    for i, v := range slice {
        if v == value {
            return i
        }
    }
    return -1
}

func Parse(filename string) markov.MarkovChain {
    log.Println("Parsing file ", filename)
    c, _ := scanner.New(filename)
    state := 0
    numberofmcstates := 0
    labelmapper := make([]string, 0,0)
    labels := make([]int, 0,0)
    var transitions [][]float64 // We only know this when we are done in state 1, so therefore we initialize it later.
    for true {
        c.EatWhitespaceAndComments()
        if c.EndOfFile() {
            if state != 2 {
                c.Fail("Expected to have seen some edges.")
            }
            break
        }
        token := scan(&c)

        if state == 0 { // We are before States:
            expectToken(token, "States")
            log.Printf("Going to states")
            state = 1
            continue
        } else if state == 1 { // We are in States
            if token.tokentype == "Edges" {
                state = 2
                log.Printf("Going to Edges")
                transitions = make([][]float64, numberofmcstates, numberofmcstates)
                for i  := range transitions {
                    transitions[i] = make([]float64, numberofmcstates, numberofmcstates)
                }
                continue
            } else if token.tokentype == "Integer" {
                label := getExpectedToken(&c, "Word")

                numberofmcstates += 1
                if strconv.Itoa(numberofmcstates) != token.value {
                    c.Fail("Expected state ", numberofmcstates, " but found ", token.value)
                }

                // In case we found a new label.
                if IndexInSlice(labelmapper, label) == -1 {
                    labelmapper = append(labelmapper, label)
                    labelmapper[len(labelmapper)-1] = label
                }

                // Setting our label transformer
                labels = append(labels, IndexInSlice(labelmapper, label))

                log.Printf("Found state %s with level %s", token.value, label)
                continue
            }

            c.Fail("Expected Edges found a %s %s", token.tokentype, token.value)

        } else if state == 2 {  // We are in Edges

            expectToken(token, "Integer")
            // The reason why it is not nessecary to check for errors, is
            // that the scanner checks if it is a integer.
            expectToken(token, "Integer")
            from, _ := strconv.Atoi(token.value)
            getExpectedToken(&c, "To")
            to, _ := strconv.Atoi(getExpectedToken(&c, "Integer"))
            num, _ := strconv.Atoi(getExpectedToken(&c, "Integer"))
            getExpectedToken(&c, "Divide")
            den, _ := strconv.Atoi(getExpectedToken(&c, "Integer"))

            // TODO make check that state exists
            transitions[from - 1][to - 1] = float64(num) / float64(den)

            log.Printf("Edge %d -> %d with prop: %d / %d", from, to, num, den)
            continue

        } else {
            c.Fail("Something went wrong while parsing")
        }
        c.Fail("A programmer error. Someone forgot a continue")
    }

    log.Printf("%+v", labelmapper)
    log.Printf("%+v", labels)
    log.Printf("%+v", transitions)
    return markov.MarkovChain{labels, transitions}
}
