package main

import (
	"centrifugo-example/apiproto"
	"context"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:10000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	api := apiproto.NewCentrifugoApiClient(conn)
	resp, err := api.Info(context.Background(), &apiproto.InfoRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Response: %s", resp)
}
