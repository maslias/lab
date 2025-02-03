package user

import (
	"context"

	"google.golang.org/grpc"

	userpb "github.com/maslias/chatroomer/pkg/common/proto/user"
)

type grpcHandler struct {
	userpb.UnimplementedUserServcieServer
}

func NewGrpcHandler(srv *grpc.Server) {
	h := &grpcHandler{}
	userpb.RegisterUserServcieServer(srv, h)
}

func (h *grpcHandler) GetUserById(
	ctx context.Context,
	payload *userpb.GetUserByIdRequest,
) (*userpb.GetUserResponse, error) {
	res := &userpb.GetUserResponse{
		Exist: true,
		User: &userpb.User{
			Id:       2,
			Username: "marciii",
			Email:    "marcel.liebreich@gmail.com",
		},
	}

	return res, nil
}
