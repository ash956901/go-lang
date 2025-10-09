package main

import (
	"fmt"
	"time"
)

// order struct
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // preicison --> nanosecond
	customer            //embedding struct
}

// customer struct --> for embdding , we can acheive compostion and kind of inhertiance , in it .. 
type customer struct {
	name  string
	phone string
}

// we dont have constructors in go (but we do have hack)
// func newOrder(id string,amount float32,status string) *order{
// 	//initial setup goes here.....
// 	myOrder := order{
// 		id: id,
// 		amount: amount,
// 		status: status,
// 		createdAt: time.Now(),
// 	}

// 	//convention->return pointer of the struct
// 	return  &myOrder
// }

// how to attach functions/behaviours to structs like classes (give first letter of struct in that bracket)
//	reciever type --> attaches function to struct
// func (o *order) changeStatus(status string) {
// 	// struct automatically derefences no need of star
// 	o.status = status
// }

// pointer not used here since we are not changing value
// func (o order) getAmount() float32 {
// 	// struct automatically derefences no need of star
// 	return o.amount
// }

// structs are custom data structure
// like classes , we rather use structs
// in go and make an instance of it (like objects)
// imp for oops in go
func main() {
	fmt.Println("Structs")
	// cust := customer{
	// 	name:  "ashutosh",
	// 	phone: "123456",
	// }
	newOrder := order{
		id:        "123",
		status:    "paid",
		amount:    321,
		createdAt: time.Now(),
		customer: customer{
			name:  "ashutosh",
			phone: "123456",
		},
	}
	newOrder.customer.name="robin"
	fmt.Println(newOrder)

	// If you dont set any field, default value is zero value (which we discussed earlier)
	// int-->0,float->0,string->"",bool->false
	// myOrder := order{
	// 	id: "123",
	// 	// amount: 22.45,
	// 	status: "received",
	// }

	//like objects , to define configuration , only one use of struct,inline struct
	// language:=struct{
	// 	name string
	// 	isGood bool
	// } {"ashutosh",true}
	// fmt.Println(language)
	// myOrder:=newOrder("123",50.00,"recieved")
	// myOrder.changeStatus("paid")
	// fmt.Println(myOrder)
	// fmt.Println(myOrder.getAmount())

	// myOrder2 := order{
	// 	id:        "2",
	// 	amount:    100,
	// 	status:    "delivered",
	// 	createdAt: time.Now(),
	// }
	// myOrder.createdAt = time.Now()

	// myOrder.status = "paid"

	// fmt.Println(myOrder.status)
	// fmt.Println(myOrder)
	// fmt.Println(myOrder2)
}
