package operations

import "github.com/OlyaIvanovs/calculator/display"

func Add(numbers ...int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func Multiply(numbers ...int) int {
	sum := 1

	for _, number := range numbers {
		sum = multiply(sum, number)
	}

	return sum
}

func Subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func add(x, y int) int {
	return x + y
}

func AdditionWithSteps(numbers ...int) {
	display.PrintAddition(add, numbers...)
}

func MultiplicationWithSteps(numbers ...int) {
	display.PrintMultiplication(multiply, numbers...)
}
