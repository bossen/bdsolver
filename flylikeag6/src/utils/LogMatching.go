package utils

import (
    "coupling"
    "log"
)


func LogMatching(matching [][]*coupling.Edge, lenrow int, lencol int) {

	for i := 0; i < lenrow; i++ {
		for j := 0; j < lencol; j++ {
			log.Println("At: u and v", matching[i][j].To.S, matching[i][j].To.T)
			log.Println(matching[i][j].Prob)
			log.Println(matching[i][j].Basic)
		}
	}
}
