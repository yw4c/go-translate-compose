package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type DefaultConn struct {

}

func (c *DefaultConn) Conn() *gorm.DB  {


	username := viper.GetString("DB_USERNAME")
	password := viper.GetString("DB_PASSWORD")
	addr := viper.GetString("DB_HOST")+":"+viper.GetString("DB_PORT")
	database := viper.GetString("DB_DATABASE")

	query := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, addr, database)
	db, err := gorm.Open("mysql", query)
	if err != nil {
		logrus.Panic(err)
	}
	return db
}
