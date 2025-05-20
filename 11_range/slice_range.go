package main

import "fmt"

// iterating over DS
func main() {

	nums := []int{6, 7, 8, 9}

	//for loop
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	//sum of the elements
	sum := 0

	for _, num := range nums { //her "_" is the index, will see in next example
		sum = sum + num
	}

	fmt.Println(sum)

	//printing index as well
	for i, num := range nums {
		fmt.Println(num, i)
	}

}
