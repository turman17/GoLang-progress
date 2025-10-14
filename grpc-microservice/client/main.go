package main

import (
	"context"
	"log"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/turman17/grpc-microservice/api/gen"
)

func main(){
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Dial: %v", err)
	}

	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	helloResp, err := c.SayHello(ctx, &pb.HelloRequest{Name: "Turman"})
	if err != nil {
		log.Fatalf("SayHello: %v", err)
	}
	fmt.Println("SayHello:", helloResp.GetMessage())

	timeResp, err := c.GetTime(ctx, &pb.TimeRequest{})
	if err != nil {
		log.Fatalf("SayTime: %v", err)
	}
	fmt.Println("SayTime:", timeResp.GetIso8601())
}