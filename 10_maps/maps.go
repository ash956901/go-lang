package main

import (
	"fmt"
	"maps"
)

// maps -> associative data structure
// hash,object,dictionary (cpp,java,python)
// similar to maps in cpp
func main() {
	//creating map
	// m := make(map[string]string)

	//setting an element
	// m["name"] = "golang"
	// m["area"] = "backend"
	//get an element
	// fmt.Println(m["name"], m["area"])
	// IMP: if key does not exists in the map , it returns zero value
	// fmt.Println(m["phone"])

	// m := make(map[string]int)
	// m["age"] = 30
	// m["price"] = 50
	// fmt.Println(m["age"], m["phone"])

	// fmt.Println(len(m))

	//removing an element
	// delete(m,"price")

	//clearing a map -> removing all the maps
	// clear(m)
	// fmt.Println(m)

	//declare map without make keyword
	// m := map[string]int{
	// 	"price": 32,
	// 	"age":   30,
	// }

	// fmt.Println(m)

	//check if element is in the map
	//value,bool checkpresent
	// v, ok := m["price"]
	// fmt.Println(v)
	// if ok {
	// 	fmt.Println("all ok")
	// } else {
	// 	fmt.Println("not ok")
	// }

	//Checking equality of maps
	m1 := map[string]int{"price": 40, "age": 30}
	m2 := map[string]int{"price": 40, "age": 30}

	//invalid equal operation == , use maps package instead , since object comparison just like slices
	fmt.Println(maps.Equal(m1, m2))

}
