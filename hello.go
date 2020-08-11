package main

import (
	"fmt"

	"example.com/user/hello/morestrings"
)

// An animal object (sort of)
type animal struct {
	name string
	age  int
}

// age getter
func (a animal) Age() int {
	return a.age
}

// name getter
func (a animal) Name() string {
	return a.name
}
func main() {
	fmt.Printf(morestrings.ReverseRunes("\nHello Go"))

	myDog := animal{"Fred", 3}

	// Really tedious print of a simple string with a few vars plugged in
	fmt.Printf(fmt.Sprintf("%s is %d\n", myDog.Name(), myDog.Age()))
}
