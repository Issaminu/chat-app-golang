package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Define the server address (IP and Port)
	serverAddr := "localhost:8080"

	// Connect to the server
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Printf("Error connecting to the server: %v\n", err)
		return
	}
	defer conn.Close()
	fmt.Printf("Connected to the server\n\n")

	done := make(chan bool)
	go sendMessage(conn, done)
	go receiveMessage(conn, done)
	<-done
	<-done
}

func sendMessage(conn net.Conn, done chan bool) {
	for {
		fmt.Print("$ You (Client): ")
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

		// Print the server's response
		fmt.Printf("\r\n\rServer: %s", buffer[:n])
		fmt.Print("\n$ You (Client): ")
	}
}
