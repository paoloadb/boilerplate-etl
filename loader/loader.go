package main

import (
	"fmt"
	"net"
	"log"
	"time"
)

func main() {


	ipc, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {
		log.Fatalln(err)
	}
	defer ipc.Close()

	for {
		conn, err := ipc.Accept()

		if err != nil {
			log.Fatalln(err)
		}

		go func (conn net.Conn) {
			buf := make([]byte,2048)

			len, err := conn.Read(buf)
			if err != nil {
				log.Fatalln(err)
			}
			loadToSomewhere(buf, len)
			time.Sleep(1 * time.Second)

			conn.Close()
		}(conn)
	}

}

func loadToSomewhere(data []byte, len int) {
	fmt.Println("Loaded",string(data[:len]))
}