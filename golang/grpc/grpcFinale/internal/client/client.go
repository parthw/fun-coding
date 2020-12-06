package client

import (
	"context"
	"log"
	"sync"
	"time"

	"example.com/grpcFinale/api/sqrtpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

// StartClient to start grpc client
func StartClient() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	creds, err := credentials.NewClientTLSFromFile("/Users/parth/.certs/server.crt", "")
	if err != nil {
		log.Fatalln(err)
	}
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalln(err)
	}

	client := sqrtpb.NewSqrtServiceClient(conn)

	// Ideal Request
	req1 := &sqrtpb.SqrtRequest{
		Num: 144,
	}

	// User Error Request
	req2 := &sqrtpb.SqrtRequest{
		Num: -100,
	}

	wg := sync.WaitGroup{}
	wg.Add(3)
	go processReq(client, req1, &wg, 5*time.Second)
	go processReq(client, req1, &wg, 1*time.Second)
	go processReq(client, req2, &wg, 5*time.Second)

	wg.Wait()
	log.Println("Client Processing completed, exiting client")
}

func processReq(client sqrtpb.SqrtServiceClient, req *sqrtpb.SqrtRequest, wg *sync.WaitGroup, t time.Duration) {
	log.Printf("Processing request for input %v\n", req.GetNum())

	ctx, cancel := context.WithTimeout(context.Background(), t)
	defer cancel()

	resp, err := client.Sqrt(ctx, req)
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			// User Created Error
			if respErr.Code() == codes.DeadlineExceeded {
				log.Println("Deadline Excceded")
			}
			log.Println(respErr.Err())
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Println(resp.GetResult())
	}

	wg.Done()
}
