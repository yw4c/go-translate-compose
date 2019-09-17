package db

import "github.com/jinzhu/gorm"

type IConn interface {
	Conn()*gorm.DB
}
