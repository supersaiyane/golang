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

	//finally
	var price float32 = 50.1
	var price2 = 50.2
	price3 := 50.3

	fmt.Println(name)
	fmt.Println(secondname)
	fmt.Println(isAdult)
	fmt.Println(age)
	fmt.Println(name1)
	fmt.Println(name3)
	fmt.Println(price)
	fmt.Println(price2)
	fmt.Println(price3)

}
