package main

import (
	"fmt"
)

func main() {
	n1, n2 := input()

	sum(n1, n2)
	sub(n1, n2)
	mul(n1, n2)
	div(n1, n2)
}

func sum(n1, n2 float64) {
	fmt.Println("Sum:", n1+n2)
}

func input() (float64, float64) {
	n1, n2 := 0.0, 0.0

	fmt.Println("Enter first number:")
	fmt.Scan(&n1)
	fmt.Println("Enter next number: ")
	fmt.Scan(&n2)
	fmt.Println("Number 1:", n1, "\nNumber 2:", n2)

	return n1, n2
}

func sub(n1, n2 float64) {
	fmt.Println("Subtract:", n1-n2)
}

func mul(n1, n2 float64) {
	fmt.Println("Multiplication:", n1*n2)
}

func div(n1, n2 float64) {
	fmt.Println("Division:", n1/n2)
}
