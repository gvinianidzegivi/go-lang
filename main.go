package main

import "fmt"

type Person struct {
	name string
	age  uint
	f    func(arg string) string
}

func (p Person) getName() string {
	return p.name
}

func getAge(person Person) uint {
	return person.age
}

func main() {
	p1 := Person{name: "Giga", age: 23}
	p1.f = func(arg string) string {
		return arg
	}
	y := p1.f("GS")
	z := p1.getName()

	// personAge := getAge(p1)
	fmt.Println(y, z)
}
