package main

import "fmt"

func main() {

	//declaring variables inside if condition
	if age := 15; age >= 18 {
		fmt.Println("person is adult", age)
	} else if age >= 12 {
		fmt.Println("teenager", age)
	}

	//go does not have ternary operator, we have to use if else

	// age := 18

	// if age >= 18 {
	// 	fmt.Println("Person is adult")
	// } else if age <= 12 {
	// 	fmt.Println("Person is teenager")
	// } else {
	// 	fmt.Println("Person is a kid")
	// }

	//logical operator

	// var role = "admin"
	// var hasPermissions = false

	// //goes inside the block
	// if role == "admin" || hasPermissions {
	// 	fmt.Println("yes")
	// }

	// //goes inside the block
	// if role == "admin" && hasPermissions {
	// 	fmt.Println("yes")
	// }

}
