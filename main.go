package main

import "fmt"

func main() {
	fmt.Println("Enter Your Name and Age: ")
	var name string
	var age int
	fmt.Scanf("%s %d", &name, &age)
	fmt.Printf("Your Name is: %s and \nAge: %d", name, age)

}
