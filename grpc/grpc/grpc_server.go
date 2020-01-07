package main

/**
Here grpc-gateway/runtime gives a multiplexer, which registers the grpc http endpoints to multiplexers.
This is not a working example because, I have not started grpc server yet.

Much better example is present here :
https://github.com/alextanhongpin/go-grpc-gateway.git
*/
import (
	"context" // Use "golang.org/x/net/context" for Golang version <= 1.6
	"flag"
	"net"
	"log"
	"google.golang.org/grpc"
	"fmt"

	gw "../proto" // Update
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9090))
	if err != nil {
			log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	gw.RegisterYourServiceServer(grpcServer, &MyServer{})
	grpcServer.Serve(lis)
}

type MyServer struct{

}

func (t *MyServer) Echo(context.Context, *gw.StringMessage) (*gw.StringMessage, error){
	v := gw.StringMessage{
		Value : "Hi from Post",
	}
	return &v , nil
}
func (t *MyServer) Hello(context.Context, *gw.StringMessage) (*gw.StringMessage, error){
	v := gw.StringMessage{
		Value : "Hi from Get",
	}
	return &v , nil
}

// Check the CPU usage with multiple http requests. Every HTTP request starts a new thread and consumes full core in this case.
func (t *MyServer) HelloFromInfitineLoop(context.Context, *gw.StringMessage) (*gw.StringMessage, error){

	sum := 1
	for sum < 1000 {
		sum =1
	}

	v := gw.StringMessage{
		Value : "Hello from infinite loop",
	}
	return &v , nil
}

