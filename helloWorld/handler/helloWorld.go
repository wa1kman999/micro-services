package handler

import (
	"context"
	pb "github.com/wa1kman999/helloWorld/proto"
)

type UserLogin struct{}

type traceIDKey struct{}

func (u *UserLogin) Login(ctx context.Context, in *pb.UserLoginRequest, out *pb.UserResponse) error {
	account := in.Account
	password := in.Password
	// 判断password
	if account == "xiaojiang" && password == "1234" {
		out.Msg = "ok"
		out.Code = 666
		out.UserName = account
	} else {
		out.Msg = "error"
		out.Code = 333
		out.UserName = "没有此人"
	}
	return nil
}
