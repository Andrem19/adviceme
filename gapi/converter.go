package gapi

import (
	db "github.com/andrem19/adviceme/db/sqlc"
	"github.com/andrem19/adviceme/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.UserAccount) *pb.User {
	return &pb.User{
		Nickname:          user.Nickname,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.CreatedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}