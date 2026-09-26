package main

// fmt is package for output/terminal
import "fmt"

// default value is "" string
// it can be used inside/outside function
var x string

func main() {
	// we can reassign it later just with = operator
	x = "My name is"

	// shorthand to create variable, only used inside funcs
	y := "Givi"

	// for regular output
	fmt.Println(x)
	// for formatted output - v for general value, d for number, T for typeof value
	fmt.Printf("%T", x)
	// more flexible string outputs, //! it returns string
	result := fmt.Sprintf("My name is %v", y)
	fmt.Println(result)

}
