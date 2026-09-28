package main

import "fmt"

func main() {
	// data types

	// for text types
	var str string = "Hello World"

	// bool -> true | false, default value false
	var isActive bool = false
	// float -> float32 | float64
	var floatNum float32 = 1.2

	// positive numbers max 20(0)
	var uNumber uint = 10000000000000000000
	// positive numbers min/max 0-255
	var uNumber8 uint8 = 255

	// positive/negative numbers min/max  -20(0)/20(0) max 18
	var number int = -1000000000000000000
	// positive/negative numbers min/max -128/127
	var number8 int8 = -127

	// byte -> like uint8 type,
	var bt byte = 255

	// rune -> like int32, mostly for symbols/characters returns number in bytecode
	// to convert and see actual value, we can use %c in Prinf method
	var rn rune = '$'

	fmt.Println(
		isActive,
		floatNum,
		bt, uNumber,
		uNumber8, number,
		number8,
		str,
		rn,
	)
}
