package service

import (
	"log"
	grpc_err "translate/P10User/src/error/grpc"
)



type IService interface {
	Register()
	Login(username string ,password string) (token string, err error)
	Show()
}

type UserService struct {

}

func (s *UserService) Register() {
	log.Print("this is register")
}

func (s *UserService) Login(username string ,password string) (token string, err error)  {

	if username == "ben" && password == "123123" {
		log.Print("this is login")
		token = "abc123"
		err = nil
	} else {
		err = grpc_err.NewGrpcError(grpc_err.LoginFail)
	}

	return token, err

}

func (s *UserService) Show() {
	log.Print("this is show")
}