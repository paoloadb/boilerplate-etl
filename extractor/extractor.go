package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

func main() {

	for {
		data := time.Now().Local().Local().UnixMicro()
		sendData(strconv.FormatInt(int64(data), 10))
		time.Sleep(1 * time.Second)
	}


}

func sendData(data string) {
	for {
		sck, err := net.Dial("tcp", "127.0.0.1:9998")
		if err != nil {
			log.Println(err)
		} else {
			sck.Write([]byte(data))
			fmt.Println("Sent", data)	
			sck.Close()
			break
		}
	}
}