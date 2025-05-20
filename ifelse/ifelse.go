package main

import "fmt"

func main() {

	age := 18

	if age >= 18 {
		fmt.Println("Person is adult")
	} else if age <= 12 {
		fmt.Println("Person is teenager")
	} else {
		fmt.Println("Person is a kid")
	}
}
