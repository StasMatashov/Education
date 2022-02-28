package handler

import (
	"github.com/gorilla/mux"
	"net/http"
	"webapi/pkg/model"
)

type Service interface {
	CreateAddress(address model.AddressUser) (int, error)
	OutputAllAddresses() []model.AddressUser
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

	r.HandleFunc("/address/createAddress", h.CreateAddress).Methods(http.MethodPost)
	r.HandleFunc("/address/outputAllAddresses", h.OutputAllAddresses).Methods(http.MethodGet)

	return r
}
