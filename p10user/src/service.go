package main

import "log"

type IService interface {
	Register()
	Login(username string ,password string) bool
	Show()
}

type service struct {

}

func (s *service) Register() {
	log.Print("this is register")
}

func (s *service) Login(username string ,password string) bool  {
	if username == "ben" && password == "123123" {
		log.Print("this is login")
		return true
	}
	return false
}

func (s *service) Show() {
	log.Print("this is show")
}