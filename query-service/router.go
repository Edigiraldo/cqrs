package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func newRouter() (router *mux.Router) {
	router = mux.NewRouter()

	router.HandleFunc("/feeds", ListFeedsHandler).Methods(http.MethodGet)
	router.HandleFunc("/search", SearchHandler).Methods(http.MethodGet)

	return router
}
