package server

import (
	"io"
	"log"
	"net"

	greetpb "example.com/grpcBiDirectionalSteaming/api/grpc"

	"google.golang.org/grpc"
)

type myserver struct {
	greetpb.UnimplementedGreetServiceServer
}

func (*myserver) Greet(stream greetpb.GreetService_GreetServer) error {
	log.Println("Invoking Greet Service Server")
	var result []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// got the request
			log.Println("Recieved complete stream")
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		result = append(result, "Hello "+req.GetFname()+" "+req.GetLname())
	}

	for _, v := range result {
		if err := stream.Send(&greetpb.GreetResponse{
			Result: v,
		}); err != nil {
			log.Fatalln(err)
		}
	}
	log.Println("Sent the stream")
	return nil
}

// StartServer to start grpc server
func StartServer() {
	lis, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(grpcServer, &myserver{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
