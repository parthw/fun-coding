package server

import (
	"log"
	"net"
	"strconv"

	greetpb "example.com/grpcServerStream/api/gRPC"
	"google.golang.org/grpc"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

func (*server) Greet(req *greetpb.GreetRequest, resp greetpb.GreetService_GreetServer) error {
	fName, lName := req.GetFirstName(), req.GetLastName()
	for i := 0; i < 10; i++ {
		result := strconv.Itoa(i) + ". Hello " + fName + " " + lName + " ✌️"

		if err := resp.Send(&greetpb.GreetResponse{
			Result: result,
		}); err != nil {
			log.Println(err)
		} else {
			log.Println("Successfully sent message ID ", strconv.Itoa(i))
		}
	}

	return nil
}

func main() {
	log.Println("Starting gRPC server")

	lis, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		log.Fatalln(err)
	}
	gRPCServer := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(gRPCServer, &server{})
	if err := gRPCServer.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
