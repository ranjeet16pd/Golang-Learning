package main

import (
	"Golang-Learning/utils"
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	println("Is golang is really most popular language ")

	lastName := "pd"

	println(utils.Name, lastName)
	ans := utils.Add(2, 4)
	println(ans)

}

// Go Created by google
// It is designed for faster compilation without then need for dependency checking
// It follows Concurrency : Ability of program to execute multiple tasks independently buy not necessarily simultaneously.
// Goroutine is a lightweight thread managed by the Go runtime.
// garbage Collection : In go we do not need to free the memory.
// Go = C + String + garbage Collection + concurrency
