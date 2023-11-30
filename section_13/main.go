/*
package main

import "fmt"

func main() {
	arr := [...]string{ //[...] -means this is an array, "..." - array literal
		"Bisquite cafe",
		"Banana Pudding",
		"Blueberry Pancacke"}

	arrLen := len(arr)

	fmt.Printf("array is %v \t length is %v", arr, arrLen)
	fmt.Printf("type is %T", arr)

}
*/

package main

import "fmt"

func main() {
	slc := []string{
		"Bisquite cafe",
		"Banana Pudding",
		"Blueberry Pancacke"}
	fmt.Printf("\nThis is a slice - %v, and its length - %v\n", slc, len(slc))
	fmt.Println("---------------")

	for i, v := range slc {
		fmt.Printf("value is %v\nindex %v\n", v, i)
		fmt.Println("---------------")
	}

	//variadic parameter (you can append ANY number of elements to a slice via append func)
	slc = append(slc, "Extra1", "Extra2", "Extra3")
	fmt.Println(slc)
	fmt.Println("---------------")

	//slicing
	// [inclusive: exclusive]
	//[:exclusive]
	//[inclusive:]
	slicedSLC := slc[:1]
	fmt.Printf("this is sliced slice - %v\n", slicedSLC)
	fmt.Println("---------------")

	//delete from slice
	xi := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("xi - %v\n", xi)
	xi = append(xi[:4], xi[5:]...) //this deletes item #4 from slice - xi[:4] == {0,1,2}
	// xi[5:]=={5,6,7,8,9} three dots == open the individual slice items
	//as append doesn't accept another slice to be appended, only individual items
	fmt.Printf("New xi == %v\n", xi)
	fmt.Println("---------------")

	//make() function
	xi = make([]int, 0, 10) //creating slice of length of 0  with up to 10 elements
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi)) //capacity of slice we created
	xi = append(xi, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(xi)
	fmt.Println("---------------")

	// multi-dimensional slice
	jb := []string{"James", "Bond", "Martini", "Chocolate"}
	jm := []string{"Jenny", "Moneypenny", "Guiness", "Wolverine Tracks"}
	fmt.Println(jb)
	fmt.Println(jm)
	xxs := [][]string{jb, jm}
	fmt.Println(xxs)
	fmt.Println("---------------")

	//slice internal

	a := []int{0, 1, 2, 3, 4, 5}
	b := make([]int, 6)
	copy(b, a)
	fmt.Printf("This is about slice internals, clice a is %v\n", a)
	fmt.Printf("This is about slice internals, clice b is %v\n", b)
	a[0] = 7
	fmt.Printf("This is about slice internals, updated clice a is %v\n", a)
	fmt.Printf("This is about slice internals, copied clice b is %v\n", b)
}
