package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func printArray(arr []int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printSlice(s []int) {
	fmt.Printf("s=%v,len=%d, cap=%d\n", s, len(s), cap(s))
}

func sliceops() {
	fmt.Println("creating slices")
	var s []int

	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	s2 := make([]int, 16)
	printSlice(s2)

	s3 := make([]int, 10, 32)
	printSlice(s3)

	fmt.Println("Copying slice")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("Deleting elements from slice")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)

	fmt.Println("Popping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail)
	printSlice(s2)
}

func main() {

	sliceops()

	// arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	// fmt.Println("arr[2:6] =", arr[2:6])
	// fmt.Println("arr[:6] =", arr[:6])
	// fmt.Println("arr[2:] =", arr[2:])
	// fmt.Println("arr[:] =", arr[:])

	// s1 := arr[2:]
	// fmt.Println("s1 =", s1)
	// s2 := arr[:]
	// fmt.Println("s2 =", s2)

	// fmt.Println("After updateSlice(s1)")
	// updateSlice(s1)
	// fmt.Println(s1)
	// fmt.Println(arr)

	// fmt.Println("After updateSlice(s2)")
	// updateSlice(s2)
	// fmt.Println(s2)
	// fmt.Println(arr)

	// fmt.Println("Reslice")
	// fmt.Println(s2)
	// s2 = s2[:5]
	// fmt.Println(s2)
	// s2 = s2[2:]
	// fmt.Println(s2)

	// fmt.Println("Extending slice")
	// arr[0], arr[2] = 0, 2
	// fmt.Println("arr =", arr)
	// s1 = arr[2:6]
	// s2 = s1[3:5]
	// fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
	// fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))

	// s3 := append(s2, 10)
	// s4 := append(s3, 11)
	// s5 := append(s4, 12)
	// fmt.Println("s3, s4, s5 = ", s3, s4, s5)
	// fmt.Println("arr = ", arr)

	// ========================================
	// var arr1 [5]int
	// arr2 := [3]int{1, 3, 5}
	// arr3 := [...]int{2, 4, 6, 8, 10}
	// var grid [4][5]int

	// fmt.Println(arr1, arr2, arr3)
	// fmt.Println(grid)

	// fmt.Println("printArray(arr1)")
	// printArray(arr1[:])

	// fmt.Println("printArray(arr3)")
	// printArray(arr3[:])

	// fmt.Println("arr1 and arr3")
	// fmt.Println(arr1, arr3)
}
