package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// dir, err := os.Getwd()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(dir)

	var fileName string
	fmt.Printf("Enter File Name : ")
	fmt.Scanf("%s \n", &fileName)

	//===================== Single Line Input ===============================//
	// Create a new scanner to read from standard input
	scanner := bufio.NewScanner(os.Stdin)
	// Prompt the user for input
	fmt.Print("Enter a line: ")
	// Read a line of input
	scanner.Scan()
	content := scanner.Text()
	//===================== Single Line Input ===============================//

	isError := createFile(fileName, content)
	fmt.Println(isError)
}

func createFile(fileName, content string) bool {
	posf, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer posf.Close()

	_, err = posf.Write([]byte(content))
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	// fmt.Printf("Number of bytes written: %d\n", n)

	return true
}
