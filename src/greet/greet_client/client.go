package main

import (
    "fmt"
    "google.golang.org/grpc"
    "grpc-hello-world/src/greet/greetpb"
    "log"
)

func main() {
    fmt.Println("Hello, from the client!")
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Could not connect: %v", err)
    }
    defer conn.Close()

    c := greetpb.NewGreetServiceClient(conn)
    fmt.Printf("Created client: %f", c)
}
