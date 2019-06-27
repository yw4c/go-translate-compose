package rpc

import "github.com/gin-gonic/gin"

type User struct {

}
func (user *User) Handler() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}