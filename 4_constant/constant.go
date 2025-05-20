package main

import "fmt"

// can be declared outside
const ag2 = 30

//shorthands cannot work outside the function

func main() {

	//cannot be changes once declared
	const name = "golang"
	const age = 30

	//we can group the constans as well
	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(port, host)
	fmt.Println(ag2)
}
