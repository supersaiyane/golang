package main

import "fmt"

//slices are dynamic arrays ---most used concept in GO -- useful methods

func main() {

	// uninitialised slice is nil
	var nums []int

	fmt.Println(nums)
	fmt.Println(nums == nil)
	fmt.Println(len(nums))

	//if you want the slic does not go nil, it takes initial size but we can add elements to it.

	var numsNonil = make([]int, 2)

	fmt.Println(numsNonil)
	fmt.Println(numsNonil == nil)
	fmt.Println(len(numsNonil))

	//cap -> capacity, maximum number of elements can fit, though it can increase dynamically
	fmt.Println(cap(numsNonil))

	var numsCaps = make([]int, 2, 5)

	fmt.Println(numsCaps)
	fmt.Println(numsCaps == nil)
	fmt.Println(len(numsCaps))
	fmt.Println(cap(numsCaps))

	//adding the element

	numsCaps = append(numsCaps, 1)
	fmt.Println(numsCaps)
	fmt.Println(cap(numsCaps))

	//capacity increase dynamically as well

	var numsCapsIncrease = make([]int, 2, 5)

	fmt.Println(numsCapsIncrease)
	fmt.Println(numsCapsIncrease == nil)
	fmt.Println(len(numsCapsIncrease))
	fmt.Println(cap(numsCapsIncrease))

	//we are adding the elements more than the allocated capcity which is 5

	numsCapsIncrease = append(numsCapsIncrease, 1)
	numsCapsIncrease = append(numsCapsIncrease, 1)
	numsCapsIncrease = append(numsCapsIncrease, 1)
	numsCapsIncrease = append(numsCapsIncrease, 1)
	numsCapsIncrease = append(numsCapsIncrease, 1)
	numsCapsIncrease = append(numsCapsIncrease, 1)
	fmt.Println(numsCapsIncrease)
	fmt.Println(cap(numsCapsIncrease)) //dyncamially increase to 10 from 5  -- it follow the algo that doubles the allocated capacity.
	// Go uses an exponential growth strategy for slices.
	// 	< 1024 capacity → doubling
	// ≥ 1024 capacity → growth slows (~25%)
	// it's a common dynamic array resizing strategy used in many systems.

	//this is very common that when slic is created it allocates the zero, like in previous example,
	//var numsCapsIncrease = make([]int, 2, 5)
	//if we want to avoid that situation then we make the allocation from 2 to 0 so that no 0's will be allocated

	var numsNoZeroinitialized = make([]int, 0, 5)
	fmt.Println(numsNoZeroinitialized)

	fmt.Println(numsNoZeroinitialized)
	fmt.Println(numsNoZeroinitialized == nil)

	numsNoZeroinitialized = append(numsNoZeroinitialized, 1)
	numsNoZeroinitialized = append(numsNoZeroinitialized, 2)
	fmt.Println(numsNoZeroinitialized)
	fmt.Println(len(numsNoZeroinitialized))
	fmt.Println(cap(numsNoZeroinitialized))

	//slicses shorthand

	numsShort := []int{}
	fmt.Println(numsShort)
	numsShort = append(numsShort, 3)
	numsShort = append(numsShort, 4)
	fmt.Println(numsShort)
	fmt.Println(len(numsShort))
	fmt.Println(cap(numsShort))

	// adding the element at index, here if we see the length of the array is 0,
	// and if we try to add the element at 0the place or a 1st place it will throw error as it is initialized as 0
	//var numsIndex = make([]int, 0, 5)
	var numsIndex = make([]int, 2, 5)

	fmt.Println(numsIndex)

	fmt.Println(numsIndex)
	fmt.Println(numsIndex == nil)

	numsIndex[0] = 3
	fmt.Println(numsIndex)
	fmt.Println(len(numsIndex))
	fmt.Println(cap(numsIndex))

	//copy fucntion, copy from one slic to another
	var numsCopy = make([]int, 0, 5)

	numsCopy = append(numsCopy, 3)

	var numsCopy2 = make([]int, len(numsCopy))

	copy(numsCopy2, numsCopy)

	fmt.Println(numsCopy, numsCopy2)

}
