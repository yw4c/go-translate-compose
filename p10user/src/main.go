package main

import (
	"context"

	kitGrpc "github.com/go-kit/kit/transport/grpc"
	"log"
	"net"
	pb "translate/P10User/src/pb"
	"google.golang.org/grpc"
)

const (
	port = ":50052"
)


type UserServer struct {
	loginHandler kitGrpc.Handler
}
func (us UserServer) Login(ctx context.Context,request *pb.LoginRequest) (*pb.LoginReply, error) {
	_, rsp, err := us.loginHandler.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return rsp.(*pb.LoginReply), err
}

func main() {


	userServer := UserServer{
		loginHandler: kitGrpc.NewServer(makeLoginEndpoint(&service{}), decodeRequest, encodeResponse),
	}



	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	gs := grpc.NewServer()
	pb.RegisterUserServer(gs, userServer)
	gs.Serve(lis)

}

func decodeRequest(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

func encodeResponse(_ context.Context, rsp interface{}) (interface{}, error) {
	return rsp, nil
}
