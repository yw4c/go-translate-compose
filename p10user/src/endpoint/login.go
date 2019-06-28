package endpoint
import (
	"context"
	"github.com/go-kit/kit/endpoint"
	pb "translate/P10User/src/pb"
	"translate/P10User/src/service"
)



func MakeLoginEndpoint(s service.IService) endpoint.Endpoint{

	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(pb.LoginRequest)

		var token string

		if t, e := s.Login(req.Username, req.Password) ;e!= nil {
			token = t
			err = e
		}


		return pb.LoginReply{Token: token }, err
	}
}
