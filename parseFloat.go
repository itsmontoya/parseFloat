package parseFloat

// THIS IS A ROUGH PROOF OF CONCEPT TO TEST SOME IDEAS AND THEORIES
// STUFF IS MISSING, LOTS OF STUFF. DON'T USE THIS AT ALL. :|

import (
	//	"fmt"
	"math"
)

const (
	char0 = '0'
	char1 = '1'
	char2 = '2'
	char3 = '3'
	char4 = '4'
	char5 = '5'
	char6 = '6'
	char7 = '7'
	char8 = '8'
	char9 = '9'
)

// Parse will parse
func Parse(s []byte) (n float64) {
	var (
		j float64
		i = len(s) - 1
	)

	for i > -1 {
		var val float64
		b := s[i]

		switch b {
		case char0:
			val = 0
		case char1:
			val = 1
		case char2:
			val = 2
		case char3:
			val = 3
		case char4:
			val = 4
		case char5:
			val = 5
		case char6:
			val = 6
		case char7:
			val = 7
		case char8:
			val = 8
		case char9:
			val = 9
		}

		n += val * math.Pow(10, j)
		i--
		j++
	}

	return
}
