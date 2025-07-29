package main

import "google.golang.org/grpc/credentials/insecure"
import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/xds"

	pb "github.com/yanivamram/td-app/gen/helloworld"
)

func main() {
	target := "xds:///sayhello-server.default.svc.cluster.local"

	// Use xDS-enabled client
	conn, err := grpc.Dial(
		target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithResolvers(xds.NewResolverWithConfigForTesting(nil)),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	name := "World"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", resp.GetMessage())
}