package main

import "fmt"

//func main() {
//	foo("Vadim")
//
//	func() {
//		fmt.Println("This is anonimous function")
//	}()
//}
//
//func foo(s string) {
//	fmt.Println("This is named function", s)
//}

//func (r receiver) identifier(p parameters) (r returns) {code of the function}
// ^ this is named function format above

//func (p parameters) (r returns) {code of the function}
// ^ this is anonymous function format above

// we can assign variables to functions and then run them with variable(funcParameter) expression
func main() {
	y := func() {
		fmt.Println("This is anonimous function")
	}
	x := func(s string) {
		fmt.Println("This is named function", s)
	}

	x("Vadim")
	y()
}
