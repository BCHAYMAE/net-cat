
package main

import (
	"bufio" //helps us read client input
	"fmt"
	"net" //package provides networking functionality in Go
	"os"
	"strings" //helps us clean the username
	"sync" //provides synchronization tools
)

type Client struct{
	name string //client name
	conn net.Conn //client's network connection
}

var clients []Client //slice to store the connected clients
var clientMutex sync.Mutex //mutex to synchronize access to the clients slice

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
	defer removeClient(conn) //remove the client from the slice when the function returns

	//send the username prompt to the client
	conn.Write([]byte("[ENTER YOUR NAME]: "))

	//create a reader to receive input from the client
	reader := bufio.NewReader(conn)

	//read the username entered by the client
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error reading username:", err)
		return
	}

	//remove spaces and the newline from the username
	name = strings.TrimSpace(name)

	//check if the username is empty
	if name == "" {
		fmt.Println("empty username")
		return
	}

	//create a new client with its name and connection
	newClient := Client{
		name: name,
		conn: conn,
	}

	//add the new client to the clients slice
	clientMutex.Lock()
	clients = append(clients, newClient)
	clientMutex.Unlock()

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

		//then we display the message received from the client
		fmt.Println("msg :", messg)

		broadcasting("[" + name + "]: " + messg) //send the message to all connected clients
	}
}

//now a disconected client from the client slice
func removeClient (conn net.Conn){
	clientMutex.Lock()

	for i, client:= range clients {
		if client.conn == conn {
			clients = append(clients[:i], clients[i+1:]...)
			break
		}
	}

	clientMutex.Unlock()
}

func broadcasting(message string){
	clientMutex.Lock() 
	//after the mutex we can safly iterate through every connected client
	for _, client := range clients {
		//send the message to this client
		_, err := client.conn.Write([]byte(message + "\n"))
		if err != nil{
			fmt.Println("error", err)
		}
	}
	clientMutex.Unlock() 
}
