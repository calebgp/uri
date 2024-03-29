package main

import (
	"fmt"
)

func main() {
	var A, B, C float64
	pi := 3.14159

	fmt.Scan(&A, &B, &C)

	fmt.Printf("TRIANGULO: %.3f\n", (A*C)/2)
	fmt.Printf("CIRCULO: %.3f\n", pi*(C*C))
	fmt.Printf("TRAPEZIO: %.3f\n", (A+B)*C/2)
	fmt.Printf("QUADRADO: %.3f\n", B*B)
	fmt.Printf("RETANGULO: %.3f\n", A*B)
}
