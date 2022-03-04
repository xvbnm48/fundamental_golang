package main

import (
	"fmt"

	"github.com/xvbnm48/fundamental_golang/calculation"
)

func main() {
	fmt.Println("Hello, World!")
	result := calculation.Add(10, 10)
	//sentence := TestAja()
	//fmt.Println(sentence)
	fmt.Println(result)

	fmt.Println("proses perkalian")
	result2 := calculation.ProsesKali(10, 10)
	fmt.Println(result2)

	var nama string = "sakura miyawaki"
	fmt.Println(nama)

	// age := 50
	// if age > 10 {
	// 	fmt.Println("age is greater than 10")
	// } else {
	// 	fmt.Println("age is less than 10")
	// }

	score := 80
	var grade string
	if score <= 50 {
		grade = "E"
	} else if score >= 60 {
		grade = "A"
	} else {
		grade = "C"
	}

	fmt.Println(grade)

	nilai := 3
	switch nilai {
	case 1:
		fmt.Println("nilai 1")
	case 2:
		fmt.Println("nilai 2")
	case 3:
		fmt.Println("nilai 3")
	default:
		fmt.Println("nilai tidak diketahui")
	}
}
