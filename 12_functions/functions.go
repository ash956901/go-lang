package main

// function to add two integer and return one integer
// can also write like this for same type arguments
// func add(a,b int) int { so on ... }
// func add(a int, b int) int {
// 	return a + b
// }

// function in go can return many values
// usually used in go to return two values (function logic return value, error)
func getLanguages() (string, string, bool, string) {
	return "golang", "javascript", true, "c"
}

// Functions are first class citizens in go:
// they can be assigned to a var, passed to a func,
// made another function out of it
// func processIt(fn func(a int) int) {
// 	fn(2)
// 	fmt.Println("doneit ")
// }

// to return a function
func processIt() func(a int) int {
	return func(a int) int {
		return 4
	}
}

// entrypoint (called itself by go)
func main() {
	// result:=add(2,5);
	// fmt.Println(result)

	// lang1,lang2,boolval,lang4:=getLanguages()
	// fmt.Println(lang1,lang2,boolval,lang4)

	// fn := func(a int) int {
	// 	return a
	// }
	// processIt(fn)

	fn := processIt()
	fn(6)
}
