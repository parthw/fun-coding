package server

import (
	"io"
	"log"
	"net"

	greetpb "example.com/grpcClientStream/api/grpc"
	"google.golang.org/grpc"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

func (*server) Greet(stream greetpb.GreetService_GreetServer) error {
	log.Println("Server Greet Service is invoked")
	result := "Hello "
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// return the result
			return stream.SendAndClose(&greetpb.GreetResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Println(err)
			return err
		}
		fname, lname := req.GetFirstName(), req.GetLastName()
		result += fname + " " + lname + " "

	}
}

// StartServer to start grpc server
func StartServer() {
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
