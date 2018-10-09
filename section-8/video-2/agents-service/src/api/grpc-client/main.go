package main

import (
	"context"
	"fmt"
	"os"
	"time"

	pb "github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-8/video-2/agents-service/src/api/transports/grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure(), grpc.WithTimeout(time.Second))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}
	defer conn.Close()

	c := pb.NewAgentsClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	getAgentReply, err := c.GetAgentByID(ctx, &pb.GetAgentByIDRequest{
		AgentID: 10,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}
	fmt.Println(getAgentReply)
}
