package main

import (
	"net"
	"log"
)

func main() {

	sendData("Testing 123")

}

func sendData(data string) {
	sck, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		log.Fatalln(err)
	}

	sck.Write([]byte(data))
	// buf := make([]byte, 1024)

	// _, err = sck.Read(buf)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	sck.Close()
}