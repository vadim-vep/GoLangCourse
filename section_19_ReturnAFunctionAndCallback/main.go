//package main
//
//import "fmt"
//
//func main() {
//	x := foo()
//	fmt.Println(x)
//
//	y := bar()() // function()() calls a function with a function inside of it
//	fmt.Println(y)
//}
//
//func foo() int {
//	return 42
//}
//
//func bar() func() int {
//	return func() int {
//		return 43
//	}
//}

package main

import "fmt"

func main() {
	x := doMath(42, 16, add)
	fmt.Println(x)
	y := doMath(42, 16, subtract)
	fmt.Println(y)
}

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}
