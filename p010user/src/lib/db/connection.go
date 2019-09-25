package db

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func NewConn(conn IConn) *gorm.DB {
	conn.Init()
	query := conn.GetConnQuery()
	logrus.Infof("connect %s", query)
	db, err := gorm.Open(conn.GetDriver(), query)

	if err != nil {
		logrus.Panic(err)
	}
	return db
}