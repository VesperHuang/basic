package main

import (
	"basic/functional/fib"
	"bufio"
	"fmt"
	"os"
)

func trydefer() {
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// fmt.Println(3)

	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		//fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
	}
}

func writeFile(filename string) {
	// file, err := os.Create(filename)
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		// fmt.Println("Error:", err.Error())
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s,%s,%s\n",
				pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibbonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	// trydefer()
	writeFile("fib.txt")
}
