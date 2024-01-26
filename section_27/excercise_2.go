package main

import "fmt"

type person struct {
	name string
	age  int
	word string
}

func (p *person) speak() {
	fmt.Println("Bla-bla")
	fmt.Println(p.word)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{
		name: "James",
		word: "I no understand Go",
	}

	saySomething(&p)
	//saySomething(p)

}
