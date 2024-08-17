package main

import "fmt"

func addFunction(a int, b int) int {
	return a + b
}

func addFuncutionWithVariable(a, b int) (result int) {
	result = a + b
	return
}

func main() {
	var a, b int
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		return
	}
	println(addFuncutionWithVariable(a, b))
}
