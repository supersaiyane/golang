package main

import (
	"fmt"
	"maps"
)

func main() {

	m := map[string]int{"price": 40, "phone": 3}

	//validating if value exist
	v, ok := m["price"] //here "OK" is a kind of bool and "v" is kind of variable which returns the value "_" as well instead of v

	fmt.Println(v)
	if ok {
		fmt.Println("all ok")
	} else {
		fmt.Println("not ok")
	}

	//maps package
	m1 := map[string]int{"price": 40, "phone": 3}
	m2 := map[string]int{"price": 40, "phone": 3}

	fmt.Println(maps.Equal(m1, m2))
}
