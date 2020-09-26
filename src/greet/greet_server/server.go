package main

import (
    "context"
    "google.golang.org/grpc"
    "grpc-hello-world/src/greet/greetpb"
    "log"
    "net"
)

type server struct {}

func (s *server) Greet(ctx context.Context, in *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
    log.Printf("Greet function was invoked with: %v\n", in)
    name := in.GetGreeting().GetFirstName() + " " + in.GetGreeting().GetLastName()
    resp := &greetpb.GreetResponse {
        Result: "Hello " + name,
    }
    return resp, nil
}

func main() {
    log.Println("Hello")
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