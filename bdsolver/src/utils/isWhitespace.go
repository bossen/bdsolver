package utils

func IsWhitespace(a byte) bool {
    return a == ' ' || a == '\t' || a == '\n' || a == '\r'
}
