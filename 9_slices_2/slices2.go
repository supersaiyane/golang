package main

import (
	"fmt"
	"slices"
)

func main() {

	//slice operators range
	var nums = []int{8, 9, 6, 4, 5, 2}

	// 	slice[start:end]
	// start is inclusive
	// end is exclusive
	// If start is omitted → defaults to 0
	// If end is omitted → goes till the end of slice

	//here [0:2] will return all the elements starting from 0 till the 2nd element , exclusing 2nd
	//in [0:2]  0 -> is from index and 2 -> is to index
	//array goes by 0 , 1 , 2  so it will return 0 and 1, exluding 2nd position
	// so, answer would be 8 and 9
	fmt.Println(nums[0:2])
	fmt.Println(nums[0:1])
	//same can be writtern as
	fmt.Println(nums[:2])

	//** we have to note here in [Index:Index2]  Index -> is always included but Index 2 is omited
	// here it will include the 1st index and put all the index values after that,till end as we dont have Index2 here to say exclude til there
	fmt.Println(nums[1:])

	//here printting willstart from index 1 and goes till index 2, excluding index 3
	fmt.Println(nums[1:3])

	//slice package
	var nums1 = []int{8, 9, 6, 4, 5, 2}
	var nums2 = []int{8, 9, 6, 4, 5, 2}

	fmt.Println(slices.Compare(nums1, nums2))

	//2d slices
	var nums3 = [][]int{{1, 2, 3}, {1, 2, 3}}
	fmt.Println(nums3)

	// nums[0:2]	[8 9]		From index 0 to before 2
	// nums[:2]		[8 9]		Same as above
	// nums[0:1]	[8]			Only index 0
	// nums[1:]		[9 6 4 5 2]	From index 1 to end
	// nums[1:3]	[9 6]		From index 1 to before 3
}
