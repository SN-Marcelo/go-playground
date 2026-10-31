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
}
