package client

import (
	"context"
	"io"
	"log"
	"sync"

	greetpb "example.com/grpcBiDirectionalSteaming/api/grpc"
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
		log.Fatalln(err)
	}
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		requests := []*greetpb.GreetRequest{
			{
				Fname: "Parth",
				Lname: "Wadhwa",
			},
			{
				Fname: "Sid",
				Lname: "Wadhwa",
			},
		}

		for _, v := range requests {
			if err := client.Send(v); err != nil {
				log.Fatalln(err)
			}
		}
		client.CloseSend()
		wg.Done()
	}()

	go func() {
		for {
			resp, err := client.Recv()
			if err == io.EOF {
				// Recieved all result
				log.Println(resp.GetResult())
				break
			}
			if err != nil {
				log.Fatalln(err)
			}
			log.Println(resp.GetResult())
		}
		wg.Done()
	}()

	wg.Wait()
}
