package db

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"

	"github.com/spf13/viper"
)

type DefaultConn struct {
	Username string
	Password string
	Addr string
	Database string
	Driver string
	ConnectionQuery string
}

func (c *DefaultConn) Init()   {

	c.Username = viper.GetString("DB_USERNAME")
	c.Password = viper.GetString("DB_PASSWORD")
	c.Addr = viper.GetString("DB_HOST")+":"+viper.GetString("DB_PORT")
	c.Database = viper.GetString("DB_DATABASE")
	c.Driver = viper.GetString("DB_DRIVER")

	c.ConnectionQuery = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.Username, c.Password, c.Addr, c.Database)

	log.Println(c.ConnectionQuery)

}

func (c *DefaultConn) GetConnQuery() string{
	if c.ConnectionQuery == "" {
		panic("Conn has not set up, call Init() first")
	}
	return  c.ConnectionQuery
}
func (c *DefaultConn) GetDriver() string{
	if c.Driver == "" {
		panic("Conn has not set up, call Init() first")
	}
	return c.Driver
}
func (c *DefaultConn) GetDatabase() string {
	if c.Driver == "" {
		panic("Conn has not set up, call Init() first")
	}
	return c.Database
}