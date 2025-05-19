package main

import "fmt"

func main() {

	//if declared have to use

	var name string = "golang"

	// golang automatically pick up the type
	var secondname = "golang2"
	var isAdult = true
	var age int = 30

	//shorthand syntax

	name1 := "gurpreet"

	//we can declare and use the variables later as well , here we need to use the type

	var name3 string

	name3 = "golang3"

	fmt.Println(name)
	fmt.Println(secondname)
	fmt.Println(isAdult)
	fmt.Println(age)
	fmt.Println(name1)
	fmt.Println(name3)

}
