package server

import (
	"context"
	"log"
	"net"

	greetpb "example.com/grpcUnary/pb"
	"google.golang.org/grpc"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	log.Println("Greet Service is invoked")
	firstName, lastName := req.GetGreeting().GetFirstName(), req.GetGreeting().GetLastName()
	result := "Hello " + firstName + " " + lastName
	return &greetpb.GreetResponse{
		Result: result,
	}, nil
}

//StartServer to start grpc server
func StartServer() {
	log.Println("Starting grpc server")
	lis, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(grpcServer, &server{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
