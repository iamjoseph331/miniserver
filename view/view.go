package view

import "context"

type CoreService interface {
	SigninQuery(ctx context.Context, query string) (string, error)
}

type ServerView struct {
	core CoreService
}

func NewServerView(core CoreService) *HayateView {
	return &HayateView{
		core: core,
	}
}

func (v *ServerView) SigninQuery(ctx context.Context, query string) (string, error) {
	return v.core.SigninQuery(ctx, query)
}
