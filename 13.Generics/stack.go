package main
import (
	"fmt"
)

type stack[T any] struct{
	elements []T
}

func main(){
	myStack := stack[string]{
		elements:[]string{"golang"},
	}

	fmt.Println(myStack)

	//nums:=[]int{1,2,3}
	//nums:=[]string{"hello1","hello2","hello3"}
}