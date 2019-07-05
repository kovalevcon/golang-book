package main

import "fmt"

func main() {
	fmt.Println("Enter a degree in fahrenheits: ")
	var fahrenheits float64
	fmt.Scanf("%f", &fahrenheits)
	celsius := (fahrenheits - 32) * 5 / 9
	fmt.Printf("Degree in celsius: %0.2f", celsius)
}
