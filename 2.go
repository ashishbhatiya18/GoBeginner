package main

import "fmt"
import "math"

func main() {
	const a string = "Hello World"
	fmt.Println("Hello")

	const b = 50000000
	const c = 3e20 / b
	fmt.Println(int64(c))
	fmt.Println(math.Sin(b))
}
