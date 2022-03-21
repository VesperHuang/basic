package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operator:" + op)
	}
}

func sum(number ...int) int {
	s := 0
	for i := range number {
		s += number[i]
	}
	return s
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("calling function %s with args (%d, %d) \n", opName, a, b)

	return op(a, b)
}

// func swap(a, b *int) {
// 	*b, *a = *a, *b
// }

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	// fmt.Println(apply(pow, 3, 4))

	// fmt.Println(apply(
	// 	func(a, b int) int {
	// 		return int(math.Pow(float64(a), float64(b)))
	// 	}, 3, 4))

	a, b := 3, 4
	// swap(&a, &b)
	a, b = swap(a, b)
	fmt.Println(a, b)
}
