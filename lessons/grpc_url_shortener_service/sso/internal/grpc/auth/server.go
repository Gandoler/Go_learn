package auth

import ssov1 "grpc_url_shortener_service/protos/gen/go/sso"

type serverApi struct {
	ssov1.UnimplementedAuthServer
}
