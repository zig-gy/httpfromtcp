package main

import (
	"fmt"
	"log"
	"net"

	"github.com/zig-gy/httpfromtcp/internal/request"
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
		requestLine, err := request.RequestFromReader(conn)
		if err != nil {
			log.Fatalln("error reading request: ", err)
		}
		fmt.Println("Request line:")
		fmt.Println("- Method:", requestLine.RequestLine.Method)
		fmt.Println("- Target:", requestLine.RequestLine.RequestTarget)
		fmt.Println("- Version:", requestLine.RequestLine.HttpVersion)
		fmt.Println("connection closed")
	}
}
