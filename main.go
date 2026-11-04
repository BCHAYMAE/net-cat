package main

import (
	"bufio" //helps us read client input
	"fmt"
	"net" //package provides networking functionality in Go
	"os"
	"strconv" //convert string to int
	"strings"
	"sync" //provides synchronization tools
	"time" //provides date and time func
)

type Client struct {
	name string   //client name
	conn net.Conn //client's network connection
}

var clients []Client        //slice to store the connected clients
var clientMutex sync.Mutex  //mutex to synchronize access to the clients slice
var messageHistory []string //slice to store previous chat messages

func main() {
	port := "8989" //set the default port by choice

	if len(os.Args) == 2 {
		port = os.Args[1]
	} else if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}

	//we should validate the port before calling net.Listen
	//check if the port is valid
	if !isValidPort(port) {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}

	listener, err := net.Listen("tcp", ":"+port) //create a TCP on the selected port
	//now we check if the listener was created successfully
	if err != nil {
		fmt.Println("error", err)
		return
	}
	defer listener.Close() //close the listener when the function returns
	fmt.Println("Listening on port", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error", err)
			continue
		}

		fmt.Println("client conn: ", conn.RemoteAddr()) //gives you the client's network address

		go handleClient(conn)
	}
}

// check if the port is valid
func isValidPort(port string) bool {
	portNumber, err := strconv.Atoi(port) //convert string to integer

	if err != nil || portNumber < 1 || portNumber > 65535 {
		return false //invalid port
	}

	return true //valid port
}

// check if the username is valid
func isValidUsername(name string) bool {
    if strings.TrimSpace(name) == "" || len(name) > 20 {
        return false
    }

    return true
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	defer removeClient(conn) //remove the client from the slice when the function returns

	//send the username prompt to the client
	conn.Write([]byte("[ENTER YOUR NAME]: "))

	//create a reader to receive input from the client
	reader := bufio.NewReader(conn)

	//read the username entered by the client
	var name string
	var err error
	for {
		name, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading username:", err)
			return
		}

		//remove spaces and the newline from the username
		name = strings.TrimSpace(name)

		//check if the username is valid
		if !isValidUsername(name) {
			conn.Write([]byte("Invalid username. Must contain 1 to 20 characters.\n"))
			conn.Write([]byte("[ENTER YOUR NAME]: "))
			continue
		}

		clientMutex.Lock()

		//check if the server already has 10 clients
		if len(clients) >= 10 {
			clientMutex.Unlock()
			conn.Write([]byte("Chat is full. Try again later.\n"))
			return
		}

		//we need to see if the username is already taken
		usernameTaken := false

		for _, client := range clients {
			if client.name == name {
				usernameTaken = true
				break
			}
		}

		if usernameTaken {
			clientMutex.Unlock()
			conn.Write([]byte("Username already taken.\n"))
			conn.Write([]byte("[ENTER YOUR NAME]: "))
			continue
		}

		//create a new client with its name and connection
		newClient := Client{
			name: name,
			conn: conn,
		}

		//add the new client to the clients slice
		clients = append(clients, newClient)

		//send previous messages to the new client
		for _, message := range messageHistory {
			_, err := conn.Write([]byte(message + "\n"))
			if err != nil {
				fmt.Println("error sending history:", err)
				break
			}
		}

		clientMutex.Unlock()
		break
	}

	broadcasting("[" + name + "] has joined our chat...") //annouce that a new client joined

	//now we need to read the data send by the client into the buffer
	for {
		//read the message using the same reader
		messg, err := reader.ReadString('\n')

		//check if reading failed
		if err != nil {
			fmt.Println("error", err)
			break
		}

		//remove spaces and newline
		messg = strings.TrimSpace(messg)

		//don't broadcast empty messages
		if messg == "" {
			continue
		}

		timestamp := time.Now().Format("2006-01-02 15:04:05")
		broadcasting("[" + timestamp + "][" + name + "]: " + messg)
	}
}

// now a disconected client from the client slice
func removeClient(conn net.Conn) {
	clientMutex.Lock()

	var name string //store the username of the disconnected client

	for i, client := range clients {
		if client.conn == conn {
			name = client.name //get the client's username
			clients = append(clients[:i], clients[i+1:]...)
			break
		}
	}

	clientMutex.Unlock()

	//announce that the client left
	if name != "" {
		broadcasting("[" + name + "] has left our chat...")
	}
}

func broadcasting(message string) {
	clientMutex.Lock() //lock shared data

	messageHistory = append(messageHistory, message) //save message in history

	activeClients := make([]Client, len(clients)) //create a copy of the clients
	copy(activeClients, clients)                  //copy connected clients

	clientMutex.Unlock() //unlock before sending messages

	for _, client := range activeClients { //loop through copied clients
		_, err := client.conn.Write([]byte(message + "\n")) //send message

		if err != nil { //check if sending failed
			fmt.Println("Error sending message to", client.name, ":", err)
		}
	}
}
