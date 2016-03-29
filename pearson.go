package main

// Collection of functions designed to get the pearson similarity between two arrays of values

import "math"

// given two floats returns a result to be used as denominator in func pearson
func pDen(v1 []int, v2 []int, c chan float64) {

	sumV1 := float64(sum(v1))
	denV1 := math.Pow(sumV1, 2) / float64(len(v1))

	sumV2 := float64(sum(v2))
	denV2 := math.Pow(sumV2, 2) / float64(len(v2))

	sumSqV1 := sumSquares(v1)
	sumSqV2 := sumSquares(v2)

	den := math.Sqrt((sumSqV1 - denV1) * (sumSqV2 - denV2))

	if math.IsNaN(den) {
		den = 0
	}
	c <- den
}

// get the numerator for the pearson's coefficient
func pNum(v1, v2 []int, c chan float64) {
	sumV1 := float64(sum(v1))
	sumV2 := float64(sum(v2))
	num := sumProducts(v1, v2) - (sumV1 * sumV2 / float64(len(v1)))
	c <- num
}

// A concurrent implementation of the func pearson
func pearsonConc(v1, v2 []int) float64 {

	d := make(chan float64)
	go pDen(v1, v2, d)
	den := <-d
	if den == 0 {
		return 0
	}

	n := make(chan float64)
	go pNum(v1, v2, n)
	num := <-n

	return (1.0 - num/den)
}

func pearson(v1, v2 []int) float64 {

	sumV1 := float64(sum(v1))
	sumV2 := float64(sum(v2))

	sumSqV1 := sumSquares(v1)
	sumSqV2 := sumSquares(v2)

	pSum := sumProducts(v1, v2)

	// Calculate r (Pearson score)
	var num float64
	var den float64

	num = pSum - (float64(sumV1) * float64(sumV2) / float64(len(v1)))

	denV1 := math.Pow(sumV1, 2) / float64(len(v1))
	denV2 := math.Pow(sumV2, 2) / float64(len(v2))
	den = math.Sqrt((sumSqV1 - denV1) * (sumSqV2 - denV2))

	if den == 0 || math.IsNaN(den) {
		return 0
	}
	return (1.0 - num/den)

}

func sumSqCh(v []int, c chan float64) {
	var sumSq float64
	for i := range v {
		value := float64(v[i])
		sumSq += math.Pow(value, 2)
	}
	c <- sumSq
}

// splits the given array into two to process separately
func sumSqFast(v []int) float64 {
	c := make(chan float64)
	go sumSqCh(v[:len(v)/2], c)
	go sumSqCh(v[len(v)/2:], c)
	x, y := <-c, <-c
	return x + y
}

// Given a vector of numbers will sum the square roots
func sumSquares(v []int) float64 {
	var sumSq float64
	for i := range v {
		value := float64(v[i])
		sumSq += math.Pow(value, 2)
	}
	return sumSq
}
