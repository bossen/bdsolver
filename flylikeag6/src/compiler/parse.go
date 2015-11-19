package compiler

import (
    "log"
)

func isAlphabic(a byte) bool {
    return (a > 'a'  && a < 'z' ) || (a > 'A' && a < 'Z')
}


func (c *Compiler) scanner() token {
    // TODO clear whitespace
    if c.peek() == 'S' { // Must be States
        return c.maketoken(c.readword(), "")
    } else if c.peek() == 'E' {
        return c.maketoken(c.readword(), "")
    else if isAlphabetic(c.peek()) {
        return c.maketoken("word",  c.readword())

    else if isNumeric(c.peek()) {
        return c.maketoken("numeric", c.readnumeric())
    }
    } else {
        log.Fatal("Unexpected keyword ", c.peek())
    }
    return token{}
}

func (c *Compiler) getExpectedToken(toexpect string) string {
    token := c.scanner()
    expectEqual(token.tokentype, toexpect)
    return token.value
}

func expectEqual(found, expected string) {
    if token.tokentype != toexpect {
        log.Fatal("Expected %s got %s", toexpect, token.tokentype)
    }

}

func (c *Compiler) Parse() {
    state := 0
    for true {
        if state == 0 { // We are before States:
            getAndExpect("States:")
            state = 1
        } else if state == 1 { // We are in States
            token := scanner()

            if token.tokentype == "Edges" {
                state = 2
                continue
            } else if token.tokentype == "numeric" {
                label := c.scanner()
                // TODO create state token.value with label.label.value
            }

            log.Fatal("Expected Edges found ", token.tokentype)

        } else if state == 2 {  // We are in Edges
            token := scanner()
            if token.tokentype == "eof" {
                break
            }

            expectEqual(token.tokentype, "word")
            from := token.value
            getExpectedToken("to")
            to := getExpectedToken("word")
            prob := getExpectedToken("float")


        } else {
            log.Fatal("Something went wrong while parsing")
        }

        break
    }
}
