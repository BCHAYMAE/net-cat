
package main

import (
	"fmt"
	"net" //package provides networking functionality in Go
	"os"
)

func main() {
	port := "8989" //set the default port by choice

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

	for {
		conn, err :=listener.Accept()
		if err != nil {
			fmt.Println("error", err)
			continue
		}

		fmt.Println("client conn: ", conn.RemoteAddr()) //gives you the client's network address

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn){
	defer conn.Close()
	buffer := make([]byte, 1024) 
	//now we need to read the data send by the client into the buffer
	for{
		n, err := conn.Read(buffer)
		//check if reading failed
		if err != nil {
			fmt.Println("error", err)
			break
			//return : exits the entire function
			//break : exits only the current loop
		}
		messg := string(buffer[:n])//convert only the received bytes ino readable text
		// then we display the message received from the client
		fmt.Println("msg :", messg)
	}
}
