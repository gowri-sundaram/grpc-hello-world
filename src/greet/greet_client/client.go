package main

import (
    "context"
    "google.golang.org/grpc"
    "grpc-hello-world/src/greet/greetpb"
    "log"
)

func main() {
    log.Println("Hello, from the client!")
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Could not connect: %v", err)
    }
    defer conn.Close()

    c := greetpb.NewGreetServiceClient(conn)
    greeting := &greetpb.Greeting{
        FirstName: "Okabe",
        LastName: "Rintarou!",
    }
    res, err := c.Greet(context.Background(), &greetpb.GreetRequest{Greeting: greeting})
    if err != nil {
        log.Fatalf("Received error on invoking greet rpc: %f", err)
    }
    log.Printf("Response from greet: %v", res)
}
