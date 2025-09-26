package main

import "fmt"

func main(){
	//slices are dynamic
	//most used contruct in go 
	//has useful methods related to arrays


	//uninitialised slice is nil
	var nums []int
	fmt.Println(nums)
	fmt.Println(nums==nil)

	var nums1=make([]int,2)
	//here intial size is 2 , we can add more elements also
	fmt.Println(nums1)
	//capacity:maximum number of elements can fit
	fmt.Println(cap(nums1))

	//if you already know the capacity do this 
	var nums2=make([]int,2,5)
	fmt.Println(cap(nums2))
	nums2=append(nums2,1)
	fmt.Println(nums2)
	fmt.Println(cap(nums2))
	fmt.Println(len(nums2))
	//If we go on adding elements , now when 5 is reached ,it doubles the capacity of the array

	
}