package main

import "fmt"

// slice --> dynamic arrays, size automatically readjusts
// most used construct in go
// + useful methods

func main() {
	//unintialized slice is nil(null in other languages)
	// var nums []int

	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	//type,intialsize,initial capacity
	// var nums = make([]int, 2, 5)
	//not nill
	//capacity->maximum numbers of element can fit , changes
	//since it automatically resizes
	// fmt.Println(nums)
	// fmt.Println(nums == nil)
	// fmt.Println(cap(nums))
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 4)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))
	//capacity gets doubled if more inserted than initial capacity

	// nums := []int{}
	// nums = append(nums, 1)
	// nums = append(nums, 1)
	// nums = append(nums, 1)

	// fmt.Println(nums)
	// nums[0] = -1
	// nums[1] = 5

	// fmt.Println(nums)

	// var nums = make([]int, 0, 5)
	// var nums2 = make([]int, len(nums))
	// nums = append(nums, 2)
	//copy function
	// copy(nums2, nums)
	// fmt.Println(nums, nums2)

	//slice operator
	// var nums = []int{1, 2, 3}
	// fmt.Println(nums[0:2])
	// fmt.Println(nums[:2]) //by default from starting
	// fmt.Println(nums[1:]) //starting one is inclusive , end is exclusive

	//use of slice package (inbuilt)
	// var nums1 = []int{1, 2}
	// var nums2 = []int{1, 2}
	// fmt.Println(slices.Equal(nums1, nums2)) //returns true if value match

	//2d slices
	var nums = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums)
}
