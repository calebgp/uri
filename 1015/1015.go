package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2 float64
	var result float64
	fmt.Scanf("%f %f\n %f %f", &x1, &y1, &x2, &y2)
	result = math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	fmt.Printf("%.4f\n", result)
}
