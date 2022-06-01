package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/nicholasaiello/go-http2-api-example/proto"
	"github.com/nicholasaiello/go-http2-api-example/routes"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

var (
	g errgroup.Group
)

// Chunked 50 + 150
// ----------------------
//      | h1.1  | h2   | h2 push |
// 1000 | 3     | 4    | 2.1     |
// 2000 | 3.7   | 5.6  | 2.4     |

// TODO:
// - add Web Worker support

func startSecureGRPCServer() {
	listener, err := net.Listen("tcp4", ":9443")
	if err != nil {
		log.Fatalf("could not attach listener to port: %v", err)
	}

	creds, err := credentials.NewServerTLSFromFile("./server.crt", "./server.key")
	if err != nil {
		log.Fatalf("Failed to generate credentials %v", err)
	}

	grpc.EnableTracing = true

	opts := []grpc.ServerOption{grpc.Creds(creds)}
	s := grpc.NewServer(opts...)
	proto.RegisterAccountHoldingsServer(s, routes.GRPCRouter())

	reflection.Register(s)

	log.Printf("Starting GRPC server listening on port 9080")

	g.Go(func() error {
		return s.Serve(listener)
	})
}

func startGRPCServer() {
	listener, err := net.Listen("tcp4", ":9080")
	if err != nil {
		log.Fatalf("could not attach listener to port: %v", err)
	}

	grpc.EnableTracing = true

	s := grpc.NewServer()
	proto.RegisterAccountHoldingsServer(s, routes.GRPCRouter())

	reflection.Register(s)

	log.Printf("Starting GRPC server listening on port 9080")

	g.Go(func() error {
		return s.Serve(listener)
	})
}

func startHTTPServer() {
	s := &http.Server{
		Addr:         ":8080",
		Handler:      routes.Router(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return s.ListenAndServe()
	})
}

func startHTTP2Server() {
	s := &http.Server{
		Addr:         ":8443",
		Handler:      routes.Router(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return s.ListenAndServeTLS("./server.crt", "./server.key")
	})
}

func main() {
	startHTTPServer()
	startHTTP2Server()
	startGRPCServer()
	startSecureGRPCServer()

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
