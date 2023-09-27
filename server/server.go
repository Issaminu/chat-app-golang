package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()

	fmt.Printf("Server is listening on port 8080\n\n")

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			continue
		}

		// Handle the connection in a new goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	done := make(chan bool)
	go sendMessage(conn, done)
	go receiveMessage(conn, done)
	<-done
	<-done
}

func sendMessage(conn net.Conn, done chan bool) {
	for {
		fmt.Print("$ You (Server): ")
		reader := bufio.NewReader(os.Stdin)
		message, _ := reader.ReadString('\n')

		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Printf("Error sending data: %v\n", err)
			done <- true
			return
		}
	}
}
func receiveMessage(conn net.Conn, done chan bool) {
	for {
		buffer := make([]byte, 4096)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Printf("Error reading data: %v\n", err)
			done <- true
			return
		}

		// Print the client's response
		fmt.Printf("\r\rClient: %s", buffer[:n])
		fmt.Print("\n$ You (Server): ")
	}
}
