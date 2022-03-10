package main

import "fmt"

type iAdder func(int) (int, iAdder)

//正統函數的寫法
func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

func main() {
	// a := adder()
	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("0 + 1 + .... + %d = %d\n ", i, a(i))
	// }

	a := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, a = a(i)
		fmt.Printf("0 + 1 + .... + %d = %d\n ", i, s)
	}
}
