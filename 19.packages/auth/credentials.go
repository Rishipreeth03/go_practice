package auth

import "fmt"

//If you write first letter of the function name as capital 
//then only it is global scope
//if written in small letter is local scope
func LoginWithPassword(username string,password string){
	fmt.Println("logn user ",username,password)
}