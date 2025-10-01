package main

import "fmt"

// range ->iterating over data structures
func main() {
	//iterating over a slice using range
	// nums := []int{6, 7, 8}
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(i)
	// }

	// sum := 0
	// for _, num := range nums {
	// 	sum = sum + num

	// }
	// fmt.Println(sum)

	// for i, num := range nums {
	// 	fmt.Println(i,"-->",num)
	// }

	//iterating over a map using range

	// m:=map[string]string{"fname":"john","lname":"doe"}
	// for k,v := range m{
	// 	fmt.Println(k,"-->",v)
	// }

	//starting byte of rune ,unicode code point rune
	// ascii  encoding --> 255 char
	// unicode goes beyond 255 char
	// <=255 -> fits inside 1 byte in memory
	// >255 --> 2 byte or more (in this case our index is the starting byte)
	// 0 1 , 2
	// for i, c := range "golang" {
	// 	fmt.Println(i, c)
	// }

	// to print the char instead of the unicode
	for i, c := range "golang" {
		fmt.Println(i, string(c))
	}

}
