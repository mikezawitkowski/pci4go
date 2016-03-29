package main

// returns the sum of an array of ints
func sum(in []int) int {
	var result int
	for i := 0; i < len(in); i++ {
		result += in[i]
	}
	return result
}

func sumProducts(v1, v2 []int) float64 {
	var pSum float64
	for i := range v1 {
		val1 := float64(v1[i])
		val2 := float64(v2[i])
		pSum += val1 * val2
	}
	return pSum
}
