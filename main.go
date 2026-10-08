package main

import (
	"fmt"
	"strings"
)

func main() {

	type Animal struct {
		hasLegs bool
		name    string
	}

	// lion := Animal{
	// 	hasLegs: true,
	// 	name:    "Lion",
	// }

	coffeeBody := "Esp"
	x := len(coffeeBody)

	fmt.Printf("(%T bytes)", x)

	// fmt.Printf("%+v\n", lion)
	s := make([]string, 0, 8)
	// s[0] = "x"

	fmt.Println(s)

	requests := make([]string, 0, 8)
	requests = append(requests, "GET /coffees HTTP/1.1")
	requests = append(requests, "BREW /tea HTTP/1.1")
	requests = append(requests, "GET /beans HTTP/1.1")
	requests = append(requests, "PING /status HTTP/1.1")
	requests = append(requests, "GET /tea HTTP/1.1")

	// fmt.Printf("Logged %d requests (capacity %d)\n", len(requests), cap(requests))

	gets := make([]string, 0)

	for _, line := range requests {
		if strings.HasPrefix(line, "GET") {
			gets = append(gets, line)
		}
	}

	// fmt.Printf("%d of them are GETs\n", len(gets))

	requests = append(requests[:3], requests[4:]...)
	requests = append(requests[:1], requests[2:]...)
	// fmt.Printf("Cleaned log (%d): %v\n", len(requests), requests)

	// TODO: build recent as the last 2 elements of requests (a subslice)
	// recent := requests[len(requests)-2:]
	// TODO: print the most recent request (last element) and the recent window
	// fmt.Println(recent)

	var mp = make(map[string]string)

	mp["name"] = "John"

	for i := range mp {
		fmt.Println(i)
	}

}
