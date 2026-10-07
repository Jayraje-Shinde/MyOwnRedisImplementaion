/*
Package myredis
This is my entry file
*/
package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":6370")
	log.Print("Started Listning on 6370 port")
	if err != nil {
		fmt.Print("Can't connect due to ", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Cannot connect")
			continue
		}

		go handleConnections(conn)
	}
}

func handleConnections(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 4096)

	for {
		n, err := conn.Read(buf)
		if 0 < n {
			log.Printf("%s read %d bytes : %q", conn.RemoteAddr(), n, buf[:n])
		}
		if err != nil {
			if err == io.EOF {
				log.Print("Connections closed")
			}
			log.Print("Error occured : ", err)
		}
	}
}
