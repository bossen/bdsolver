package compiler

import (
    "log"
    "strconv"
)

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

