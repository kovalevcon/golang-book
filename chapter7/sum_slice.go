package main

import "fmt"

func sumSlice(xs []int) int {
	total := 0
	for _, v := range xs {
		total += v
	}
	return total
}

func main() {
	slice := []int{98, 93, 77, 82, 83}
	fmt.Println(sumSlice(slice)) // 433
}
