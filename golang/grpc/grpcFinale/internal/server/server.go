package server

import (
	"context"
	"fmt"
	"log"
	"math"
	"net"
	"os"
	"os/signal"
	"time"

	"example.com/grpcFinale/api/sqrtpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

type server struct {
	sqrtpb.UnimplementedSqrtServiceServer
}

// Sqrt service registered with gRPC server
func (*server) Sqrt(ctx context.Context, req *sqrtpb.SqrtRequest) (*sqrtpb.SqrtResponse, error) {

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			log.Println("Client cancelled the request")
			return nil, status.Error(codes.DeadlineExceeded, "Client canncelled the request")
		}
		time.Sleep(1 * time.Second)
	}
	if req.GetNum() < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Sqrt of negative number %v is not supported", req.GetNum()))
	}
	resp := &sqrtpb.SqrtResponse{
		Result: math.Sqrt(float64(req.GetNum())),
	}
	return resp, nil
}

// StartServer to start gRPC server
func StartServer() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	lis, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		log.Fatalln("Failed to listen on port 9090 with error - ", err)
	}

	creds, err := credentials.NewServerTLSFromFile("/Users/parth/.certs/server.crt", "/Users/parth/.certs/server.key")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	sqrtpb.RegisterSqrtServiceServer(grpcServer, &server{})
	log.Println("Successfully registered sqrt service to  gRPC server")

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalln("Failed to server grpc server with error - ", err)
		}
	}()

	exitChan := make(chan os.Signal, 1)
	signal.Notify(exitChan, os.Interrupt)
	<-exitChan
	log.Println("Stopping the server and listener")
	grpcServer.Stop()
	lis.Close()
	log.Println("Server is exited gracefully.")
}
