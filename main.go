package main

import (
	"log"
	"net"
)

func main() {
	lis, err :=net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("cannot create listener %s", err)
	}
}  