package main

import (
	"fmt"
	"net/http"

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

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func main() {
	fmt.Printf(morestrings.ReverseRunes("\nHello Go"))

	myDog := animal{"Fred", 3}

	// Really tedious print of a simple string with a few vars plugged in
	fmt.Printf(fmt.Sprintf("%s is %d\n", myDog.Name(), myDog.Age()))

	// Add a web server
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}
