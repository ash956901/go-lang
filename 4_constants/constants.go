package main

import "fmt"

//can declare variabels and const outside func
//shorthand outside function is not allowed

const age =30
var diet string="strict"

func main(){
	//constants-->value cant change after we declare it 
	const name string = "golang"
	//name = "ajkldfjkl" not allowed since const 
	fmt.Println(name)
	fmt.Println(diet)
	fmt.Println(age)

	//Constant grouping : to declare multiple constants together
	const (
		host="localhost"
		port=5000
	)

	fmt.Println(host,port)
}