package utils

import (
    "coupling"
    "log"
)


func LogMatching(matching [][]*coupling.Edge, lenrow int, lencol int) {
    log.Println("Logging matching:")
	for i := 0; i < lenrow; i++ {
		for j := 0; j < lencol; j++ {
            log.Printf(" - At: u %d, %d prob: %f, basic: %t",
                matching[i][j].To.S,
                matching[i][j].To.T,
			    matching[i][j].Prob,
			    matching[i][j].Basic)
		}
	}
}
