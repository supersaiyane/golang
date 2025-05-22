package main

import "fmt"

func main() {

	// result := add(3, 5)

	// fmt.Println(result)

	fmt.Println(getlanguages())

	//another way of writing
	// lang1, lang2, lang3 := getlanguages()
	// fmt.Println(lang1, lang2, lang3)

	//sometimes if we do not want to use any parameter then we simply pass "_"
	// lang1, lang2, _ := getlanguages()
	// fmt.Println(lang1, lang2)

}

// func add(a int, b int) int {
// 	return a + b
// }

//return multiple values

func getlanguages() (string, string, bool) { //define the return type in the sequece of value returning

	return "golang", "javascript", true

}
