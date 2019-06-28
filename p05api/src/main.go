package src

import (
	"github.com/gin-gonic/gin"
	"translate/P10User/src/rpc-cli"
)

func main() {
	r:= gin.Default()

	v1 := r.Group("/v1")
	{
		//userRpc := rpc-cli.User{};
		//v1.POST("/user/login", userRpc.Handler())
	}

}
