package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"context"

	"github.com/iamjoseph331/miniserver/domain"
	"github.com/iamjoseph331/miniserver/log"
)

const (
	auth  = "Authorization"
)

type Receptor interface {
	Signup(ctx context.Context, req *domain.ReceptorObject) (domain.SigninQueryResponse, error)
	GetUser(ctx context.Context, req *domain.ReceptorObject) (string, error)
	PatchUser(ctx context.Context, req *domain.ReceptorObject) (string, error)
	Close(ctx context.Context, req *domain.ReceptorObject) (string, error)
}

type ServerViewService interface {
	SigninQuery(ctx context.Context, user_id string, password string) (domain.SigninQueryResponse, error)
	GetUserQuery(ctx context.Context, user_id string) (domain.GetUserQueryResponse, error)
	PatchUserQuery(ctx context.Context, user_id string, nickname string, comment string) (domain.PatchUserResponse, error)
	Close(ctx context.Context) (domain.CloseResponse, error)
}

var view ServerViewService

type ReturnError struct {
	Message string `json:"message"`
	Cause string `json:"cause, omitempty"`
}

func Signup(ctx *gin.Context) {
	request := &domain.SigninQueryObject{
		UserId: "000",
		Password: "000",
	}
	if err := ctx.ShouldBind(&request); err != nil {
		log.Logger.Error(log.ApplicationLog(ctx, "%v", err.Error()))
		handleShouldBindError(ctx, err)
		return
	}
	response, err := view.SigninQuery(
		context.WithValue(ctx, auth, ctx.Request.Header.Get("Authorization")),
		request.UserId,
		request.Password,
	)
	if err != nil {
		log.Logger.Error(log.ApplicationLog(ctx, "%v", err.Error()))
		status_code := err.(domain.Err).StatusCode
		returnerror := ReturnError{
			Message: err.(domain.Err).Message,
			Cause: err.(domain.Err).Cause,
		}
		ctx.JSON(status_code, gin.H{
			"response": returnerror,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"response": response,
	})
}
// /users/{user_id}
func GetUser(ctx *gin.Context, user_id string) {
	response, err := view.GetUserQuery(
		context.WithValue(ctx, auth, ctx.Request.Header.Get("Authorization")),
		user_id,
	)
	if err != nil {
		log.Logger.Error(log.ApplicationLog(ctx, "%v", err.Error()))
		status_code := err.(domain.Err).StatusCode
		returnerror := ReturnError{
			Message: err.(domain.Err).Message,
			Cause: err.(domain.Err).Cause,
		}
		ctx.JSON(status_code, gin.H{
			"response": returnerror,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"response": response,
	})
}

//"/users/{user_id}"
func PatchUser(ctx *gin.Context, user_id string) {
	request := &domain.PatchUserObject{
		Nickname: "",
		Comment: "",
	}
	if err := ctx.ShouldBind(&request); err != nil {
		log.Logger.Error(log.ApplicationLog(ctx, "%v", err.Error()))
		handleShouldBindError(ctx, err)
		return
	}
	response, err := view.PatchUserQuery(
		context.WithValue(ctx, auth, ctx.Request.Header.Get("Authorization")),
		user_id,
		request.Nickname,
		request.Comment,
	)
	if err != nil {
		log.Logger.Error(log.ApplicationLog(ctx, "%v", err.Error()))
		status_code := err.(domain.Err).StatusCode
		returnerror := ReturnError{
			Message: err.(domain.Err).Message,
			Cause: err.(domain.Err).Cause,
		}
		ctx.JSON(status_code, gin.H{
			"response": returnerror,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"response": response,
	})
}

func Close(ctx *gin.Context) {
	response, err := view.Close(
		context.WithValue(ctx, auth, ctx.Request.Header.Get("Authorization")),
	)	
	if err != nil {
		log.Logger.Error(log.ApplicationLog(ctx, "%v", err.Error()))
		status_code := err.(domain.Err).StatusCode
		returnerror := ReturnError{
			Message: err.(domain.Err).Message,
			Cause: err.(domain.Err).Cause,
		}
		ctx.JSON(status_code, gin.H{
			"response": returnerror,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"response": response,
	})
}

func Healthy(ctx *gin.Context) {}

func handleShouldBindError(ctx *gin.Context, err error) {
	ctx.JSON(400, gin.H{
		"error": err.Error(),
	})
}
