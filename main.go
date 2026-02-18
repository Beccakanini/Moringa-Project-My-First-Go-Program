package main //every Go program starts with this

import "fmt" //used for printing

func main(){  //the main function starts from here
	fmt.Println("What is your name?") //print the question on the cli

	var name string //create a variable to store name
	fmt.Scanln(&name) // Read user input

	fmt.Printf("Hello, %s! Welcome to Go!\n You just run your first Go program",name) //print greeting
}