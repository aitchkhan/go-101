package main

import (
	"101/Calculator"
	"fmt"
)

type callback func(int, int) int

func main() {
	num1 := 2
	num2 := 3
	fmt.Println("this is another package")
	sum := Calculator.Calculate(add, num1, num2)
	fmt.Printf("\n sum of %v and %v is %v", num1, num2, sum)
}

func add(num1, num2 int) int {
	return num1 + num2
}

func subtract(num1, num2 int) int {
	return num1 - num2
}

func multiply(num1, num2 int) int {
	return num1 * num2
}

func divide(num1, num2 int) int {
	return num1 / num2
}
