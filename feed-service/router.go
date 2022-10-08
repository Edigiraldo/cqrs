package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func newRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/feeds", createdFeedHandler).Methods(http.MethodPost)

	return router
}
