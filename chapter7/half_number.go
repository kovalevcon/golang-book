package main

import "fmt"

func half(x int) (int, bool) {
	if x%2 == 0 {
		return x / 2, true
	}

	return 0, false
}

func main() {
	fmt.Println(half(1)) // 0, false
	fmt.Println(half(2)) // 1, true
	fmt.Println(half(4)) // 2, true
}
