package db

import (
	"fmt"
	"github.com/spf13/viper"
)

type LocalConn struct {
	Username string
	Password string
	Addr string
	Database string
	Driver string
	ConnectionQuery string
}

func (c *LocalConn) Init()   {

	c.Username = viper.GetString("DB_USERNAME")
	c.Password = viper.GetString("DB_PASSWORD")
	c.Addr = "127.0.0.1:3306"
	c.Database = viper.GetString("DB_DATABASE")
	c.Driver = viper.GetString("DB_DRIVER")

	c.ConnectionQuery = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.Username, c.Password, c.Addr, c.Database)

}

func (c *LocalConn) GetConnQuery() string{
	if c.ConnectionQuery == "" {
		panic("Conn has not set up, call Init() first")
	}
	return  c.ConnectionQuery
}
func (c *LocalConn) GetDriver() string{
	if c.Driver == "" {
		panic("Conn has not set up, call Init() first")
	}
	return c.Driver
}
func (c *LocalConn) GetDatabase() string {
	if c.Driver == "" {
		panic("Conn has not set up, call Init() first")
	}
	return c.Database
}