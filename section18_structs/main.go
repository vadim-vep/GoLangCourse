/* // Struct with Slices and maps example
package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	iceCreams []string //creating a slice inside a struct
}

func main() {
	p1 := person{
		firstName: "James",
		lastName:  "Bond",
		iceCreams: []string{"strawberry"}}
	p2 := person{
		"Jenny",
		" Moneypenny",
		[]string{"chocolate", "vanilla"}}

	for _, v := range p1.iceCreams {
		fmt.Printf("These are %v %v's favorite icecreams: %v\n", p1.lastName, p1.firstName, v)
	}

	for _, v := range p2.iceCreams {
		fmt.Printf("And these are %v %v's favorite icecreams: %v\n", p2.firstName, p2.lastName, v)
	}

	//creating a map with keys strings and values of type person
	m := map[string]person{
		p1.lastName: p1, //key == lastName of p1,p2 we created above (string):
		p2.lastName: p2, // value == values of p1,p2 (type person)
	}
	fmt.Println(m[p1.lastName]) //print out values by key p1.lastName
	fmt.Println(m[p2.lastName])

	for _, v := range m {
		fmt.Printf("These are %v's icecreams from the map m: %v\n", v.lastName, v.iceCreams)
		for _, v2 := range v.iceCreams {
			fmt.Printf("Each individual ice cream %v likes is %v\n", v.lastName, v2)
		}
	}

}
*/

/*
//embed struct

package main

import "fmt"

type engine struct {
	electric bool
}
type vehicle struct {
	engine
	make  string
	model string
	doors int64
	color string
}

func main() {
	car1 := vehicle{
		engine: engine{
			true},
		make:  "Ford",
		model: "Mustang",
		doors: 2,
		color: "Black",
	}
	car2 := vehicle{
		engine: engine{
			false},
		make:  "Toyota",
		model: "Prius",
		doors: 4,
		color: "red",
	}

	fmt.Printf("The first car is %v and the second one is%v\n", car1, car2)

	fmt.Printf("Does the cars have electronic engine? First one - %v and second one - %v\n",
		car1.engine.electric, car2.engine.electric)

}
*/

//anonymous structs

package main

import "fmt"

func main() {

	p1 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first:     "Vadym",
		friends:   map[string]int{"Pasisnychenko": 4, "Sanyok": 3, "I'm forever alone": 100},
		favDrinks: []string{"tequila", "beer with vodka", "beer"},
	}

	fmt.Println(p1)
}
