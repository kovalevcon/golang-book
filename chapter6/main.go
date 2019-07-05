package main

import "fmt"

func main() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)

	y := [5]float64{98, 93, 77, 82, 83}
	var total float64 = 0
	for _, value := range y {
		total += value
	}
	fmt.Println(total / float64(len(y)))

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	slice3 := make([]int, 2)
	copy(slice3, slice1)
	fmt.Println(slice3)

	z := make(map[string]int)
	z["key"] = 10
	fmt.Println(z["key"])

	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}

	name, ok := elements["O"]
	fmt.Println(name, ok)

	if name, ok := elements["F"]; ok {
		fmt.Println(name, ok)
	}

	c := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(c[2:5])
}
