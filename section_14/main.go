package main

import "fmt"

func main() {
	arr := [5]int{} //create array that holds 4 values
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5

	for i := 0; i < len(arr); i++ {
		fmt.Printf("The value from array is %v and the type is %T\n", arr[i], arr[i])

	}
	fmt.Println("============")

	slc := []int{42, 43, 44, 45, 46, 47, 48, 48, 50, 51} // composite literal and slice

	for i := 0; i < len(slc); i++ {
		fmt.Printf("The value from a slice is %v and its index is %v and type is %T\n", slc[i], i, i)
	}
	fmt.Println("============")

	//slicing slices
	a := slc[:5]
	b := slc[5:]
	c := slc[2 : len(slc)-2] //cool trick: len(slc)-1 gives you the index value of the last item in the slice
	d := slc[1:6]
	fmt.Println(a, b, c, d)

	fmt.Println("============")

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...) //We add three dots for appending a slice
	fmt.Println("============")

	///Deleting from a slice
	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(slice)
	sliceNew := append(slice[:3], slice[6:len(slice)]...)
	fmt.Println(sliceNew)
	fmt.Println("============")

	//creating a slice of some length of 0 and maximum capacity(50) with MAKE function
	states := make([]string, 0, 50)
	states = append(states, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`,
		` Colorado`, ` Connecticut`, `Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`,
		` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, `Kentucky`, ` Louisiana`, ` Maine`,
		` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, `Mississippi`, ` Missouri`,
		` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`,
		` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`,
		`Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`,
		` Texas`, `Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`,
		` Wisconsin`, ` Wyoming`)
	fmt.Println(states)

	for i := 0; i < len(states); i++ {
		fmt.Printf("The State is %v and its index is %v\n", states[i], i)
	}

	fmt.Printf("The slice length is %v\n", len(states))
	fmt.Printf("And its capacity is %v\n", cap(states))
	fmt.Println("============")

	//creating a slice of slices
	bondJames := []string{"James", "Bond", "Shaken,not stirred"}
	missMonypenny := []string{"Miss", "Moneypenny", "I'm 008"}
	sliceOfSlice := [][]string{bondJames, missMonypenny}
	//iterating through values of each slice that are stored within parent slice
	for i := range sliceOfSlice {
		fmt.Printf("This is a slice stored within another slice - %v\n", sliceOfSlice[i])
		for i2 := range sliceOfSlice[i] {
			fmt.Printf("This is slice #%v and the at index %v, the value is - %v\n", i, i2, sliceOfSlice[i][i2])
		}
	}
}
