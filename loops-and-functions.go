package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var numOfLoops int = 0
	var z float64 = 1.0
	for {
		nextZ := z - (z*z-x)/(2*z)
		if math.Abs(nextZ-z) < 0.000000000000001 {
			break
		}
		z = nextZ
		numOfLoops++
	}
	fmt.Println("Number of loops: " + fmt.Sprint(numOfLoops))
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
