package main

import "fmt"

var message = `Nama saya "John Wick".
Salam kenal.
Mari belajar "Golang".`

var left = false
var right = true
var firstName string = "john"

var lastName string

func main() {
	lastName = "wick"
	fmt.Println(message)
	fmt.Printf("halo %s %s!\n", firstName, lastName)

	var leftAndRight = left && right
	fmt.Printf("left && right \t(%t) \n", leftAndRight)

	var leftOrRight = left || right
	fmt.Printf("left || right \t(%t) \n", leftOrRight)

	var leftReverse = !left
	fmt.Printf("!left \t\t(%t) \n", leftReverse)
}
