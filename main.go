package main

import "fmt"

func main() {
	str := "Hello World"

	for i := 0; i < len(str); i++ {
		// fmt.Printf("%c", str[i])
		// fmt.Println(string(str[i]))
	}

	for i, char := range str {
		fmt.Println(i, string(char))
	}
}
