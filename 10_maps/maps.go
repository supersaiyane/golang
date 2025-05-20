package main

import "fmt"

// maps -> associative ds, like hash, object, dict
func main() {

	//creating map

	m := make(map[string]string) // [string]string --> [string] --> type of key / string -> type of value

	//setting elements

	m["name"] = "golang"
	m["area"] = "backend"

	fmt.Println(m["name"], m["area"])

	//if we try to access a key does not exist, then it return empty if string, 0 if int
	fmt.Println(m["dance"])

	m1 := make(map[string]int) // [string]int --> [string] --> type of key / int -> type of value

	m1["age"] = 30
	m1["phone"] = 30245156

	fmt.Println(m1["age"])
	fmt.Println(m1["phone"])

	//to get lenght
	fmt.Println(len(m1))

	//to delete
	delete(m1, "age")
	fmt.Println(m1)
	fmt.Println(len(m1))

	//to empty map
	clear(m1)

	//create map without make

	m2 := map[string]int{"price": 40, "phone": 3}
	fmt.Println(m2)

}
