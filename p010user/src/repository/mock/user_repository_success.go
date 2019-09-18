package mock

import (
	"translate/P10User/src/orm"
)

type UserRepositorySuccess struct {

}

func (r *UserRepositorySuccess) FindOne(user *orm.User) *orm.User {
	return &orm.User{
		Id:1,
		Username:"ben",
		Password:"hashed",
	}
}