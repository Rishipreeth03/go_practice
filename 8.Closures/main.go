package main
import(
	"fmt"
)

func counter() func() int{
	var cnt int=0
	return func() int{
		cnt+=1
		return cnt
	}
}
func main(){
	increment:=counter()

	fmt.Println(increment())
	fmt.Println(increment())
}