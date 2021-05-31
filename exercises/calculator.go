package calculator

import "fmt"

func sum(x int, y int) int {
	return x + y
}

func substract(x int, y int) int {
	return x - y
}

func multiplication(x int, y int) int {
	return x * y
}

func division(x int, y int) float64 {
	return float64(x) / float64(y)
}

func example() {
	fmt.Println("Sum result: ", sum(10, 4))
	fmt.Println("Subtraction result: ", substract(10, 4))
	fmt.Println("Multiplication result: ", multiplication(10, 4))
	fmt.Println("Division result: ", division(10, 4))
}
