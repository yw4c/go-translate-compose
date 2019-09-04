package grpc_cli

import (

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
)

func Requestor(addr string, handlerFunc gin.HandlerFunc)  {

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	handlerFunc(&gin.Context{})


}