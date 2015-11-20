package utils


func IsAlphabetic(a byte) bool {
    return ('a' <= a && a <= 'z') || ('A' <= a && a <= 'Z')
}
