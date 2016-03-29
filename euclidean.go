package main

import "math"

// Get the euclidean distance between two points
/*
if points are:
[xa, ya, za]
(xb, yb, zb)
then dist = sqrt((xa-xb)^2 + (za-zb)^2)
*/
func euclidean(v1, v2 []int) float64 {
	sumSq := 0.0
	for i := 0; i < len(v1); i++ {
		sumSq += math.Pow(float64(v1[i]-v2[i]), 2)
	}
	return math.Pow(sumSq, 0.5)
}
