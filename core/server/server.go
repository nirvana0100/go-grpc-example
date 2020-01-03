package main

import (
	"fmt"
	pb "go-grpc-example/protocol"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
)

const Address = "127.0.0.1:8080"

type helloService struct {
}

func (h *helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	resp := new(pb.HelloReply)
	if in.Name == "Hi" {
		resp.Message = "Hello"
		return resp, nil
	}
	resp.Message = "not correct say again "
	return resp, nil
}
func main() {
	conn, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &helloService{})
	fmt.Println("listen on " + Address)
	s.Serve(conn)
}
