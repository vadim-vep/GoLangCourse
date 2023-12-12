/*
// func (r receiver) identifier (p parameters(s)) (return(s)) {<our code here>}
package main

import "fmt"

func main() {

	// function without arguments and no returns
	foo()

	// function with one parameter s of type string
	bar("Vadim")

	// function with one parameter S type string and a return of type string
	v := aloha("Bond")
	fmt.Println(v)

	// function with two parameters and two returns
	dogName, dogAge := dogEars("Tobias", 12)
	fmt.Println(dogName, dogAge)

	// function with variadic parameters
	result := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("The sum is", result)

	//and we can also pass a slice to function with variadic parameters.
	//note we need to 'unfurl' the slice -> "xi..."
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result2 := sum(xi...)
	fmt.Println("Sum of items from a slice passed is", result2)

}

// function without arguments and no returns
func foo() {
	fmt.Println("I am from foo")
}

// function with one parameter s of type string
func bar(s string) {
	fmt.Println("My name is", s)
}

// function with one parameter S type string and a return of type string
func aloha(s string) string {
	return fmt.Sprint("My name is ", s) //configures return to a type string
}

// function with two parameters and two returns
func dogEars(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(name, " is this old in dog years - "), age
}

// function with variadic parameters
func sum(ii ...int) int {
	fmt.Println(ii)
	fmt.Printf("%T\n", ii)
	s := 0
	for _, v := range ii {
		s += v
	}
	return s
}
*/

/*
// Defer statement
package main

import "fmt"

func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("Foo")
}

func bar() {
	fmt.Println("Bar")
}
*/

/* adding methods
package main

import "fmt"

type person struct {
	first string
}

// function bellow accepts variable p of type person, then prints the result
func (p person) speak() {
	fmt.Println("I am", p.first)
}

func main() {
	p1 := person{
		first: "Vadim",
	}
	p1.speak()
}
*/

/*
// polymorphism
package main

import "fmt"

type person struct {
	first string
}

type secretAgent struct {
	person
	licenseToKill bool
}

// function bellow accepts variable p of type person, then prints the result
func (p person) speak() {
	fmt.Println("I am", p.first)
}

func (sa secretAgent) speak() {
	fmt.Println("I am a secret agent", sa.first)
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "Vadim",
		},
		licenseToKill: true,
	}

	p2 := person{
		first: "Jenny"}

	sa1.speak()
	p2.speak()
}
*/

/*
//interface

package main

import "fmt"

type person struct {
	first string
}

type secretAgent struct {
	person
	licenseToKill bool
}

// function bellow accepts variable p of type person, then prints the result
func (p person) speak() {
	fmt.Println("I am", p.first)
}

func (sa secretAgent) speak() {
	fmt.Println("I am a secret agent", sa.first)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "Vadim",
		},
		licenseToKill: true,
	}

	p2 := person{
		first: "Jenny"}

	sa1.speak()
	p2.speak()

	saySomething(sa1)
	saySomething(p2)
}
*/

/*
//stringer interface
package main

import (
	"fmt"
	"strconv"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("The title of the book is ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("This is the number ", strconv.Itoa(int(c)))
}

func main() {
	b := book{
		title: "West with the Night",
	}
	var n count = 42

	fmt.Println(b)
	fmt.Println(n)
}
*/

/*
package main

import (
	"fmt"
	"log"
	"strconv"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("The title of the book is ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("This is the number ", strconv.Itoa(int(c)))
}

func logInfo(s fmt.Stringer) {
	log.Println("Log From section 19", s.String())
}

func main() {
	b := book{
		title: "West with the Night",
	}
	var n count = 42

	logInfo(b)
	logInfo(n)
}
*/

// Writer type
package main

import (
	"bytes"
	"fmt"
	"io"
)

type person struct {
	first string
}

//
//func (p person) writeOut(w io.Writer) error {
//	_, err := w.Write([]byte(p.first))
//	return err
//}
//func main() {
//	f, err := os.Create("output.txt")
//	if err != nil {
//		log.Fatalf("error %s", err)
//	}
//	defer f.Close()
//
//	s := []byte("Hello gophers!")
//
//	_, err = f.Write(s)
//	if err != nil {
//		log.Fatalf("error %s", err)
//	}
//}

func (p person) writeOut(w io.Writer) error {
	_, err := w.Write([]byte(p.first))
	return err
}

func main() {
	b := bytes.NewBufferString("Hello ")
	fmt.Println(b.String())
	b.WriteString("Gophers! ")
	fmt.Println(b.String())
	b.Reset()
	b.WriteString("It's Thursday ")
	fmt.Println(b.String())

	b.Write([]byte("Happy Happy"))
	fmt.Println(b.String())
}
