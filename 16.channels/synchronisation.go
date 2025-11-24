package main
import (
	"fmt"
)

//goroutine synchroniser 
//this type are mainly used for single channels
func task(done chan bool){
	defer func(){
		done<-true
	}
	fmt.Println("Processing...")
}
func main(){
	//this is unbuffered channel 
	//at a time only one value is sent 
	done:=make(chan bool)
	go task(done)
	<-done //blocking
	//this also can be used for making the process wait till completion

}