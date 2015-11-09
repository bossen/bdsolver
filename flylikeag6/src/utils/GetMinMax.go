package utils

func GetMinMax(u, v int) (int, int) {
	if v < u {
		return v, u
	}
	return u, v
}
