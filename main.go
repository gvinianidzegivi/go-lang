package main

import "fmt"

func main() {
	var favoriteColor = "Black"
	birthYear, age := 1993, 33

	var (
		firstInitial rune = 'G'
		lastInitial  rune = 'G'
	)

	var ageInDays int
	ageInDays = 365 * age

	fmt.Println("Favorite color:", favoriteColor)
	fmt.Println("Born in", birthYear, "- age:", age)
	fmt.Println("Initials:", firstInitial, lastInitial)
	fmt.Println("Age in days:", ageInDays)
}
