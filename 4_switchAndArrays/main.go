package main

import (
	"fmt"
	"time"
)

func main() {
	i := 5
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("other")
	}

	//multiple condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		{
			fmt.Println("Weekend")
		}
	default:
		fmt.Println("Its working day")
	}

	//typeSwitch
	whoAmI := func(i interface{}) {
		switch i := i.(type) {
		case int:
			fmt.Println("Integer")
		case string:
			fmt.Println("String")
		default:
			fmt.Println("Other", i)
		}
	}
	whoAmI("Hello")

	//Arrays
	var nums [4]int

	//array length
	fmt.Println(len(nums))

	nums[0]=1
	fmt.Println(nums)


	var arr [4]bool
	fmt.Println(arr)
	//declaring in single line 
	numarray:=[3]int{1,2,3}
	fmt.Println(numarray)

	//2D arrays
	nums1:=[2][2]int{{3,4},{5,6}}
	fmt.Println(nums1);

	//uses of arrays
	//-fixed size ,that is predictable 
	//-Memory Optimisation
	//-constant time access
}
