package main

import "fmt"

var data int = 21    //this variable is declared outside of any packages
const data2 int = 30 //this is a constant variable outside of any package

func main() {
	local_data := 300

	fmt.Println(data)
	fmt.Println(data)
	fmt.Println(local_data)
}
