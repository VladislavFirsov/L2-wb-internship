package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("Server is listening")
	lstn, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := lstn.Accept()

		if err != nil {
			fmt.Println("connection issue", err)
			conn.Close()
			continue
		}

		fmt.Println("connected")

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("You got a message: %s", message)
	}
}
