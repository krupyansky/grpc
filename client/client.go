package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc/gen/proto"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}

	client := pb.NewTestApiClient(conn)

	resp, err := client.Echo(context.Background(), &pb.ResponseRequest{Msg: "Hello everyone!"})
	if err != nil {
		return
	}

	fmt.Println(resp)
	fmt.Println(resp.Msg)
}
