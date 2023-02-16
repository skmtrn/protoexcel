package main

import (
	"log"

	"github.com/skmtrn/protoexcel/excelizegrpc"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {

	te := excelizegrpc.TestEntry{Id: 3, Lp: "example LP", Value: 0.12}

	var ste []*excelizegrpc.TestEntry
	ste = append(ste, &te)

	entries := excelizegrpc.Entries{Entries: ste}

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := excelizegrpc.NewExcelizeServiceClient(conn)

	response, err := c.ProvideEntries(context.Background(), &entries)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response)

}
