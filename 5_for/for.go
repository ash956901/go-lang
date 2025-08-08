package main

import "fmt"

// for is the only loop in go
// no while , its the only construct for looping
func main() {

	//while loop
	// i:=1
	// for i <=3 {
	// 	fmt.Println(i)
	// 	i=i+1
	// }

	//infinite loop
	// for {
	// 	fmt.Println("1")
	// }

	//classic for loop
	// for j := 0; j <= 3; j++ {
	// 	//break
	// 	//continue
	// 	if j == 2 {
	// 		continue
	// 	}
	// 	fmt.Println(j)
	// }

	//go 1.22 new feature
	//range , it is exclusive like python range

	for i := range 11 {
		fmt.Println(i)
	}

}
