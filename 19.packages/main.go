package main

import (
	"fmt"
	"pack/auth"
	"pack/user"
	"github.com/fatih/color"
)

func main(){
	auth.LoginWithPassword("rishipreeth","secert")
	session:=auth.GetSession()
	fmt.Println(session)

	User1:=user.User{
		Email: "hello@gmail.com",
		Name: "mypassword",
	}
	// fmt.Println(User1)
	color.Red(User1.Email)
}
