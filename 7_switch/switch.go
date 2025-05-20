package main

import "fmt"

func main() {

	//simple switch

	// i :=5

	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// default:
	// 	fmt.Println("other")
	// }

	//multiple swich

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("its weekend")
	// default:
	// 	fmt.Println("itwork")
	// }

	//type switch

	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("it integer")
		case string:
			fmt.Println("it string")
		case bool:
			fmt.Println("it boolean")
		default:
			fmt.Println("other")
		}

	}

	whoAmI(50.36)

}
