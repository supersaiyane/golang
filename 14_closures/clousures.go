package main

import "fmt"

//A closure in Go is a function value that references variables from outside its body. The function remembers the variables even after the outer function has finished executing.

func counter() func() int {

	count := 0

	return func() int {
		count++

		return count
	}
}

func main() {
	inc := counter()
	fmt.Println(inc()) // 1
	fmt.Println(inc()) // 2
	fmt.Println(inc()) // 3

	// New instance of the counter
	newInc := counter()
	fmt.Println(newInc()) // 1 (separate state)

}
