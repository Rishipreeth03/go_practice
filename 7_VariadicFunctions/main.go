package main
import(
	"fmt"
)

func sum(nums ...int) int {
	total:=0

	for _,num :=range nums{
		total=total + num
	}
	return total
}

func main(){
	//you can pass n number of parameters
	fmt.Println(1,4,5,86,"hello")
	nums:=[]int{3,4,5,6}
	fmt.Println(sum(nums ...))
}