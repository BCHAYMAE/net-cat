
package main

import (
	"fmt"
	"net" //package provides networking functionality in Go
	"os"
)

func main() {
	port := "8989"

	if len(os.Args) == 2 {
		port = os.Args[1]
	} else if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}

	listener, err := net.Listen("tcp", ":" + port) //create a TCP on the selected port
	//now we check if the listener was created successfully
	if err != nil{
		fmt.Println("error", err)
		return
	}
	defer listener.Close() //close the listener when the function returns
	fmt.Println("Listening on port", port) 

	conn, err :=listener.Accept()
	if err != nil {
		fmt.Println("error", err)
		return
	}
	defer conn.Close()
	fmt.Println("client conn: ", conn.RemoteAddr()) //gives you the client's network address

}