package auth

import "fmt"

// beggining with small letter -> scope within the same package
// beggining with large letter -> outside too 
func LoginWithCredentials(username string,password string){
	fmt.Println("logging with using",username,password)
}

