package main

import (
	"fmt"
	"math/rand"
)

/*
func main() {
	x := rand.Intn(251)
	fmt.Printf("the value is %v and it's name is x \n", x)
	if x >= 0 && x <= 100 {
		fmt.Println("Value between 0 and 100")
	} else if x >= 101 && x <= 200 {
		fmt.Println("Value is between 101 and 200")
	} else {
		fmt.Println("Value is between 201 and 250")
	}

}
*/
/*
func init() {
	fmt.Println("This is where the initialization of my my program starts")
}

func main() {

	x := rand.Intn(251)
	fmt.Printf("the value is %v and it's name is x \n", x)

	switch {
	case x <= 100:
		fmt.Println("Value between 0 and 100")
	case x >= 101 && x <= 200:
		fmt.Println("Value is between 101 and 200")
	default:
		fmt.Println("Value is between 201 and 250")
	}
}
*/

/*
func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)
	fmt.Printf("The variables are x=%v, y=%v\n", x, y)

	if x < 4 && y < 4 {
		fmt.Printf("Both x and y are less than 4")

	} else if x > 6 && y > 6 {
		fmt.Printf("Both x and y are greater than 6")

	} else if x >= 4 && x <= 6 {
		fmt.Printf("x is equal or greater than 4 and less or equal to 6")

	} else if y != 5 {
		fmt.Printf("y is not equal to 5")

	} else {
		fmt.Printf("None of the conditions are met")

	}

}
*/

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)
	fmt.Printf("The variables are x=%v, y=%v\n", x, y)

	switch {
	case x < 4 && y < 4:
		fmt.Printf("Both x and y are less than 4")
	case x > 6 && y > 6:
		fmt.Printf("Both x and y are greater than 6")
	case x >= 4 && x <= 6:
		fmt.Printf("x is equal or greater than 4 and less or equal to 6")
	case y != 5:
		fmt.Printf("y is not equal to 5")
	default:
		fmt.Printf("None of the conditions are met")
	}
}
