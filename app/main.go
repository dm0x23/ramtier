package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

var _ = net.Listen
var _ = os.Exit

func main() {

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go handleConnection(conn)

	}

}

func handleConnection(conn net.Conn) {
	for {
		buf := make([]byte, 1024)

		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Connection terminated")
			os.Exit(1)
		}

		command := strings.Split(string(buf[:n]), "\r\n")

		switch strings.ToUpper(command[2]) {
		case "PING":
			conn.Write([]byte("+PONG\r\n"))
		case "ECHO":
			fmt.Fprintf(conn, "$%d\r\n%s\r\n", len(command[4]), command[4])
		}
	}
}
