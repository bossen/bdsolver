package compiler

import (
    "os"
    "log"
    "strconv"
)

type token struct {
    tokentype string
    val string
}


type Compiler struct {
    lastchar []byte
    uselastchar bool
    file os.File
    line int
}

func New(filename string) (Compiler, error) {
    filelocal, err := os.Open(filename)
    return Compiler{
        make([]byte, 1,1),
        false,
        *filelocal,
        0}, err
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


func (c *Compiler) peek() byte {
    if !c.uselastchar {
        c.uselastchar = true
        c.readchar()
    }
    return c.lastchar[0]
}

func (c *Compiler) readchar() byte {
    if c.uselastchar {
        c.uselastchar = false
        return c.lastchar[0]
    }

    n, _ := c.file.Read(c.lastchar)

    if c.lastchar[0] == '\n' {
        c.line += 1
    }

    if n == 0 {
        log.Fatal("Unexpected end of file")
    }

    return c.lastchar[0]
}

func (c *Compiler) endoffile() bool {
    // TODO Implement
    return true
}

func expect(a, b byte) {
    if a != b {
        log.Fatal("Expected %s got %s", a, b)
    }
}

func (c *Compiler) eatuntil(a byte) {

    for !c.endoffile()  {
        chr :=  c.readchar()
        if chr == a {
            break
        }
    }
    if c.endoffile() {
        log.Fatal("Unexpected end of file")
    }
}

func (c *Compiler) eatwhitespaceandcomments() {
    for !c.endoffile() {
        if c.peek() == ' ' || c.peek() == '\t' {
            c.readchar()
        }
    }

    if !c.endoffile() && c.peek() == '/' {
        expect(c.readchar(), '/')
        c.eatuntil('\n')
    }
}


func (c *Compiler) readnumber() int {
    c.eatwhitespaceandcomments()

    number := ""
    for !c.endoffile() {
        if !(c.peek() > '0' && c.peek() < '9') {
            if c.peek() == ' ' || c.peek() == '\t' || c.peek() == '\n' {
                break
            } else {
                log.Fatal("Expected whitespace after number, but got %s", c.peek())
            }
        }
    }
    if len(number) == 0 {
        log.Fatal("Did not read any number")
    }

    numbasint, err := strconv.Atoi(number)
    if err != nil {
        log.Fatal("Could not convert to number. Something really bad happened")
    }

    return numbasint
}

func (c *Compiler) readword() string {
    c.eatwhitespaceandcomments()
    word := ""
    for !c.endoffile() {
        if (c.peek() > 'A' &&  c.peek() < 'Z') || (c.peek() > 'a' && c.peek() < 'z') {
            word += string(c.readchar())
        } else if c.peek() == ' ' || c.peek() == '\t' || c.peek() ==  '\n' {
            break
        } else {
            log.Fatal("Unexpected %s", c.peek())
        }
    }
    return word
}
