package scanner

import (
	"github.com/stretchr/testify/assert"
	"testing"
    "strings"
    "bufio"
)




func TestReadChar(t *testing.T) {
    scanner, _ := New("main.go")

    // Overwrite the reader, so we can fake it
    scanner.reader = *bufio.NewReader(strings.NewReader("someword"))
    assert.True(t, scanner.ReadChar() == 's', "Read unexpected byte, should be 's'")
    assert.True(t, scanner.ReadChar() == 'o', "Read unexpected byte, should be 'o'")
    assert.True(t, scanner.ReadChar() == 'm', "Read unexpected byte, should be 'm'")
}

func TestReadNumber(t *testing.T) {
    scanner, _ := New("main.go")

    // Overwrite the reader, so we can fake it
    scanner.reader = *bufio.NewReader(strings.NewReader("321321"))

    assert.Equal(t, scanner.ReadNumber(), 321321)

}


func TestReadNumberWithPrefixWhitespace(t *testing.T) {
    scanner, _ := New("main.go")

    // Overwrite the reader, so we can fake it
    scanner.reader = *bufio.NewReader(strings.NewReader("\t    \t321321"))

    assert.Equal(t, scanner.ReadNumber(), 321321)

}

func TestReadNumberWithSufixWhitespace(t *testing.T) {
    scanner, _ := New("main.go")

    // Overwrite the reader, so we can fake it
    scanner.reader = *bufio.NewReader(strings.NewReader("321321\t   "))

    assert.Equal(t, scanner.ReadNumber(), 321321)

}


func TestReadWord(t *testing.T) {
    scanner, _ := New("main.go")

    // Overwrite the reader, so we can fake it
    scanner.reader = *bufio.NewReader(strings.NewReader("someword"))

    assert.Equal(t, scanner.ReadWord(), "someword")

}

func TestReadWordWithPrefixWhitespace(t *testing.T) {
    scanner, _ := New("main.go")

    // Overwrite the reader, so we can fake it
    scanner.reader = *bufio.NewReader(strings.NewReader("\t    someword"))

    assert.Equal(t, scanner.ReadWord(), "someword")

}

func TestReadWordWithSufixWhitespace(t *testing.T) {
    scanner, _ := New("main.go")

    // Overwrite the reader, so we can fake it
    scanner.reader = *bufio.NewReader(strings.NewReader("someword\t    "))

    assert.Equal(t, scanner.ReadWord(), "someword")

}
