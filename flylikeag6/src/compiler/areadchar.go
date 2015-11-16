package compiler



func (c *Compiler)readchar() byte {
    buf := make([]byte, 1, 1)
    c.file.Read(buf)
    return buf[0]
}
