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
    c.line = 1
    return c, nil
}

func (c *Scanner) Fail(m ...interface{}) {
    log.Fatal("In line ", c.line, ": ", m)
}

func (c *Scanner) Peek() byte {
    char, err := c.reader.Peek(1)

    if err != nil {
        c.Fail("Unexpected end of file")
    }

    return char[0]
}

func (c *Scanner) ReadChar() byte {
    char := make([]byte, 1, 1)

    n, err := c.reader.Read(char)

    if n == 0 || err != nil {
        c.Fail("Unexpected end of file")
    }

    if char[0] == '\n' {
        c.line += 1
    }

    return char[0]
}

func (c *Scanner) EndOfFile() bool {
    _, err := c.reader.Peek(1)
    // There might be other cases, e.g. buffer full.
    if err == nil {
        return false
    } else {
        return true
    }
}

func (c *Scanner) eatuntil(a byte) {

    for !c.EndOfFile()  {
        chr :=  c.ReadChar()
        if chr == a {
            break
        }
    }
    if c.EndOfFile() {
        c.Fail("Unexpected end of file")
    }
}

func (c *Scanner) EatWhitespaceAndComments() {
    for !c.EndOfFile() {
        if utils.IsWhitespace(c.Peek()) {
            c.ReadChar()
        } else {
            break
        }
    }

    chars, err := c.reader.Peek(2)

    _ = err
    // TODO check error.

    if !c.EndOfFile() && chars[0] == '/' && chars[1] == '/' {
        c.ReadChar()
        c.ReadChar()
        c.eatuntil('\n')
    }
}


func (c *Scanner) LineNumber() int {
    return c.line
}



func (c *Scanner) ReadNumber() int {
    c.EatWhitespaceAndComments()

    number := ""
    for !c.EndOfFile() {
        if utils.IsNumeric(c.Peek()) {
            number += string(c.ReadChar())
        } else if utils.IsWhitespace(c.Peek()) || c.Peek() == '/' {
            break
        } else {
            c.Fail("Expected whitespace after the number " + number + ", but got ", string(c.Peek()))
        }
    }

    if len(number) == 0 {
        c.Fail("Did not read any number")
    }

    numbasint, err := strconv.Atoi(number)
    if err != nil {
        c.Fail("Could not convert to number. Something really bad happened, and should never happen. Please contact the developers at https://github.com/jbossen/P7-code")
    }

    return numbasint
}

func (c *Scanner) ReadWord() string {
    c.EatWhitespaceAndComments()
    word := ""
    for !c.EndOfFile() {
        if utils.IsAlphabetic(c.Peek()) {
            word += string(c.ReadChar())
        } else if utils.IsWhitespace(c.Peek()) || c.Peek() == '/' {
            break
        } else {
            c.Fail("Reading a word failed, found '", word, "' but then unexpected ", string(c.Peek()))
        }
    }
    return word
}

func (c *Scanner) Close() {
    c.file.Close()
}
