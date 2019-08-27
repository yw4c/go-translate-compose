package endpoint

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"testing"
	pb "translate/P10User/src/pb/user"
)

func TestLogin(t *testing.T) {
	conn, err := grpc.Dial("localhost:6010", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewUserClient(conn)
	r, err := client.Login(context.Background(), &pb.LoginRequest{
		Username: "ben",
		Password: "123123",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	fmt.Println(r)
	t.Log("hi")
	t.Log(r)
}