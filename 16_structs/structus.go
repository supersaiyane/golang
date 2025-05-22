package main

import (
	"time"
)

//structs are custome data structures like, othe langs has classes
//go doest not have classes, so we use structs

// struct is also a blueprint
type order struct {
	id     string
	amount float32
	status string
	time   time.Time
}

// constructor
func newOrder(id string, amount float32, status string) *order {
	// initial setup goes here
	myorder := order{
		id:     id,
		amount: amount,
		status: status,
	}

	return &myorder

}

//say if we want to change the status or any other variable then e write recevier

// here (o order) --> o is the first letter of the struct and and order is the name, by this
// this fucntion is connected to the recevier fucntion below
func (o *order) changestatus(status string) { //*order as we made a order as pointer, struct automatically takes care of derfrencing
	o.status = status
}

func (o order) getAmount() float32 { // here we have not used * as this recevie is using only GET not changing the variable or value
	return o.amount
}

func main() {

	//if we dont't see any value after print. default value is zero or blank
	//int/float -> 0, string -> empty, bool-> false
	// myorder := order{
	// 	id:     "1",
	// 	amount: 42,
	// 	status: "recevied",
	// }

	// myorder.changestatus("confirmed") //here status will not change as it is not passed as reference, we need to use pointers here, we need to make the recevier fucntion as pointer
	// fmt.Println(myorder.getAmount())
	// myorder.time = time.Now()

	// myorder2 := order{
	// 	id:     "2",
	// 	amount: 43,
	// 	status: "closed",
	// }

	// myorder2.time = time.Now()

	// fmt.Println(myorder)
	// fmt.Println(myorder2)

}
