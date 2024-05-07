// Example 1 – map dispatcher

package main

import "fmt"

type calculateFunc func(int, int) int

var (
	operations = map[string]calculateFunc{
		"+": add,
		"-": sub,
		"*": mult,
		"/": div,
	}
)

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mult(a, b int) int {
	return a + b
}

func div(a, b int) int {
	if b == 0 {
		panic("divide by zero")
	}
	return a / b
}

func calculateWithMap(a, b int, opString string) int {
	if operation, ok := operations[opString]; ok {
		return operation(a, b)
	}
	panic("operation not supported")
}

func main() {
	fmt.Println("add=", calculateWithMap(1, 2, "+"))
	fmt.Println("sub=", calculateWithMap(1, 2, "-"))
	fmt.Println("mult=", calculateWithMap(1, 2, "*"))
	fmt.Println("div=", calculateWithMap(1, 2, "/"))
}
