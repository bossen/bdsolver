package compiler

import (
    "log"
    "scanner"
    "utils"
    "strconv"
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
            log.Fatal("Found -, and expected a \"to\" statement(->) but found ", chr)
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
        log.Fatal("Unexpected keyword ", c.Peek())
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

func Parse(filename string) {
    log.Println("Parsing file ", filename)
    c, _ := scanner.New(filename)
    state := 0
    for !c.EndOfFile() {
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
                continue
            } else if token.tokentype == "Integer" {
                label := getExpectedToken(&c, "Word")
                log.Printf("Found state %s with level %s", token.value, label)
                continue
            }

            log.Fatal("Expected Edges found a %s %s", token.tokentype, token.value)

        } else if state == 2 {  // We are in Edges

            expectToken(token, "Integer")
            from := token.value
            getExpectedToken(&c, "To")
            to := getExpectedToken(&c, "Integer")
            num := getExpectedToken(&c, "Integer")
            getExpectedToken(&c, "Divide")
            den := getExpectedToken(&c, "Integer")

            log.Printf("Edge %s -> %s with prop: %s / %s", from, to, num, den)
            continue

        } else {
            log.Fatal("Something went wrong while parsing")
        }
        log.Fatal("A programmer error. Someone forgot a continue")

    }
}
