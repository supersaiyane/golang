package main

import "fmt"

func main() {
	m := map[string]int{"price": 40, "phone": 3}

	for k, v := range m {
		fmt.Println(k, v)
	}

	for k := range m {
		fmt.Println(k)
	}

	//get unicode point rune
	// here i is the starting byte of the rune
	for i, c := range "golang" {
		fmt.Println(i, c)
		fmt.Println(i, string(c))
	}
}
