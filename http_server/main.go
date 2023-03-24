package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	nl, err := net.Listen("tcp", ":8188") //1 to 65535
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1) //1 = exit with error
	}

	conn, err := nl.Accept()
	if err != nil {
		fmt.Println(err.Error())
	}

	bs := make([]byte, 1024)
	n, err := conn.Read(bs)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(n)
	reqStr := string(bs) // convert ByteString to string
	fmt.Println(reqStr)

	body := `<!DOCTYPE html> <html leng="en"> <head> <title>Page Title</title> </head> <body> <h1>My First Heading</h1> </body> </html>`

	resp := "HTTP/1.1 200 OK\r\n" +
		"Content-Length: %d\r\n" +
		"Content-Type: text/html\r\n" +
		"\r\n%s"

	message := fmt.Sprintf(resp, len(body), body)
	fmt.Println(message)
	conn.Write([]byte(message))

}
