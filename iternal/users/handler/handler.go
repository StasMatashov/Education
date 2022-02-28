package handler

import (
	"github.com/gorilla/mux"
	"net/http"
	"webapi/pkg/model"
)

type Service interface {
	CreateUser(user model.User) (int, error)
	SearchUserByIDs(IDs []int) []model.User
	OutputAllUsers() []model.User
}

type Handler struct {
	service Service
}

func NewHandler(svc Service) *Handler {
	return &Handler{
		service: svc,
	}
}

func (h *Handler) Init() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/users/createUser", h.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/users/searchUserByIDs", h.SearchUserByIDs).Methods(http.MethodGet)
	r.HandleFunc("/users/outputAllUsers", h.OutputAllUsers).Methods(http.MethodGet)

	return r
}
