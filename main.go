
package main

import (
	"fmt"
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

	fmt.Println("Listening on the port :" + port)
}
