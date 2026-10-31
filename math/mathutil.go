package math

//Upper Case to allow usage from other modules
func Add(x, y int) int {
	return x + y
}

//Lower Case to forbid usage from other modules
func substract(x, y int) int {
	return x - y
}

func Divide(x, y int) int {
	if y <= 0 {
		return 0
	}
	return x / y
}
