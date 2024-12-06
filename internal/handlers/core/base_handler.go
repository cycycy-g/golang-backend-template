package core

import (
	"your-project-name/internal/auth"
	"your-project-name/internal/store"
)

type BaseHandler struct {
	Store      store.Store
	TokenMaker auth.Maker
}

func NewBaseHandler(store store.Store, tokenMaker auth.Maker) *BaseHandler {
	return &BaseHandler{
		Store:      store,
		TokenMaker: tokenMaker,
	}
}
