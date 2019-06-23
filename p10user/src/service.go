package main

import (
	"errors"
	"log"
)

var ErrLoginInvalid = errors.New("username or password were wrong")

type IService interface {
	Register()
	Login(username string ,password string) (token string, err error)
	Show()
}

type service struct {

}

func (s *service) Register() {
	log.Print("this is register")
}

func (s *service) Login(username string ,password string) (token string, err error)  {

	if username == "ben" && password == "123123" {
		log.Print("this is login")
		token = "abc123"
		err = nil
	} else {
		err = ErrLoginInvalid
	}

	return token, err

}

func (s *service) Show() {
	log.Print("this is show")
}