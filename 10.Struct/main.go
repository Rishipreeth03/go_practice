package main
import(
	"fmt"
	"time"
)
//order struct 
type order struct{
	id string
	amount float32
	status string
	createdAt time.Time
}

//reciever type
//sending order by address so chnages take affect
//since we are using struct dont need to use * while changing 
//it is done by default
func (o *order) changeStatus(status string){
	o.status=status
}
func (o order) getAmount() float32{
	return o.amount
}

//constructor of struct order
func newOrder(id string,amount float32,status string) *order{
	myOrder:=order{
		id: id,
		amount:amount,
		status:status,
	}
	return &myOrder
}

func main(){
	myOrder := order{
		id: "1",
		amount: 900.00,
		status: "recieved",
	}
	myOrder.createdAt=time.Now()
	myOrder.changeStatus("Sending")
	fmt.Println("Order struct ",myOrder)


	myOrder1:=newOrder("2",30.24,"recieved")
	fmt.Println("Order 1 using contructor",myOrder1)
}