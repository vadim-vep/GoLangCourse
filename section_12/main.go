/*
package main
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
package main
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
package main
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
/*
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
*/

/*
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Printf("i is %v \n", i)
		x := rand.Intn(10)
		y := rand.Intn(10)
		fmt.Printf("The variables are x=%v, y=%v\n", x, y)

		switch {
		case x < 4 && y < 4:
			fmt.Printf("Both x and y are less than 4 \n")
		case x > 6 && y > 6:
			fmt.Printf("Both x and y are greater than 6 \n")
		case x >= 4 && x <= 6:
			fmt.Printf("x is equal or greater than 4 and less or equal to 6 \n")
		case y != 5:
			fmt.Printf("y is not equal to 5 \n")
		default:
			fmt.Printf("None of the conditions are met \n")
		}
	}
}
*/

/*
package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for i := 0; i < 42; i++ {
		x := rand.Intn(5)
		switch {
		case x == 0:
			fmt.Printf("x is equal to %v \n", x)
		case x == 1:
			fmt.Printf("x is equal to %v \n", x)
		case x == 2:
			fmt.Printf("x is equal to %v \n", x)
		case x == 3:
			fmt.Printf("x is equal to %v \n", x)
		default:
			fmt.Printf("x is equal to %v \n", x)
		}
	}
}
*/

/*
package main

import "fmt"

func main() {
	i := 20
	for i >= 0 {
		fmt.Println(i)
		i--
	}
}
*/

/*
package main

import "fmt"

func main() {
	i := 20
	for {
		fmt.Println(i)
		i--
		if i < 0 {
			break
		}
	}
}
*/

package main

import "fmt"

func main() {
	for z := 0; z < 20; z++ {
		if z%2 == 0 {
			continue //skips the iteration of the for cycle if z is EVEN
		}
		fmt.Println(z)
	}
}
