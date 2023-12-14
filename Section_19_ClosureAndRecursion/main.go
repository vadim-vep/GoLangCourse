//package main
//
//import "fmt"
//
//func main() {
//	f := increment()
//	fmt.Println(f()) //1
//	fmt.Println(f()) //2
//	fmt.Println(f()) //3
//	fmt.Println(f()) //4
//	fmt.Println(f()) //5
//
//}
//
//// the bellow is closure because increment function encloses another function
//func increment() func() int {
//	x := 0
//	return func() int {
//		x++
//		return x
//	}
//}

package main

import "fmt"

func main() {
	fact := factorial(4)
	fmt.Println(fact)
}

func factorial(n int) int {

	fmt.Println("This is n ", n)
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
