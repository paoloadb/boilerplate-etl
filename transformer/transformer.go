package main

import (
	"fmt"
	"net"
	"log"
	"time"
)

func main() {

	srcIpc, err := net.Listen("tcp", "127.0.0.1:9998")
	if err != nil {
		log.Fatalln(err)
	}
	defer srcIpc.Close()

	for {
		conn, err := srcIpc.Accept()

		if err != nil {
			log.Fatalln(err)
		}

		go func (conn net.Conn) {
			buf := make([]byte,2048)

			len, err := conn.Read(buf)
			if err != nil {
				log.Fatalln(err)
			}
			relayToLoader(buf, len)
			time.Sleep(1 * time.Second)

			conn.Close()
		}(conn)
	}

}

func relayToLoader(data []byte, len int) {
	for {
		destIpc, err := net.Dial("tcp", "127.0.0.1:9999")
		if err != nil {
			log.Println(err)
		} else {
			
		destIpc.Write(data)
		fmt.Println("Relaying",string(data[:len]))
			break
		}
	}

}