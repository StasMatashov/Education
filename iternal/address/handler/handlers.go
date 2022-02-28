package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webapi/pkg/model"
)

func (h *Handler) CreateAddress(w http.ResponseWriter, r *http.Request) {
	var req model.AddressUser
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		return
	}

	_, err = h.service.CreateAddress(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Welcome to the club, body")
}

func (h *Handler) OutputAllAddresses(w http.ResponseWriter, r *http.Request) {
	allAddress := h.service.OutputAllAddresses()

	encoder := json.NewEncoder(w)
	err := encoder.Encode(allAddress)
	if err != nil {
		return
	}
}
