/*
Package myredis
This is my entry file
*/
package main

import (
	"errors"
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":6370")
	if err != nil {
		log.Fatalf("cannot listen on :6370: %v", err)
	}
	log.Print("My redis is listning on port 6370")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print("Cannot connect")
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 4096)
	var pending []byte
	for {

		n, err := conn.Read(buf)
		if 0 < n {
			pending = append(pending, buf[:n]...)
			log.Printf("%s read %d bytes : %q", conn.RemoteAddr(), n, buf[:n])
			if checkIfReadHasCRLF(pending) {
				log.Print("Proccessed pending data : ", pending)
				//			result := processdata()
				// implement a function in which the commands are stored in a byte buffer by breaking them
				// using /r/n delimeter and then put it to process function which need not to process and can just print
				// and send back msg like "printed {this command}"
			}
		}
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Print("Connections closed for : ", conn.RemoteAddr())
				break
			} else {
				log.Printf("read error from %s: %v", conn.RemoteAddr(), err)
			}
		}
		return
	}
}

func checkIfReadHasCRLF(data []byte) bool {
	return len(data) >= 2 &&
		data[len(data)-2] == 13 &&
		data[len(data)-1] == 10
}
