package main

import (
	"net"
	"log"
	"fmt"
)

func main() {

	sendData("Testing 123")

}

func sendData(data string) {
	sck, err := net.Dial("tcp", "127.0.0.1:9998")
	if err != nil {
		log.Fatalln(err)
	}

	sck.Write([]byte(data))
	fmt.Println("Sent", data)

	sck.Close()
}