package client

import (
	"context"
	"log"

	greetpb "example.com/grpcClientStream/api/grpc"
	"google.golang.org/grpc"
)

// StartClient to start grpc client
func StartClient() {
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client, err := greetpb.NewGreetServiceClient(conn).Greet(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 5; i++ {
		if err := client.Send(&greetpb.GreetRequest{
			FirstName: "Parth",
			LastName:  "Wadhwa",
		}); err != nil {
			log.Fatalln(err)
		}
	}

	resp, err := client.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(resp.GetResult())
}
