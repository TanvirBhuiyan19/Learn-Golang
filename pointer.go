package main

import "fmt"

func update(a *int) {
	fmt.Println("A value is: ", a)
	*a = *a + 10
}

func main() {
	var x int
	var y *int

	fmt.Println("X value is: ", x)
	fmt.Println("X memory address is: ", &x)

	fmt.Println("Y value is:", y)
	fmt.Println("Y memory address is: ", &y)

	x = 10 //Assigning
	y = &x //Referencing

	fmt.Println("X value is: ", x) //Accessing
	fmt.Println("Y value is:", y)
	fmt.Println("Y dereference is: ", *y) //Dereferencing

	update(&x)

	fmt.Printf("X value is: ", x)

}
