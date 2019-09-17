package db

import "github.com/jinzhu/gorm"

func NewConn(conn IConn) *gorm.DB {
	return conn.Conn()
}