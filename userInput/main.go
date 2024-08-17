package main

import "fmt"

func main() {
	fmt.Println("Hey, What is your name ")
	var name string

	//it only read until the first whitespace character
	scan, err := fmt.Scan(&name)
	if err != nil {
		return
	}
	println(scan)
	fmt.Printf("Hello , %s!\n", name)
}
