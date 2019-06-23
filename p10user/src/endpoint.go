package main

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type loginRequest struct {
	username string
	password string
}

type LoginResponse struct {
	Token   string `json:"token"`
	Err error `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

func makeLoginEndpoint(s IService) endpoint.Endpoint{

	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(loginRequest)

		var token string

		if t, e := s.Login(req.username, req.password) ;e!= nil {
			token = t
			err = e
		}


		return LoginResponse{Token: token, Err: err }, err
	}
}
