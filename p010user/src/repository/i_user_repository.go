package repository

import "translate/P10User/src/orm"


type IUserRepository interface {
	FindOne(user *orm.User) *orm.User
}