package service

import (
	"crypto/md5"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	grpc_err "translate/P10User/src/error/grpc"
	"translate/P10User/src/orm"
	"translate/P10User/src/repository"
)


type UserService struct {

}

func (s *UserService) Register() {
	log.Print("this is register")
}

func (s *UserService) Login(username string ,password string) (token string, err error)  {

	data := []byte(password)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)

	logrus.Info("hashed psw" + md5str1)

	filter := &orm.User{
		Username:username,
		Password:md5str1,
	}

	repo := repository.UserRepository{}
	if user := repo.FindOne(filter); user == nil {
		err = grpc_err.NewGrpcError(grpc_err.LoginFail)
	} else {
		token = "123"
		err = nil
	}


	return token, err

}

func (s *UserService) Show() {
	log.Print("this is show")
}