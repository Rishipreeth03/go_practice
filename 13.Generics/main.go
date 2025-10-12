package main
import (
	"fmt"
)
func printSlice[T any](items []T){
	for _,item:=range items{
		fmt.Println(item)
	}
}

//or another way
func printSlice[T int | string | bool] (items []T){
	for _,item:=range items{
		fmt.Println(item)
	}
}
//another way 
func printSlice[T interface{}] (items []T){
	for _,item:=range items{
		fmt.Println(item)
	}
}

func main(){
	nums:=[]int{1,4,5,6}
	printSlice(nums)
}