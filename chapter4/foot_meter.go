package main

import "fmt"

func main() {
	fmt.Println("Enter a foots: ")
	var foots float64
	fmt.Scanf("%f", &foots)
	meters := foots * 0.3048
	fmt.Printf("Meters: %0.2f", meters)
}
