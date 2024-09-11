package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"context"

	"github.com/iamjoseph331/miniserver/domain"
	"github.com/iamjoseph331/miniserver/log"
)

const (
	requestID  = "X-Kong-Request-ID"
	instanceID = "instance-id"
)

type Receptor interface {
	Heard(ctx context.Context, req *domain.ReceptorObject) (string, error)
	Thought(ctx context.Context, req *domain.ReceptorObject) (string, error)
	HeardVoice(ctx context.Context, req *domain.ReceptorObject) (string, error)
	Lookat(ctx context.Context, req *domain.ReceptorObject) (string, error)
}

type ServerViewService interface {
	SigninQuery(ctx context.Context, query string) (string, error)
}

var view ServerViewService

func Lookat(ctx *gin.Context) {
	/*
		request := &CreateClassRequest{
			Class: domain.FaceDefault512ClassName,
		}
		if err := ctx.ShouldBind(&request); err != nil {
			log.Logger.Error(log.ApplicationLog(ctx, "%v", err.Error()))
			handleShouldBindError(ctx, err)
			return
		}
		className, err := view.CreateClass512D(
			context.WithValue(ctx, requestID, ctx.Request.Header.Get(requestID)),
			&domain.CreateClassRequest{
				Class:      request.Class,
				InstanceID: request.InstanceId,
			},
		)
		if err != nil {
			handleErrors(ctx, err)
			return
		}
		resp := &CreateClassResponse{Class: className}
		ctx.JSON(http.StatusCreated, resp)*/
}

func Heard(ctx *gin.Context) {
	request := &domain.SigninQueryObject{
		UserID: "000",
	}
	if err := ctx.ShouldBind(&request); err != nil {
		log.Logger.Error(log.ApplicationLog(ctx, "%v", err.Error()))
		handleShouldBindError(ctx, err)
		return
	}
	response, err := view.SigninQuery(
		context.WithValue(ctx, requestID, ctx.Request.Header.Get(requestID)),
		request.Query,
	)
	if err != nil {
		log.Logger.Error(log.ApplicationLog(ctx, "%v", err.Error()))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"response": response,
	})
}

func Thought(ctx *gin.Context) {

}

func HeardVoice(ctx *gin.Context) {

}

func Healthy(ctx *gin.Context) {}

func handleShouldBindError(ctx *gin.Context, err error) {
	ctx.JSON(400, gin.H{
		"error": err.Error(),
	})
}
