package scanner

import (
    "os"
    "log"
    "strconv"
    "utils"
    "bufio"
)


type Scanner struct {
    reader bufio.Reader
    file os.File
    line int
}

func New(filename string) (Scanner, error) {
    c := Scanner{}

    filelocal, err := os.Open(filename)
    if err != nil {
        return c, err
    }

    c.file = *filelocal

    c.reader = *bufio.NewReader(filelocal)
    return c, nil
}

func (c *Scanner) Peek() byte {
    char, err := c.reader.Peek(1)

    if err != nil {
        log.Fatal("Unexpected end of file")
    }

    return char[0]
}

func (c *Scanner) ReadChar() byte {
    char := make([]byte, 1, 1)

    n, err := c.reader.Read(char)

    if n == 0 || err != nil {
        log.Fatal("Unexpected end of file")
    }

    return char[0]
}

func (c *Scanner) endoffile() bool {
    _, err := c.reader.Peek(1)
    // There might be other cases, e.g. buffer full.
    if err == nil {
        return false
    } else {
        return true
    }
}

func (c *Scanner) eatuntil(a byte) {

    for !c.endoffile()  {
        chr :=  c.ReadChar()
        if chr == a {
            break
        }
    }
    if c.endoffile() {
        log.Fatal("Unexpected end of file")
    }
}

func (c *Scanner) EatWhitespaceAndComments() {
    for !c.endoffile() {
        if c.Peek() == ' ' || c.Peek() == '\t' {
            c.ReadChar()
        } else {
            break
        }
    }

    if !c.endoffile() && c.Peek() == '/' {
        char := c.ReadChar()
        if char != '/' {
            log.Fatal("Expected / got %s", char)
        }
        expect(c.ReadChar(), '/')
        c.eatuntil('\n')
    }
}



func (c *Scanner) ReadNumber() int {
    c.EatWhitespaceAndComments()

    number := ""
    for !c.endoffile() {
        if utils.IsNumeric(c.Peek()) {
            number += string(c.ReadChar())
        } else if utils.IsWhitespace(c.Peek()){
            break
        } else {
            log.Fatal("Expected whitespace after the number " + number + ", but got ", c.Peek())
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

func (c *Scanner) ReadWord() string {
    c.EatWhitespaceAndComments()
    word := ""
    for !c.endoffile() {
        if utils.IsAlphabetic(c.Peek()) {
            word += string(c.ReadChar())
        } else if utils.IsWhitespace(c.Peek()) {
            break
        } else {
            log.Fatal("Unexpected %s", c.Peek())
        }
    }
    return word
}

func (c *Scanner) Close() {
    c.file.Close()
}
