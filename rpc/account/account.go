// Code generated by goctl. DO NOT EDIT!
// Source: account.proto

package account

import (
	"context"

	"github.com/xh-polaris/account-svc/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	SendVerifyCodeReq  = pb.SendVerifyCodeReq
	SendVerifyCodeResp = pb.SendVerifyCodeResp
	SetPasswordReq     = pb.SetPasswordReq
	SetPasswordResp    = pb.SetPasswordResp
	SignInReq          = pb.SignInReq
	SignInResp         = pb.SignInResp

	Account interface {
		SignIn(ctx context.Context, in *SignInReq, opts ...grpc.CallOption) (*SignInResp, error)
		SetPassword(ctx context.Context, in *SetPasswordReq, opts ...grpc.CallOption) (*SetPasswordResp, error)
		SendVerifyCode(ctx context.Context, in *SendVerifyCodeReq, opts ...grpc.CallOption) (*SendVerifyCodeResp, error)
	}

	defaultAccount struct {
		cli zrpc.Client
	}
)

func NewAccount(cli zrpc.Client) Account {
	return &defaultAccount{
		cli: cli,
	}
}

func (m *defaultAccount) SignIn(ctx context.Context, in *SignInReq, opts ...grpc.CallOption) (*SignInResp, error) {
	client := pb.NewAccountClient(m.cli.Conn())
	return client.SignIn(ctx, in, opts...)
}

func (m *defaultAccount) SetPassword(ctx context.Context, in *SetPasswordReq, opts ...grpc.CallOption) (*SetPasswordResp, error) {
	client := pb.NewAccountClient(m.cli.Conn())
	return client.SetPassword(ctx, in, opts...)
}

func (m *defaultAccount) SendVerifyCode(ctx context.Context, in *SendVerifyCodeReq, opts ...grpc.CallOption) (*SendVerifyCodeResp, error) {
	client := pb.NewAccountClient(m.cli.Conn())
	return client.SendVerifyCode(ctx, in, opts...)
}
