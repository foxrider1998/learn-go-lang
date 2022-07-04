package main

import "fmt"

// use fmt for print anything

// this section is about how to declare an variable
var message = `Nama saya "John Wick".  
Salam kenal.
Mari belajar "Golang".`

var left = false
var right = true
var firstName string = "john"

var lastName string

var chicken map[string]int

// end of  about how to declare an variable

// this is how you run code in main function
func main() {

	chicken = map[string]int{}
	chicken["januari"] = 50
	chicken["februari"] = 40

	main2()
	// this is ho you call an function other than main
	fmt.Println("januari", chicken["januari"]) // januari 50
	fmt.Println("mei", chicken["mei"])         // mei 0 ,karena belum di definisikan

}

// end of how you run code in main function

func main2() {

	fmt.Println("\n and this is how you run a function outside main  \n")

}
