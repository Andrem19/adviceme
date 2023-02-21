package gapi

import (
	"context"

	db "github.com/andrem19/adviceme/db/sqlc"
	"github.com/andrem19/adviceme/pb"
	"github.com/andrem19/adviceme/util"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server * Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserRsponse, error) {
	hashedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Unimplemented, "failed to hash password")
	}

	arg := db.CreateUserParams{
		Nickname: req.GetNickname(),
		Email: req.GetEmail(),
		Balance: 0,
		HashedPassword: hashedPassword,
		Resp: false,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "username already exist %s", err)
			}
		}
		return nil, status.Errorf(codes.Internal, "failed to create user %s", err)
	}

	rsp := &pb.CreateUserRsponse{
		User: convertUser(user),
	}

	return rsp, nil
}