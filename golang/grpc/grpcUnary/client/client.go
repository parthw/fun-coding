package client

import (
	"context"
	"log"

	greetpb "example.com/grpcUnary/pb"
	"google.golang.org/grpc"
)

// StartClient to start the grpc client
func StartClient() {
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := greetpb.NewGreetServiceClient(conn)
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Parth",
			LastName:  "Wadhwa",
		},
	}
	resp, err := client.Greet(context.Background(), req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(resp.GetResult())
}
