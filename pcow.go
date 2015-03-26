package main

import (
	//	"log"
	"net"
)

func main() {
	fox := "fa"
	lis, err := net.Listen("tcp", "0.0.0.0:7778")
	if err != nil {
		panic(err)
	}

	for {
		conn, _ := lis.Accept()

	}
}
