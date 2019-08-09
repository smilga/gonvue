package router

import (
	"gonvue/pkg/config"
	"gonvue/pkg/http/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func New(config *config.Config) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", handler.AppHandler(config))

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	return r
}
