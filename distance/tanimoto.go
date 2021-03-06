package main

import "math"

// Tanimoto is more useful for where values are 1s and 0s
// Given two vectors such as [0,1,0,0...] returns similarity score
func tanimoto(v1, v2 []bool) float64 {
	var result float64
	c1 := 0.0
	c2 := 0.0
	shr := 0.0

	max := len(v2)
	if len(v1) > max {
		max = len(v1)
	}

	// Loop through and increment values
	for i := 0; i < max; i++ {
		if v1[i] != false {
			c1++ // in v1
		}
		if v2[i] != false {
			c2++                // in v2
			if v1[i] != false { // If both != 0
				shr++ // increment shr
			}
		}
	}
	result = math.Abs(1.0 - (float64(shr) / (c1 + c2 - shr)))
	// log.Println("result:", result)
	return result
}
