package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	udpAddress, err := net.ResolveUDPAddr("udp", ":42069")
	if err != nil {
		log.Fatalln("could not resolve udp address:", err.Error())
	}

	udpConnection, err := net.DialUDP("udp", nil, udpAddress)
	if err != nil {
		log.Fatalln("could not dial udp address:", err.Error())
	}
	defer udpConnection.Close()

	udpReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, err := udpReader.ReadString('\n')
		if err != nil {
			fmt.Println("could not read string:", err.Error())
		}

		if _, err := udpConnection.Write([]byte(line)); err != nil {
			fmt.Println("could not write to string:", err.Error())
		}
	}
}
