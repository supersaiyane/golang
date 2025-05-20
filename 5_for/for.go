package main

import "fmt"

//for -> only constrcut in go that is available for looping

func main() {

	//constructing a while loop
	i := 1

	for i <= 3 {
		fmt.Println(i)
		i++
	}

	//infinte loop

	// for {
	// 	fmt.Println(1)
	// }

	//classic for loop

	for i = 0; i <= 3; i++ {

		//break -----> this keyword will stop the execution

		//this is the continue variable it will basically will not print the 2
		if i == 2 {
			continue
		}

		fmt.Println(i)
	}

	//range function
	for i := range 11 {
		fmt.Println(i)
	}

}
