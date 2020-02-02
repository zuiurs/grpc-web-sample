package main

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/zuiurs/grpc-web-sample/protobuf"
	"google.golang.org/grpc"
)

func main() {
	if len(os.Args) != 2 {
		log.Print("not enough argument")
		return
	}
	name := os.Args[1]

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := pb.NewGreeterClient(conn)

	ctx := context.Background()
	resp, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})

	fmt.Println(resp.GetMessage())
}
