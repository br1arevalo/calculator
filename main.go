package main

import (
	"fmt"
)

func main() {
	n1, n2 := input()
	dec(n1, n2)

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

func dec(n1, n2 float64) {
	var decision int

	fmt.Println("What do you want to do?\nPress 1 for Sum\nPress 2 for subtract\nPress 3 for multiplication\nPress 4 for division")
	fmt.Scan(&decision)
	switch {
	case decision == 1:
		sum(n1, n2)
	case decision == 2:
		sub(n1, n2)
	case decision == 3:
		mul(n1, n2)
	case decision == 4:
		div(n1, n2)
	default:
		fmt.Println("Incorrect number, please try again")
		dec(n1, n2)
	}
}
