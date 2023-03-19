package main

import "fmt"

func sum(a, b int) (result int) {
	result = a + b
	return
}

func main() {
	var a, b int
	fmt.Printf("Enter Two numbers: ")
	fmt.Scanf("%d %d", &a, &b)
	result := sum(a, b)
	fmt.Printf("Result is : %d", result)
}
