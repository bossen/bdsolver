package compiler

import "os"


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
