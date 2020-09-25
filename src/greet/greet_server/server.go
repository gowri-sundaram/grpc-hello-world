package main

import (
    "fmt"
    "google.golang.org/grpc"
    "grpc-hello-world/src/greet/greetpb"
    "log"
    "net"
)

type server struct {}
func main() {
    fmt.Println("Hello")
    listener, err := net.Listen("tcp", "0.0.0.0:50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    s := grpc.NewServer()
    greetpb.RegisterGreetServiceServer(s, &server{})

    if err := s.Serve(listener); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}