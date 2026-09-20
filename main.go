package main

import (
	"fmt"
	// "math"
	"strconv"
)

func main() {
	 num1 := "12"
	 result, err := strconv.Atoi(num1) 
	 y := result + 5
 	// x := -1000
	// y := uint(x)
	// fmt.Println(x, y)
	// var z string
	// z = "Giga"
	// z := uint(0 )
	fmt.Println(err, y)
	fmt.Printf("%T", result)
	// fmt.Printf("value of the number is %.2f", num)
}