package main

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type loginRequest struct {
	username string `json:"s"`
	password string `json:"s"`
	Err string `json:"err,omitempty"`
}

type LoginResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

func makeLoginEndpoint(s IService) endpoint.Endpoint{

	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		//req := request.(loginRequest)
		//s.Login(req)
		//response
	}
}
