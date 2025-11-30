package sso

import (
	"context"
	ssov1 "github.com/Unpakenman/protos/gen/go/sso/rpc"
	"google.golang.org/grpc"
)

type authServer struct {
	ssov1.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
	ssov1.RegisterAuthServer(gRPC, &authServer{})
}

func (s *authServer) Login(ctx context.Context, req *ssov1.LoginRequest) (
	*ssov1.LoginResponse, error) {
	panic("implement me")
}
func (s *authServer) Register(ctx context.Context, req *ssov1.RegisterRequest) (
	*ssov1.RegisterResponse, error) {
	panic("implement me")
}

func (s *authServer) IsAdmin(ctx context.Context, req *ssov1.IsAdminRequest) (
	*ssov1.IsAdminResponse, error) {
	panic("implement me")
}
