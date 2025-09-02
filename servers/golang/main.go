package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// Port listening
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error ", err)
		return
	}
	defer ln.Close()

	fmt.Println("Server live and running in port 8080")

	for {

		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error trying to accept conn", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("New connection %s\n", conn.RemoteAddr().String())

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		msg := scanner.Text()
		fmt.Printf("«Client msg» (%s): %s\n", conn.RemoteAddr().String(), msg)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading client", err)
	}
}
