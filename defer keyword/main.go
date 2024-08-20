package main

import "fmt"

func main() {

	// defered line is always executed in the last of the function

	defer println("Prasad")
	fmt.Println("Ranjeet")

	defer println("count 1")
	defer println("count 2")
	defer println("count 3")
	defer println("count 4")

	// It follow the LIFO property (Last defer is executed first)

	println("Lets Learn thorugh loops")
	for i := 0; i < 5; i++ {
		defer fmt.Println("I am defer ", i)
	}

}
