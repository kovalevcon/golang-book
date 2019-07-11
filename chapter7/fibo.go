package main

import "fmt"

func fibonacci(x int) int {
	if x == 0 || x == 1 {
		return x
	}

	return fibonacci(x-1) + fibonacci(x-2)
}

func main() {
	fmt.Println(fibonacci(3))  // 2
	fmt.Println(fibonacci(7))  // 13
	fmt.Println(fibonacci(10)) // 55
}
