package core

import (
	"context"

	"github.com/iamjoseph331/miniserver/domain"
)

type ServerCore struct {
}

func NewServerCore() *ServerCore {
	return &ServerCore{}
}

func (h *ServerCore) SigninQuery(ctx context.Context, user_id string, password string) (domain.SigninQueryResponse, error) {
	// todo: implement fail cases
	domain.Database[user_id] = domain.UserPublic{
		UserId: user_id,
		Nickname: user_id,
	}
	domain.DatabasePassword[user_id] = password
	response := domain.SigninQueryResponse{
		Message: "Account successfully created",
		User: domain.UserPublic{
			UserId: user_id,
			Nickname: user_id,
		},
	}
	return response, nil
}

func (h *ServerCore) GetUserQuery(ctx context.Context, user_id string) (domain.GetUserQueryResponse, error) {
	user := domain.Database[user_id]
	response := domain.GetUserQueryResponse{
		Message: "User details by user_id",
		User: user,
	}
	return response, nil
}

func (h *ServerCore) PatchUserQuery(ctx context.Context, user_id string, nickname string, comment string) (domain.PatchUserResponse, error) {
	user := domain.Database[user_id]
	user.Nickname = nickname
	user.Comment = comment
	domain.Database[user_id] = user
	response := domain.PatchUserResponse{
		Message: "User successfully updated",
		Recipe: []domain.UserPublic{
			{
				Nickname: nickname,
				Comment: comment,
			},
		},
	}
	return response, nil
}

func (h *ServerCore) Close(ctx context.Context, user_id string) (domain.CloseResponse, error) {
	delete(domain.Database, user_id)
	delete(domain.DatabasePassword, user_id)
	response := domain.CloseResponse{
		Message: "Account and user successfully removed",
	}
	return response, nil
}