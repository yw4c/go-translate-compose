package src

import (
	"github.com/gin-gonic/gin"
	"translate/P10User/src/rpc"
)

func main() {
	r:= gin.Default()

	v1 := r.Group("/v1")
	{
		userRpc := rpc.User{};
		v1.POST("/user/login", userRpc.Handler())
	}

}
