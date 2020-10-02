package main

import (
    "context"
    "fmt"
    "google.golang.org/grpc"
    "grpc-hello-world/src/greet/greetpb"
    "io"
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
    //doUnary(c)
    doServerStreaming(c)

}

func doUnary(c greetpb.GreetServiceClient) {
    fmt.Println("Doing unary request...")
    greeting := &greetpb.Greeting{
        FirstName: "Okabe",
        LastName: "Rintarou!",
    }
    res, err := c.Greet(context.Background(), &greetpb.GreetRequest{Greeting: greeting})
    if err != nil {
        log.Fatalf("Received error on invoking greet rpc: %v", err)
    }
    log.Printf("Response from Greet: %v", res)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
    fmt.Println("Doing server streaming request...")
    greeting := &greetpb.Greeting{
        FirstName: "Okabe",
        LastName: "Rintarou!",
    }

    resStream, err := c.GreetManyTimes(context.Background(), &greetpb.GreetRequest{Greeting: greeting})
    if err != nil {
        log.Fatalf("Received error on invoking greet rpc: %v", err)
    }

    for {
        msg, err := resStream.Recv()
        if err == io.EOF {
            // We've reached the end of the stream.
            break
        }
        if err != nil {
            log.Fatalf("Error on reading stream: %v", err)
        }
        log.Printf("Response from GreetManyTimes: %v", msg)
    }
}
