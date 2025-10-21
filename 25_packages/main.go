package main

import (
	"fmt"

	"github.com/ash956901/go-lang/auth"
	"github.com/ash956901/go-lang/user"
	"github.com/fatih/color"
)

//how to use third party package
// copy the url
// and do: go get url

// in order to make package --> you have to intialize a module
// module is like your project
// go mod init module_name(project_name)
func main() {
	auth.LoginWithCredentials("ashutosh", "hello")
	session := auth.GetSession()
	fmt.Println("Session:", session)
	user := user.User{
		Email: "user@email.com",
		Name:  "John Doe",
	}

	// fmt.Println(user.Name, user.Email)
	color.Cyan(user.Name)
}

// go mod tidy -->fixes a lot of things helps install undownlaoded pakcages also
