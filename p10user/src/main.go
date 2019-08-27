package main

import (

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"translate/P10User/src/endpoint"
	pb "translate/P10User/src/pb/user"
)

func main() {
	lis, err := net.Listen("tcp", ":6000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}


	s := grpc.NewServer()
	pb.RegisterUserServer(s, &endpoint.UserEndpoints{})
	reflection.Register(s)

	if err:= s.Serve(lis); err!= nil {
		log.Fatalf("failed to serve: %v", err)
	}

}