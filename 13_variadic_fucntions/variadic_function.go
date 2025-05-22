package main

import "fmt"

func main() {

	//println is a variadic fucntion
	result := sum(1, 5, 8, 9, 6, 4, 7, 2, 5, 4, 25)

	//if in a slice

	num := []int{1, 2, 3}

	results := sum(num...)

	fmt.Println(result, results)
}

// here "...int" means unlimited parameters of type int
func sum(nums ...int) int {

	total := 0

	for _, nums := range nums {
		total = total + nums
	}

	return total
}

// anytype
// func anytype(nums ...interface{}) int {
// 	return 1
// }
