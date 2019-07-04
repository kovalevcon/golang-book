package main

import "fmt"

func main() {
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("1 + 1 =", 1.0+1.0)

	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

	fmt.Println(32132 * 42452)

	var x string = "Hello World"
	fmt.Println(x)

	x = "first "
	fmt.Println(x)
	x += "second"
	fmt.Println(x)

	var y = "first second"
	fmt.Println(x == y)

	z := 5
	fmt.Println(z)

	dogsName := "Max"
	fmt.Println("My dog's name is", dogsName)
}
