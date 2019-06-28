package main

import (
	"context"
	kitGrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
	"log"
	"net"
	"translate/P10User/src/endpoint"
	"translate/P10User/src/pb"
	"translate/P10User/src/service"
)

var (
	// etcd服务地址
	etcdServer = "consul1:8400"
	// 服务的信息目录
	prefix     = "/services/user/"
	// 当前启动服务实例的地址
	instance   = "p10user:50052"
	// 服务实例注册的路径
	key        = prefix + instance
	// 服务实例注册的val
	value      = instance
	ctx        = context.Background()
	// 服务监听地址
	serviceAddress = ":50052"
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
		loginHandler: kitGrpc.NewServer(endpoint.MakeLoginEndpoint(&service.UserService{}), decodeRequest, encodeResponse),
	}


	lis, err := net.Listen("tcp", serviceAddress)
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
