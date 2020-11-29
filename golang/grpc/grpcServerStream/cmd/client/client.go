package client

import (
	"context"
	"log"

	greetpb "example.com/grpcServerStream/api/gRPC"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := greetpb.NewGreetServiceClient(conn)
	respClient := client.Greet(context.Background(), &greetpb.GreetRequest{
		FirstName: "Parth",
		LastName:  "Wadhwa",
	})
	respClient
}
