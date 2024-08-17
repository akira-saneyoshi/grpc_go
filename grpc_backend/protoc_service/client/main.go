package main

import (
	"context"
	"fmt"
	"log"
	"udemy_study/go_grpc/pb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewFileServiceClient(conn)
	callListFiles(client)
}

func callListFiles(client pb.FileServiceClient) {
	res, err := client.ListFiles(context.Background(), &pb.ListFilesRequest{})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res.GetFilenames())
}
