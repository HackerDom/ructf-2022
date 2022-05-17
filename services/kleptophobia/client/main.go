package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "kleptophobia/models"
	"log"
	"time"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewKleptophobiaClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	r, err := c.Register(ctx, &pb.RegisterRequest{Person: &pb.PrivatePerson{
		Username:  "username",
		Password:  "passwd",
		FirstName: "Name",
	}})

	if err != nil {
		log.Printf("could not register user: %v", err)
	}

	if r.Status != pb.RegisterReply_OK {
		log.Printf("Can not register user: " + r.GetMessage())
	} else {
		log.Println("Success!")
	}
}
