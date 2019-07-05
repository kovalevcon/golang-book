package main

import "fmt"

const x = "Hello, my name is Constantine"

func main() {
	fmt.Println(x)

	fmt.Println("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2
	fmt.Println(output)
}
