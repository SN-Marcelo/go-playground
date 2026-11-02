package main

import (
	"fmt"
	"test/math"
	"time"
)

const Pi = 3.14

type Celsius float64
type Farenheit float64

func (c Celsius) ToFarenheit() float64 {
	return float64(c * 9 / 5 * 32)
}

type Person struct {
	Name string
	Age  int
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func Salute() {
	fmt.Println("Hello from goroutine")
}

type Hello interface {
	Bark() string
}

type Dog struct{}

func (d Dog) Bark() string {
	return "Wow!"
}

func printAnything(anything interface{}) {
	fmt.Println(anything)
}

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

	day := 2
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
		fallthrough //Executes the next case
	default:
		fmt.Println("Other day")
	}

	m := map[string]int{
		"a": 1,
		"b": 2,
	}

	for key, value := range m {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	//defer waits till the function finishes
	//It uses a Lifo sequence
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Loop finished")

	temp := Celsius(25)
	fmt.Println(temp.ToFarenheit())

	person := Person{
		Name: "Marcelo",
		Age:  22,
	}

	fmt.Println(person)
	fmt.Printf("Person name: %s\n", person.Name)

	rectangle := Rectangle{
		Width:  5,
		Height: 3,
	}
	fmt.Println(rectangle.Area())

	go Salute() //Go Routine / Background thread
	time.Sleep(time.Second)
	fmt.Println("main function completed")

	ch := make(chan int)

	go func() {
		ch <- 42
	}()
	val := <-ch //waits
	fmt.Println(val)

	ch2 := make(chan string, 2)
	ch2 <- "Hello"
	ch2 <- "World"
	fmt.Println(<-ch2) //Hello
	fmt.Println(<-ch2) //World

	ch10 := make(chan string)
	ch20 := make(chan string)

	go func() { ch10 <- "Channel 10" }()
	go func() { ch20 <- "Channel 20" }()

	select {
	case msg1 := <-ch10:
		fmt.Println(msg1)
	case msg2 := <-ch20:
		fmt.Println(msg2)
	}

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Done"
	}()

	select {
	case msg := <-ch2:
		fmt.Println(msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout!")
	}

	mappi := map[string]int{"manzana": 2, "banana": 3}
	fmt.Println(mappi["manzana"])

	for key, value := range mappi {
		fmt.Printf("key: %s\nvalue: %d\n", key, value)
	}

	var dog Hello = Dog{}
	fmt.Println(dog.Bark())

	printAnything(40)
	printAnything("Hello World")
	printAnything(3.14)

	fmt.Println("Before goto")
	goto End
	fmt.Println("Won't be shown")
End:
	fmt.Println("After Goto")
}
