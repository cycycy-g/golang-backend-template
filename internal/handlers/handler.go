package handlers

import (
	"your-project-name/internal/auth"
	"your-project-name/internal/handlers/core"
	"your-project-name/internal/handlers/user"
	"your-project-name/internal/store"
)

type Handler struct {
	User *user.Handler
}

func New(store store.Store, tokenMaker auth.Maker) *Handler {
	base := core.NewBaseHandler(store, tokenMaker)

	return &Handler{
		User: user.New(base),
	}
}
