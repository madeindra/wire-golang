package user

import "net/http"

type UserHandler interface {
	FetchByUsername() http.HandlerFunc
}

type handler struct {
	s UserService
}

func (h *handler) FetchByUsername() http.HandlerFunc {
	panic("implement this")
}
