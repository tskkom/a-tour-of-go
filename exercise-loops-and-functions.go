package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	e := 0.0001
	d := 0.0
	z := 1.0
	for z*z-x > e || x-z*z > e {
		d = (z*z - x) / (2 * z)
		// fmt.Printf("z = %f, d = %f\n", z, d)
		z -= d
	}
	return z
}

func main() {
	fmt.Println(Sqrt(3))

}
