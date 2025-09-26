package main
import (
	"fmt"
	"maps"
)

//maps->hash ,object,dict
func main(){
	//creating map
	m:=make(map[string]string)

	//setting an element
	m["name"]="golang"
	fmt.Println(m["name"])

	//If key is not present in the map then it returns zero value 
	//i.e returns empty if return value is string 
	delete(m,"name")
	clear(m)
	
	m1:=map[string]int{"price":40,"age":25}
	k,ok :=m1["name"]
	//returns value and boolean 
	//if value not present then k stores 0
	if ok{
		fmt.Println("All ok",k)
	}else {
		fmt.Println("Something wrong")
	}
	m2:=map[string]int{"price":40,"age":25}
	fmt.Println(maps.Equal(m1,m2))
}