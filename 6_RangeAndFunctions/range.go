package main
import(
	"fmt"
)
func main(){
	arr:=[]int{6,7,8}
	//retrieves index and num
	for i,num:=range arr{
		fmt.Println(i,num)
	}

	m:=map[string]int{"name":1,"age":2}
	for k,v:=range m{
		fmt.Println(k,v)
	}

	//unicode code point rune
	//starting byte of rune
	//if less than 255 then it fits in 1 byte
	//greater than 255 then 2 bytes and so on.. will be assigned 
	for i,c:=range "golang"{
		fmt.Println(i,c)
	}
}