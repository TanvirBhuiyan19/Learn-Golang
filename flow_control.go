package main

import "fmt"

func main() {
	/*
		fmt.Printf("Enter age : ")
		var age int
		fmt.Scanf("%d", &age)
	*/
	/*
		if age <= 5 {
			fmt.Printf("Baby")
		} else if age > 5 && age <= 12 {
			fmt.Printf("Child")
		} else if age > 12 && age <= 18 {
			fmt.Printf("Teen age")
		} else {
			fmt.Printf("Adult")
		}
	*/
	/*
		switch age {
		case 1, 2, 3, 4, 5:
			fmt.Printf("Baby\n")
			fallthrough //If this condition will true, it will also execute this and next both case;
		case 6, 7, 8, 9, 10, 11, 12:
			fmt.Printf("Child")
		case 13, 14, 15, 16, 17, 18:
			fmt.Printf("Teen Age")
		default:
			fmt.Printf("Adult")

		}
	*/

	/*
		//For Loop
			for i := 1; i <= 10; i++ {
				fmt.Println(i)
			}
	*/

	/*
		//Foreach Loop
		students := []string{"Tanvir", "Shuvo", "Jahir", "Oshim"}
		for i, student := range students {
			fmt.Println(i, student)
		}
	*/

	//While Loop
	i := 0
	for i < 5 {
		fmt.Printf("%d => Hello Tanvir\n", i+1)
		i++
	}

}
