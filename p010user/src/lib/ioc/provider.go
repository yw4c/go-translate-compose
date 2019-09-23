package ioc

import (
	"github.com/spf13/viper"
	"translate/P10User/src/repository"
	"translate/P10User/src/repository/mock"
)

////////// register providers here ////////////

func userRepositoryProvider()repository.IUserRepository {

	if viper.GetString("APP_ENV") != "test" {
		return &repository.UserRepository{}
	}
	return &mock.UserRepositorySuccess{}
}








////////// test provider ////////////

type ITest interface {
	echo() string
}
type TestA struct {}


func (t *TestA) echo() string {
	return "success"
}

func testProvider() ITest {
	return &TestA{}
}

