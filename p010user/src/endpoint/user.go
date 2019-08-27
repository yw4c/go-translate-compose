package endpoint

import (
	"context"
	pb "translate/P10User/src/pb/user"
	"translate/P10User/src/service"
)

type UserEndpoints struct {

}

func (*UserEndpoints) Login(ctx context.Context, req *pb.LoginRequest) (resp *pb.LoginReply,err error) {

	// todo: validate request so on

	service := &service.UserService{}
	token, err := service.Login(req.Username, req.Password)

	resp = &pb.LoginReply{
		Token:token,
	}

	return

}

