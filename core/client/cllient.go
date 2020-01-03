package main

import (
	"fmt"
	pb "go-grpc-example/protocol"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const Address = "127.0.0.1:8080"

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()
	client := pb.NewHelloClient(conn)
	hr := new(pb.HelloRequest)
	hr.Name = "grpc"
	r, err := client.SayHello(context.Background(), hr)
	if err != nil {
		grpclog.Fatalln(err)
	}
	fmt.Println(r.Message)
}
