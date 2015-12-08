package utils


type IntPair struct {
	I, J int
}


func IsIntPairInSlice(pair IntPair, pairs []IntPair) bool {
    for _, n := range pairs {
        if n == pair {
            return true
        }
    }
    return false
}
