package main

import (
	"io"
	"io/ioutil"
	"os"
)

func main() {

	//creating a new file
	file, err := os.Create("./test.txt")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	_, err = io.WriteString(file, "I love learning the Go programming language!")
	if err != nil {
		return
	}
	ReadFile("./test.txt")

}

func ReadFile(file string) {

	dataByte, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	println(string(dataByte))

}
