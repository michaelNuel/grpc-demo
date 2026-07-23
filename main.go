package main

import (
	"context"
	"log"
	"net"

	"github.com/michaelNuel/demo-grpc/invoicer"
	"google.golang.org/grpc"
)

type myInvoicerServer struct { 
  invoicer.UnimplementedInvoicerServer
} 

func (s myInvoicerServer ) Create(ctx context.Context, req *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
  return &invoicer.CreateResponse{
	Pdf: []byte("pdf content"), 
	Docx: []byte("docx content"),
  }, nil 
}

func main() { 
	lis, err :=net.Listen("tcp", ":8080") 
	if err != nil { 
		log.Fatalf("cannot create listener %s", err)
	} 

	serverRegister := grpc.NewServer()
	service := &myInvoicerServer{} 
	invoicer.RegisterInvoicerServer(serverRegister, service)
    err = 	serverRegister.Serve(lis)
    if err != nil {
	  log.Fatalf("cannot serve %s", err)
   }
}   

