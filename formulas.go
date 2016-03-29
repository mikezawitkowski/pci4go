package main

/*
formulas.go

These are Go versions of some of the mathematical formulas
found in Appendix B of Programming Collective Intelligence.

The following formulas are covered elsewhere and not
included in this file:
 - Euclidean Distance is in euclidean.go (for []int) and recommendations.go: simDistance() (for [string]float)
 - Pearson Correlation Coefficient is in pearson.go
 - Tanimoto is in tanimoto.go

*/

// dotProduct multiplies two vectors
func dotProduct(a, b []int) int {
	max := len(a)
	if len(b) > max {
		max = len(b)
	}

	result := 0 // will hold the aggregate for each round
	for i := 0; i < max; i++ {
		result = result + a[i]*b[i]
	}
	return result
}
