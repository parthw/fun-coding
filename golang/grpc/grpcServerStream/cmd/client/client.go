package client

import (
	"context"
	"io"
	"log"

	greetpb "example.com/grpcServerStream/api/gRPC"
	"google.golang.org/grpc"
)

// StartClient to start gRPC client
func StartClient() {
	log.Println("Started gRPC client")
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := greetpb.NewGreetServiceClient(conn)
	respClient, err := client.Greet(context.Background(), &greetpb.GreetRequest{
		FirstName: "Parth",
		LastName:  "Wadhwa",
	})
	if err != nil {
		log.Fatalln(err)
	}

	for {
		resp, err := respClient.Recv()
		if err == io.EOF {
			log.Println("Recieved complete stream")
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		log.Println(resp.GetResult())
	}

}
