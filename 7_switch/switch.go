package main

import (
	"fmt"
	"time"
)

func main() {
	//simple switch (many conditions)
	i := 3
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("other")
	}
	//default is optional,no need to write break here

	//multiple condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("weekend")
	case time.Friday:
		fmt.Println("party day")
	default:
		fmt.Println("Workday")

	}

	//type switch,use for type check
	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("Its an integer")
		case string:
			fmt.Println("its a string")
		case bool:
			fmt.Println("Its a boolean")
		default:
			fmt.Println("other")
		}
	}
	whoAmI(20.1)
	//i interface{} means we can pass anything

}
