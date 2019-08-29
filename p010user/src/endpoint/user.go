package endpoint

import (
	"context"
	grpc_err "translate/P10User/src/error/grpc"
	pb "translate/P10User/src/pb/p010user"
	"translate/P10User/src/service"
)

type UserEndpoints struct {

}

func (*UserEndpoints) Login(ctx context.Context, req *pb.LoginRequest) (resp *pb.LoginReply,err error) {

	// todo: validate request so on

	service := &service.UserService{}
	token, serviceError := service.Login(req.Username, req.Password)

	if grpcError, ok := serviceError.(*grpc_err.GrpcError); ok {
		resp = &pb.LoginReply{
			Error: &pb.Error{
				Code: grpcError.Code,
				Message: grpcError.Message,
			},
		}
		return
	}

	if serviceError != nil {
		err = serviceError
		return
	}

	resp = &pb.LoginReply{
		Token:token,
	}
	return
}

