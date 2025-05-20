package main

import "fmt"

func main() {

	//fixed size -- use array -- helps in memory optimizations -- constant time access, fast

	//declaring a fix type of array
	var nums [4]int

	//getting the length
	fmt.Println(len(nums))

	//adding the element in array
	nums[0] = 1

	fmt.Println(nums)

	var boolsele [4]bool
	fmt.Println(boolsele)
	boolsele[3] = true
	fmt.Println(boolsele)

	var stringsele [4]string
	fmt.Println(stringsele)
	stringsele[3] = "true"
	fmt.Println(stringsele)

	//shorthand fix declaration
	numint := [3]int{1, 2, 3}
	fmt.Println(numint)

	//2d Arrays

	nums2d := [2][2]int{{1, 2}, {3, 2}}
	fmt.Println(nums2d)

}
