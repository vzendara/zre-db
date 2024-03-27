package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	port := os.Args[1]
	addr := fmt.Sprintf(":%s", port)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("Failed to start listener", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Printf("Started listening at %s", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection")
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		line, err := reader.ReadBytes(byte('\n'))
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading bytes: ", err)
			}
			return
		}

		fmt.Print(string(line))
		conn.Write([]byte(line))
	}
}
