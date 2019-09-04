package main

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"translate/P10User/src/endpoint/v1"
	pb "translate/P10User/src/pb/p010user/v1"
)

func main() {

	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stdout)

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