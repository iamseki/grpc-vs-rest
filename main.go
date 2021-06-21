package main

import (
	"log"
	"net"
	"net/http"

	"github.com/iamseki/grpc-vs-rest/grpc/proto"
	"github.com/iamseki/grpc-vs-rest/grpc/server"

	"github.com/iamseki/grpc-vs-rest/rest/handler"
	"google.golang.org/grpc"
)

func main() {
	go runGRPC()
	runHTTP()
}

func runGRPC() {
	lis, err := net.Listen("tcp", ":50502")
	if err != nil {
		log.Fatalln("Error on listen on TCP:50502:", err)
	}

	s := grpc.NewServer()
	server := server.NewCryptoPricingServer()
	proto.RegisterCryptoPricingServiceServer(s, server)

	log.Printf("Starting gRPC listener on port :50502")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func runHTTP() {
	http.HandleFunc("/pricing/eth/", handler.CryptoPricingHandler)
	log.Printf("Starting HTTP server on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
