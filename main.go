package main

import "fmt"

func foo(num1, num2 int) (x int, y string) {
	x = num1 + num2
	y = "ok"
	return
}

func main() {
	twoNumberSum, str := foo(4, 45)
	fmt.Println(twoNumberSum, str)
}
