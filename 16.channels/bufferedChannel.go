package main
import (
	"fmt"
	"time"
)

func emailSender(emailchan chan string,done chan bool){
	defer func(){
		done<-true
	}()
	for email:=range emailchan{
		fmt.Println("sending email to",email)
		time.Sleep(time.Second)
	}
}
// you can also limit sending or recieving of data ,
// here now emailchan can only recieve ,but cant send for emailChan
//in case of done boolean , can only send ,not recieve
func emailSender(emailchan <-chan string,done chan<- bool){
	defer func(){
		done<-true
	}()

	//this cant be done
	//emailChan<- "hello@program.com"
	for email:=range emailchan{
		fmt.Println("sending email to",email)
		time.Sleep(time.Second)
	}
}

func main(){
	//this is a buffered channel
	//this can send up to 100 data ,after that it will become blocking
	emailchan:=make(chan string,100)
	done:=make(chan bool)

	go emailSender(emailchan,done)
	for i:=0;i<5;i++{
		emailchan<-fmt.Sprintf("%d@gmail.com",i)
	}

	fmt.Println("done sending")
	//closing channel is most important
	close(emailchan)
	<-done


	chan1:=make(chan int)
	chan2:=make(chan string)
	go func(){
		chan1<-10
	}()

	go func(){
		chan2<-"pong"
	}()

	for i:=0;i<2;i++{
		select{
		case chan1val:=<-chan1:
			fmt.Println("recieved data from chan1",chan1val);
		case chan2val:=<-chan2:
			fmt.Println("recieved data from chan2",chan2val);
		}
	}
	// emailchan<-"1@example.com"
	// emailchan<-"2@example.com"
	// fmt.Println(<-emailchan)
	// fmt.Println(<-emailchan)
}