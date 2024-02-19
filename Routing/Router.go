package Routing

import (// "net/http" 
	"github.com/gorilla/mux"
)
var (
    GlobalRouter *mux.Router = nil
)

func globalRouter() *mux.Router {
    if GlobalRouter == nil {
        GlobalRouter = mux.NewRouter()
    }
    return GlobalRouter
}


