package main

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/turman17/grpc-microservice/api/gen"
)

type server struct{
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error){
	name := req.GetName()
	if name == "" {
		name = "world";
	}

	return &pb.HelloReply{Message: "Hello " + name + "!"}, nil
}

func (s *server) GetTime(ctx context.Context, _ *pb.TimeRequest) (*pb.TimeReply, error){
	return &pb.TimeReply{ Iso8601: time.Now().UTC().Format(time.RFC3339)}, nil
}

func main(){
	lis, err := net.Listen("tcp", ":50051")
	if err != nil{
		log.Fatalf("listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, &server{})

	reflection.Register(grpcServer)
	log.Println("gRPC server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("serve: %v", err)
	}
}