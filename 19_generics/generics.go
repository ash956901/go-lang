package main

import "fmt"

// generics in structs
// type stack[T any] struct {
// 	elements []T
// }

// func printSlice(items []int) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// func printStringSlice(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// Generics --> no duplication
// not good practise to use any --> since anything can be passed
// any or interface{} --> same thing
// func printSlice[T int | string | bool | float32](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// we can make this more generic  with comparable keywoard
func printSlice[T comparable, V string](items []T, name V) {
	for _, item := range items {
		fmt.Println(item, name)
	}
}

// lots of duplication for diff types
// we are duplicating same code again and again
// violates DRY-> dont repeat urself
// use generics instead
func main() {
	// myStack := stack[int]{
	// 	elements: []int{1, 2, 3},
	// }
	// myStringStack := stack[string]{
	// 	elements: []string{"abc", "gef"},
	// }
	// fmt.Println(myStack)
	// fmt.Println(myStringStack)
	// nums := []int{1, 2, 3}
	// names := []string{"golang", "typescript"}
	vals := []bool{true, false, true}
	// printSlice(nums)
	// printSlice(names)
	printSlice(vals, "john")
}
