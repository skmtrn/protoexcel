package excelizegrpc

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) ProvideEntries(ctx context.Context, entries *Entries) (*Response, error) {

	log.Printf("Received entries from client: %s", entries)

	return &Response{Response: "Thank you."}, nil

}
