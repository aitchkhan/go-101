package Calculator

type callback func(int, int) int

// Calculate expect a callback function with operation
func Calculate(fn callback, num1, num2 int) int {
	return fn(num1, num2)
}
