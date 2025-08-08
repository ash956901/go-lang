package main

import "fmt"

func main() {
	//var name type , have to use every variable created,else delete it
	// var name string = "golang"
	//type inference automatically,for all types
	var name = "golang"

	var isAdult bool = true

	//many int version , int 32 ,64
	//most of the time we use int, go automatially optimizes it
	//due to its architecture
	var age int = 32

	fmt.Println(name)
	fmt.Println(isAdult)
	fmt.Println(age)

	//Short hand syntax
	naam := "golang"
	fmt.Println(naam)

	//In the cases where value will be given later we will have
	//to give the type as we dont have any vale
	//for eg: var name , if i dont give type it wont be able to infer
	//so var name string is our variable

	var n1 string

	n1 = "naanu"

	fmt.Println(n1)

	//floats
	// var price float32 =50.5
	// var price=50.5
	price := 50.5

	fmt.Println(price)

}
