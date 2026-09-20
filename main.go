package main

import (
	"fmt"
)

func main() {
	var x int = 21
	num := 15
	
	switch num {
	case 12:
		fmt.Println("eq")
	case 15:
		fmt.Println("not eq")
    default:
		fmt.Println("unknown")		
	}

	fmt.Println(x)
} 