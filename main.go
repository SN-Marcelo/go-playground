package main

import (
	"fmt"
	"test/math"
)

const Pi = 3.14

func main() {
	var x int = 5
	var y int = 3
	z := 6 //Declaration only valid in functions
	fmt.Println(math.Add(x, y))
	fmt.Println(z)
	fmt.Println(Pi)
	fmt.Println(math.Divide(x, y))

	if x > 5 {
		fmt.Println("X es mayor a 5")
	} else if x == 5 {
		fmt.Println("X es 5")
	} else {
		fmt.Println("X es menor que 5")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	nums := []int{1, 2, 3}
	for _, num := range nums {
		fmt.Println(num)
	}
}
