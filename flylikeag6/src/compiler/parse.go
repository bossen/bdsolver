package compiler

import (
    "log"
)


func (c *Compiler) scanner() token {
    // TODO clear whitespace
    if c.peek() == 'S' { // Must be States
        return c.maketoken(c.readword(), "")
    } else if c.peek() == 'E' {
        return c.maketoken(c.readword(), "")
    } else {
        log.Fatal("Unexpected keyword ", c.peek())
    }
    return token{}
}


func (c *Compiler) Parse() {
    state := 0
    for true {
        token := c.scanner()

        if state == 0 { // We are before States:
            if token.tokentype != "States" {
                log.Fatal("Expected States found ", token.tokentype)
            }

        } else if state == 1 { // We are in States

        } else if state == 2 {  // We are in Edges

        } else {
            log.Fatal("Something went wrong while parsing")
        }


        break
    }
}
