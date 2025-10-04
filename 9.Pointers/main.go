package main
import "fmt"

//using pass by value
func changeNum(num int){
	num=5
	fmt.Println("In change ",num)
}
// pass by address
func changeNumPassByAddress(num *int){
	*num=5
	fmt.Println("In change ",*num)
}
func main(){
	num:=1
	fmt.Println("First pass by value then by address")
	changeNum(num)
	fmt.Println("After Change",num)

	changeNumPassByAddress(&num)
	fmt.Println("After Change",num)
}