package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"context"

	pb "github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-3/video-3/proto"
	"google.golang.org/grpc"
)

const (
	address       = "localhost:50051"
	defaultNumber = uint32(1)
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewFibonacciClient(conn)

	// Contact the server and print out its response.
	num := defaultNumber
	if len(os.Args) > 1 {
		tmp, err := strconv.ParseUint(os.Args[1], 10, 32)
		if err != nil {
			log.Fatalf("Wrong Argument: %s", os.Args[1])
		}
		num = uint32(tmp)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Calculate(ctx, &pb.FibonacciRequest{Number: num})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Fibonacci(%d): %d\nProcessing Time: %d", num, r.Result, r.ProcessingTime)
}
