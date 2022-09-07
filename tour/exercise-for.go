package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	z := x
	err := 0.0000001
	count := 10
	for i := 0; i < count; i++ {
		if math.Abs(z * z - x) > err {
			fmt.Println("continue: ", i, "z: ", z)
			z -= (z*z - x) / (2*z)
		} else {
			return z
		}
	}
	return z
}

func main()  {
	x := 101.0
	z1 := sqrt(x)
	z2 := math.Sqrt(x)
	fmt.Println(x, "sqrt is:", z1, "true is:", z2)
}