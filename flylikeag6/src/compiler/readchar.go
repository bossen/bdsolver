package compiler

import (
    "log")

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
