package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	tcpListener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatalln("error setting up listener:", err)
	}
	defer tcpListener.Close()

	for {
		conn, err := tcpListener.Accept()
		if err != nil {
			log.Fatalln("error accepting request:", err)
		}
		fmt.Println("connection established")
		for line := range getLinesChannel(conn) {
			fmt.Println(line)
		}
		fmt.Println("connection closed")
	}
}
