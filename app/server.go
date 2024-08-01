package main

import (
	"fmt"
	"net"
	"os"
)

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

		dat := make([]byte, 1000)
		for {
			n, err := conn.Read(dat)
			if err != nil {
				fmt.Println(err.Error())
				break
			} else if n > 0 {
				fmt.Println(string(dat))
			}

			conn.Write([]byte("+PONG\r\n"))
		}
	}
}
