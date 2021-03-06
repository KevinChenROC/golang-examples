package main

import (
	"fmt"

	"kchen.com/greetings"
	"kchen.com/greetings/hello"
	"kchen.com/math"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Kevin")
	fmt.Println(message)
	fmt.Println(hello.Hi())

	fmt.Println("\nOutput from math module")
	fmt.Println(math.Average([]float64{1, 2, 2.3}))
	fmt.Println(math.Add(1, 2.2))
}
