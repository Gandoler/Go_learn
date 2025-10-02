package auth

import (
	"context"

	ssov1 "github.com/JustSkiv/protos/gen/go/sso"
	"google.golang.org/grpc"
)

type serverApi struct {
	ssov1.UnimplementedAuthServer
}

func Register(grpcServer *grpc.Server) {
	ssov1.RegisterAuthServer(grpcServer, &serverApi{})
}

func (s *serverApi) Login(ctx context.Context, req *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {
	panic("implement me")
}

func (s *serverApi) Register(ctx context.Context, req *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error) {
	panic("implement me")
}

func (s *serverApi) IsAdmin(ctx context.Context, req *ssov1.IsAdminRequest) (*ssov1.IsAdminResponse, error) {
	panic("implement me")
}
