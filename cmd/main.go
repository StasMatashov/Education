package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	addressHandler "webapi/iternal/address/handler"
	addressService "webapi/iternal/address/service"
	userHandler "webapi/iternal/users/handler"
	userService "webapi/iternal/users/service"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Main page")
}

func testSearchId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println(id)
}

func testSearchString(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["s"]
	fmt.Println(id)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/mainPage", mainPage)
	router.HandleFunc("/testSearchId/{id:[0-9]+}", testSearchId)
	router.HandleFunc("/testSearchString/{s:[a,b]+}", testSearchString)

	userSvc := userService.NewService()
	userHandler := userHandler.NewHandler(userSvc)
	router.PathPrefix("/users").Handler(userHandler.Init())

	addressSvc := addressService.NewService()
	addressHandler := addressHandler.NewHandler(addressSvc)
	router.PathPrefix("/address").Handler(addressHandler.Init())

	fmt.Println("Server is listening...")
	http.ListenAndServe("localhost:8181", router)
}
