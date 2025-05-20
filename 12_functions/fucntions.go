package main

import "fmt"

func main() {

	// result := add(3, 5)

	// fmt.Println(result)

	fmt.Println(getlanguages())

}

// func add(a int, b int) int {
// 	return a + b
// }

//return multiple values

func getlanguages() (string, string, string) { //define the return type in the sequece of value returning

	return "golang", "javascript", "c"

}
