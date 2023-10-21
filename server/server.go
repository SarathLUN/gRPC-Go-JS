package main

import (
	"context"
	"fmt"
	helloworld "github.com/SarathLUN/grpc-go-js/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	helloworld.UnimplementedGreeterServer
}

func (s server) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	message := "Hello, " + req.GetName()
	return &helloworld.HelloReply{Message: message}, nil
}

func (s *server) SayRepeatHello(req *helloworld.RepeatHelloRequest, stream helloworld.Greeter_SayRepeatHelloServer) error {
	for i := 0; i < int(req.Count); i++ {
		msg := fmt.Sprintf("server stream response: %v", i)
		err := stream.Send(&helloworld.HelloReply{Message: "Hello, " + msg + "!"})
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	log.Println("Starting gRPC server")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listent: %v", err.Error())
	}
	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err.Error())
	}
}
