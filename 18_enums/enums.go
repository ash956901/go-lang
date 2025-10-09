package main

import "fmt"

// enums--> enumerated types
// there is no inbuilt enums in go
// we implement in go using const

// We can also make our custom types in go
// for eg
// type myType string
type OrderStatus string

// grouping of const variables
// const (
//
//	Received OrderStatus= iota // predeclared identifier , for untyped integer which keeps on incrementing starting from 0
//	Confirmed  // hover over each to see value of iota  , 1
//	Prepared // 2
//	Delivered // 3
//
// )
// in string instead of iota you have to give individual values
const (
	Received  OrderStatus = "received"
	Confirmed             = "confirmed"
	Prepared              = "prepared"
	Delivered             = "delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status to", status)
}
func main() {
	changeOrderStatus(Prepared)
}
