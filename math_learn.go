package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Printf("%.3f", math.Log(math.E))
	var x, y float64
	x = 5000000000000
	y = math.Log(x) / 2
	fmt.Printf(
		"%E\n",
		y,
	)
}
