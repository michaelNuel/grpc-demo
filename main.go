package main

import (
	"log"
	"net"

	"github.com/michaelNuel/demo-grpc/invoicer"
)

func main() {
	lis, err :=net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("cannot create listener %s", err)
	}
	invoicer.RegisterInvoicerServer()
}  