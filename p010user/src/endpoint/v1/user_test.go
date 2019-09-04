package v1

import (
	"context"
	"google.golang.org/grpc"
	"testing"
	pb "translate/P10User/src/pb/p010user/v1"
	grpc_err "translate/P10User/src/error/grpc"
)

func TestLogin(t *testing.T) {
	conn, err := grpc.Dial("localhost:6010", grpc.WithInsecure())
	if err != nil {
		t.Logf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewUserClient(conn)

	// Correct Request
	{
		r, err := client.Login(context.Background(), &pb.LoginRequest{
			Username: "ben",
			Password: "123123",
		})
		if err != nil {
			t.Logf("could not greet: %v", err.Error())
			t.Fail()
		}
		t.Log(r)
		if r.GetError() != nil {
			t.Log(r.GetError().GetMessage())
			t.Fail()
		}
	}


	// Incorrect password
	{
		r, _ := client.Login(context.Background(), &pb.LoginRequest{
			Username: "ben",
			Password: "2222",
		})
		t.Log(r)
		if (r.Error.GetCode() != grpc_err.LoginFail) {
			t.Log(r.Error.GetCode())
			t.Fail()
		}
	}



}