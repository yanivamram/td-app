package main

import (
	"context"
	"log"
	"time"

	pb "your/module/path/helloworld"

	"google.golang.org/grpc"
	_ "google.golang.org/grpc/xds" // Register xDS resolver
)

func main() {
	conn, err := grpc.Dial(
		"xds:///greeter",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "TrafficDirector"})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Greeting: %s", resp.GetMessage())
}
