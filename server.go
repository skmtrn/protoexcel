package main

import (
	"log"
	"net"
	"protoexcel/excelizegrpc"

	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":9000")

	if err != nil {

		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	s := excelizegrpc.Server{}

	grpcServer := grpc.NewServer()

	excelizegrpc.RegisterExcelizeServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {

		log.Fatalf("Failed to serve gRPC on port 9000: %v", err)

	}

}
