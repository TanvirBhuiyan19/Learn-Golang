package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
	}
	// fmt.Println(dir)

	// var fileName string
	// fmt.Printf("Enter File Name : ")
	// fmt.Scanf("%s \n", &fileName)

	// //===================== Single Line Input ===============================//
	// // Create a new scanner to read from standard input
	// scanner := bufio.NewScanner(os.Stdin)
	// // Prompt the user for input
	// fmt.Print("Enter a line: ")
	// // Read a line of input
	// scanner.Scan()
	// content := scanner.Text()
	// //===================== Single Line Input ===============================//

	// isError := createFile(fileName, content)
	// fmt.Println(isError)

	// fileInfo, err := os.Stat("Tanvir.txt")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println(fileInfo.IsDir())
	// fmt.Println(fileInfo.ModTime().Date())
	// fmt.Println(fileInfo.Name())
	// fmt.Println(fileInfo.Size())

	//How to create a directory
	// err := os.Mkdir("New Folder", 0777)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	path := filepath.Base(dir)
	relativePath := filepath.Join("New Folder")
	absolutePath, err := filepath.Abs("New Folder")
	if err != nil {
		fmt.Println(err.Error())
	}
	newPath := filepath.Join(absolutePath, "..", "..", "New Folder")
	err = os.Mkdir(newPath, 777)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(path)
	fmt.Println(relativePath)
	fmt.Println(absolutePath)
	fmt.Println(newPath)

	//Rename existing folder
	// os.Rename(absolutePath, `New Folder 2`)

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
