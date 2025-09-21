package main

import "fmt"

func main() {

	const value int=20

	const(
		port=5000
		host="local"
	)
	age := 29
	if age < 18 {
		fmt.Println("He is small")
	}else if age <20 {
		fmt.Println("He is in 18-20")
	}else{
		fmt.Println("He is old")
	}

	// general for loop
	for i:=0 ; i<3 ;i++ {
		fmt.Println("i is",i)
	}
	//infinite loop 
	for{
		fmt.Println("Hello- Press Ctrl+c to exit");
	}
	//while loop
	i:=3
	for i<=3{
		fmt.Println(i)
		i=i+1
	}
}
