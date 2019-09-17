package repository

import (
	"github.com/sirupsen/logrus"
	"translate/P10User/src/lib/db"
	"translate/P10User/src/orm"
)

type UserRepository struct {

}

func (r *UserRepository) FindOne(user *orm.User) *orm.User {
	conn := db.NewConn(&db.DefaultConn{})
	defer conn.Close()

	resultUser := &orm.User{}
	if conn.Where(user).First(&resultUser).RecordNotFound() {
		logrus.Debug("user not found")
		return nil
	}

	logrus.Debugf("found user id : %v",resultUser.Id)

	return resultUser

}
