package main

import (
	"fmt"
)

// receives n number of inputs of int type
// basically takes a slice of int
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total = total + num
	}

	return total
}

// variadic functions --> can pass n number of parameters
func main() {
	// fmt.Println(1,2,3,5,"hello")

	nums:=[]int{3,4,5,6}
	result := sum(nums...)
	fmt.Println(result)

}
