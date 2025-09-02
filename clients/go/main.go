package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Cant connet to server", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to the server. Write your messages, or type 'EXIT' to close")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		msg := scanner.Text()
		if msg == "EXIT" {
			break
		}

		_, err := fmt.Fprintln(conn, msg)
		if err != nil {
			fmt.Println("Error sending message", err)
			break
		}
	}
}
