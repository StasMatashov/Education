package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webapi/pkg/model"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req model.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		return
	}

	_, err = h.service.CreateUser(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Welcome to hell")
}

func (h *Handler) SearchUserByIDs(w http.ResponseWriter, r *http.Request) {
	type UserCheck struct {
		Id []int `json:"id"`
	}

	var req UserCheck
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		return
	}

	sliceUser := h.service.SearchUserByIDs(req.Id)

	w.WriteHeader(http.StatusCreated)
	encoder := json.NewEncoder(w)
	err = encoder.Encode(sliceUser)
	if err != nil {
		return
	}
}

func (h *Handler) OutputAllUsers(w http.ResponseWriter, r *http.Request) {
	allUsers := h.service.OutputAllUsers()

	encoder := json.NewEncoder(w)
	err := encoder.Encode(allUsers)
	if err != nil {
		return
	}
}
