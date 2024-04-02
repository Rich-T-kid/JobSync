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

func settingSubrouter(r  *mux.Router) *mux.Router {
	settingrouter := r.PathPrefix("/settings").Subrouter()
	return settingrouter
}
func localConnectionsSubrouter(r *mux.Router) *mux.Router {
	localConnectionsRouter := r.PathPrefix("/local-connections").Subrouter()
	// Define routes for the "/local-connections" subrouter
	return localConnectionsRouter
}

func chillSubrouter(r *mux.Router) *mux.Router {
	chillRouter := r.PathPrefix("/chill").Subrouter()
	// Define routes for the "/chill" subrouter
	return chillRouter
}

func jobsSubrouter(r *mux.Router) *mux.Router {
	jobsRouter := r.PathPrefix("/jobs").Subrouter()
	// Define routes for the "/jobs" subrouter
	return jobsRouter
}

func apiRouter(r *mux.Router) *mux.Router{
	apiRouter := r.PathPrefix("/Api").Subrouter()
	// Define Routers for /Api subrouter
	return apiRouter
}
