package main

import "fmt"

// the inner function even after executing
// keeps the outer scope context with itself
// that is closes the variables in the outer scope within itself
// so next time it remembers it
func counter() func() int {
	var count int = 0
	return func() int {
		count += 1
		return count
	}
}

func main() {
	increment := counter()
	fmt.Println(increment())
	fmt.Println(increment())
}
