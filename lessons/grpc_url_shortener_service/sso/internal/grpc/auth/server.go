package auth

import (
	"context"

	ssov1 "github.com/JustSkiv/protos/gen/go/sso"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const emptyValue = 0

type Auth interface {
	Login(ctx context.Context, username string, password string, appId int) (token string, err error)
	Register(ctx context.Context, username string, password string) (userID int, err error)
	IsAdmin(ctx context.Context, userID int) (bool, error)
}

type serverApi struct {
	ssov1.UnimplementedAuthServer
	auth Auth
}

func Register(grpcServer *grpc.Server) {
	ssov1.RegisterAuthServer(grpcServer, &serverApi{auth: auth})
}

func (s *serverApi) Login(ctx context.Context, req *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {
	if req.GetEmail() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "email required")
	}
	if req.GetPassword() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "password required")
	}

	if req.GetAppId() == emptyValue {
		return nil, status.Errorf(codes.InvalidArgument, "app id required")
	}

	return &ssov1.LoginResponse{
		Token: "",
	}, nil
}

func (s *serverApi) Register(ctx context.Context, req *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error) {
	panic("implement me")
}

func (s *serverApi) IsAdmin(ctx context.Context, req *ssov1.IsAdminRequest) (*ssov1.IsAdminResponse, error) {
	panic("implement me")
}
