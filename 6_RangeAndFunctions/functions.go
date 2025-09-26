package main
import(
	"fmt"
)

func add(a int,b int)int{
	return a+b
}

func getLanguages()(string,string,string){
	return "js","java","c"
}

func processIt(fn func(a int) int)(int){
	return fn(1)
}

func mulFunc() func(a int)int{
	return func(a int)int {
		return 4
	}
}
func main(){
	fmt.Println(add(3,5))
	fmt.Println(getLanguages())
	lang1,lang2,lang3:=getLanguages()
	fmt.Println(lang1,lang2,lang3)

	fn:=func(a int)int{
		return 2
	}
	fmt.Println(processIt(fn))
}