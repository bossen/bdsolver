package compiler

import (
    "log"
    "os"
)


type token struct {
    tokentype string
    val string
}


func (c *Compiler)  maketoken(tokentype, value string) token {
    validtokentypes := []string {
        "States",
        "Edges",
    }

    for _, vtokens := range validtokentypes {
        if vtokens == tokentype {
            return token{tokentype, value}
        }
    }

    log.Fatal("Did not find expected keyword, found " + tokentype)
    return token{}
}


func (c *Compiler) scanner() token {
    // TODO clear whitespace
    chr := c.readchar()
    if chr  == 'S' { // Must be States
        return c.maketoken("S" + c.readword(), "")
    }
    return token{}
}

func New(filename string) (Compiler, error) {
    filelocal, err := os.Open(filename)
    return Compiler{*filelocal}, err
}

func (c *Compiler) Parse() {


    state := 0
    for true {
        token := c.scanner()
        _ = token 

        if state == 0 { // We are before States:
        } else if state == 1 {

        } else if state == 2 {

        } else {
            log.Fatal("Something went wrong while parsing")
        }
    }
}
