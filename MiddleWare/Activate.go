package MiddleWare

import (
	"github.com/gorilla/mux"
	"net/http"
)

func ReturnChainedMiddleware(r *mux.Router) http.Handler {
	return New(LogRequest, StandardHeaders).Then(r)
}
