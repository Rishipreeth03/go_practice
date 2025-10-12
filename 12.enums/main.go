package main
import (
	"fmt"
)

type OrderStatus string

const(
	Recieved OrderStatus="Recieved"
	Confirmed OrderStatus="Confirmed"
	Prepared OrderStatus="Prepared"
	Delivered OrderStatus="Delivered"
)

func changeOrderStatus(status OrderStatus){
	fmt.Println("chnaging order status to",status)
}
func main(){
	changeOrderStatus(Delivered)
	changeOrderStatus(Recieved)
}