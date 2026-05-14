package main

import (
	"fmt"
	"strings"
)

func factorial(val int) int {
	if val <= 1 {
		return val
	}

	return val * factorial(val-1)
}

func increment() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func add(numbers ...int) int {
	fmt.Println(numbers)

	total := 0
	for _, val := range numbers {
		total += val
	}

	return total
}

func mightPanic(shouldPanic bool) {
	if shouldPanic {
		panic("Something went bad..")
	}

	fmt.Println("This function is executed well..")
}

func recoverable(flag bool) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recoverd from panic.")
		}
	}()

	mightPanic(flag)
}

type MathErr struct {
	Operation string
	Message   string
	InputA    int
	InputB    int
}

const (
	division       = "Division"
	divisionErrMsg = "Divide by 0 is not allowed."
)

func (e *MathErr) Error() string {
	var inputs []string

	if e.Operation == division {
		inputs = append(inputs, fmt.Sprintf("X: %d", e.InputA))
		inputs = append(inputs, fmt.Sprintf("Y: %d", e.InputB))
	}

	return fmt.Sprintf("Math Error in %s => %s, %s", e.Operation, strings.Join(inputs, ","), e.Message)
}

func safeDivision(x, y int) (int, error) {
	if y == 0 {
		return 0, &MathErr{
			Message:   divisionErrMsg,
			Operation: division,
			InputA:    x,
			InputB:    y,
		}
	}

	return x / y, nil
}

func main() {

	val, err := safeDivision(10, 0)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(val)

	// panic("Something went wrong..")

	// recoverable(true)
	// mightPanic(true)

	// fmt.Println(add())

	// counter := increment()

	// fmt.Println(counter())
	// fmt.Println(counter())
	// fmt.Println(counter())

	// logger := func(msg string) {
	// 	fmt.Println(msg)
	// }

	// logger("Hello1")
	// logger("Hello2")
	// logger("Hello3")

}
