package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"translate/P10User/src/config"
	"translate/P10User/src/endpoint/v1"
	pb "translate/P10User/src/pb/p010user/v1"
	mLog "translate/P10User/src/lib/log"
)

func main() {

	config.LoadEnv("")
	mLog.Setup()

	lis, err := net.Listen("tcp", ":6000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServer(s, &v1.UserEndpoints{})
	reflection.Register(s)

	if err:= s.Serve(lis); err!= nil {
		log.Fatalf("failed to serve: %v", err)
	}

}