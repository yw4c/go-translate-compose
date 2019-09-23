package service

import (
	"testing"
	"translate/P10User/src/lib/log"
	"translate/P10User/src/repository/mock"
)

func TestLogin(t *testing.T) {

	log.Setup()

	s:= NewUserService(&mock.UserRepositorySuccess{})
	token, err := s.Login("ben", "1234")

	if err != nil  {
		t.Fail()
	}
	t.Logf("token : %s", token)

}
