package view

import "context"
import "encoding/base64"
import "strings"

import "github.com/iamjoseph331/miniserver/domain"

type CoreService interface {
	SigninQuery(ctx context.Context, user_id string, password string) (domain.SigninQueryResponse, error)
	GetUserQuery(ctx context.Context, user_id string) (domain.GetUserQueryResponse, error)
	PatchUserQuery(ctx context.Context, user_id string, nickname string, comment string) (domain.PatchUserResponse, error)
	Close(ctx context.Context, user_id string) (domain.CloseResponse, error)
}

type ServerView struct {
	core CoreService
}

func NewServerView(core CoreService) *ServerView {
	return &ServerView{
		core: core,
	}
}

func isAlphanumeric(s string) bool {
	for _, c := range s {
		if !('a' <= c && c <= 'z') && !('A' <= c && c <= 'Z') && !('0' <= c && c <= '9') {
			return false
		}
	}
	return true
}

func isAlphanumericSpecial(s string) bool {
	for _, c := range s {
		if !('a' <= c && c <= 'z') && !('A' <= c && c <= 'Z') && !('0' <= c && c <= '9') && !('!' <= c && c <= '/') && !(':' <= c && c <= '@') && !('[' <= c && c <= '`') && !('{' <= c && c <= '~') {
			return false
		}
	}
	return true
}

func Authorization(ctx context.Context) (string, string) {
	// base64 decode -> split username and password by ':' -> return username and password
	auth := ctx.Value("Authorization").(string)
	decoded, err := base64.StdEncoding.DecodeString(auth)
	if err != nil {
		return "", ""
	}
	credentials := string(decoded)
	//split by ':'
	usernamepassword := strings.Split(credentials, ":")
	// check length is 2
	if len(usernamepassword) != 2 {
		return "", ""
	}
	return usernamepassword[0], usernamepassword[1]
}

func (v *ServerView) SigninQuery(ctx context.Context, user_id string, password string) (domain.SigninQueryResponse, error) {
	// handle potential errors
	// 400 require user_id and password
	if user_id == "" || password == "" {
		return domain.SigninQueryResponse{}, domain.Err{
			StatusCode: 400,
			Message: "Account creation failed",
			Cause: "require user_id and password",
		}
	}
	// user_id must be 6 ~ 20 characters, password must be 8 ~ 20 characters
	if len(user_id) < 6 || len(user_id) > 20 || len(password) < 8 || len(password) > 20 {
		return domain.SigninQueryResponse{}, domain.Err{
			StatusCode: 400,
			Message: "Account creation failed",
			Cause: "user_id must be 6 ~ 20 characters, password must be 8 ~ 20 characters",
		}
	}
	// user_id must be consist of only alphabets and numbers, password must be consist of only alphabets, numbers, and special characters
	if !isAlphanumeric(user_id) || !isAlphanumericSpecial(password) {
		return domain.SigninQueryResponse{}, domain.Err{
			StatusCode: 400,
			Message: "Account creation failed",
			Cause: "user_id must be consist of only alphabets and numbers, password must be consist of only alphabets, numbers, and special characters",
		}
	}
	// 400 already same user_id is used
	if _, ok := domain.Database[user_id]; ok {
		return domain.SigninQueryResponse{}, domain.Err{
			StatusCode: 400,
			Message: "Account creation failed",
			Cause: "already same user_id is used",
		}
	}
	return v.core.SigninQuery(ctx, user_id, password)
}

func (v *ServerView) GetUserQuery(ctx context.Context, user_id string) (domain.GetUserQueryResponse, error) {
	userauth, passauth := Authorization(ctx)
	// 404 can not find user_id
	if _, ok := domain.Database[user_id]; !ok {
		return domain.GetUserQueryResponse{}, domain.Err{
			StatusCode: 404,
			Message: "No user found",
		}
	}
	// 401 unauthorized
	if userauth != user_id || passauth != domain.DatabasePassword[user_id] {
		return domain.GetUserQueryResponse{}, domain.Err{
			StatusCode: 401,
			Message: "Unauthorized",
			Cause: "Authernication failed",
		}
	}
	return v.core.GetUserQuery(ctx, user_id)
}

func (v *ServerView) PatchUserQuery(ctx context.Context, user_id string, nickname string, comment string) (domain.PatchUserResponse, error) {
	userauth, passauth := Authorization(ctx)
	// 404 can not find user_id
	if _, ok := domain.Database[user_id]; !ok {
		return domain.PatchUserResponse{}, domain.Err{
			StatusCode: 404,
			Message: "No user found",
		}
	}
	// 400 require nickname or comment
	if nickname == "" && comment == "" {
		return domain.PatchUserResponse{}, domain.Err{
			StatusCode: 400,
			Message: "User updation failed",
			Cause: "require nickname or comment",
		}
	}
	// 403 no permission to update
	if userauth != user_id {
		return domain.PatchUserResponse{}, domain.Err{
			StatusCode: 403,
			Message: "no permission to update",
		}
	}

	// 401 unauthorized
	if passauth != domain.DatabasePassword[user_id] {
		return domain.PatchUserResponse{}, domain.Err{
			StatusCode: 401,
			Message: "Authernication failed",
		}
	}

	return v.core.PatchUserQuery(ctx, user_id, nickname, comment)
}

func (v *ServerView) Close(ctx context.Context) (domain.CloseResponse, error) {
	userauth, passauth := Authorization(ctx)
	// 401 Authernication failed: username and password are not matched
	if passauth != domain.DatabasePassword[userauth] {
		return domain.CloseResponse{}, domain.Err{
			StatusCode: 401,
			Message: "Authernication failed",
		}
	}
	return v.core.Close(ctx, userauth)
}