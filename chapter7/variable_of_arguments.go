package main

import "fmt"

func maxInList(args ...int) int {
	max := args[0]
	for _, v := range args {
		if v > max {
			max = v
		}
	}

	return max
}

func main() {
	fmt.Println(maxInList(1, 4, 5, 0, -11))     // 5
	fmt.Println(maxInList(42, 99, 51, 0, -652)) // 99
}
