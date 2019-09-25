package db

import (
	"fmt"
	"github.com/spf13/viper"
)

type LocalConn struct {
	conn
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

