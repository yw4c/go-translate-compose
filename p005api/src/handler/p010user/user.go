package p010user

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"

	pb "translate/P005api/src/pb/p010user"
)

const addr  = "p010user:6000"

func LoginHandler() gin.HandlerFunc {

	return func(context *gin.Context) {
		conn, err := grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		client := pb.NewUserClient(conn)

		request := &pb.LoginRequest{
			Username: context.PostForm("username"),
			Password: context.PostForm("password"),
		}
		log.Println(request)

		resp, err := client.Login(context, request)
		if err != nil {
			context.JSON(500, gin.H{
				"msg" : err.Error(),
			})
			return
		}

		log.Println(resp)
		if err := resp.GetError(); err != nil {
			context.JSON(400, gin.H{
				"code": err.GetCode(),
				"message": err.GetMessage(),
			})
			return
		}
		context.JSON(200, gin.H{
			"token": resp.GetToken(),
		})
		return
	}
}