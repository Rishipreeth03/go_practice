package main
import (
	"fmt"
	"time"
	//"math/rand"
)

// --------Sending of data
// Goroutine function that continuously reads from numChan
// and processes each received number.
func processNum(numChan chan int){
	for num :=range(numChan){
		fmt.Println("Processing num",num)
		time.Sleep(time.Second)
	}
}

//recieving
func sum(result chan int,num1 int,num2 int){
	numResult:=num1+num2
	result<-numResult
}


func main(){

	result:=make(chan int)
	go sum(result,4,5)
	res:=<-result // this is also blocking operation
	fmt.Println(res)


	/*
		//Sending
		numChan := make(chan int)
		// Start a goroutine that reads and processes numbers.
		go processNum(numChan)

		// Continuously send random numbers to the channel.
		for {
			numChan <- rand.Intn(100) // ✅ Needs "math/rand" import
			time.Sleep(500 * time.Millisecond) // optional: prevent overwhelming the goroutine
		}

		// The below commented code shows an example of a deadlock scenario.
		
			messageChan := make(chan string)
			messageChan <- "ping" //  Blocking since there's no receiver
			msg := <-messageChan
			fmt.Println(msg)
		
	*/
}