package Routing

import (
	"github.com/gorilla/mux"
	//"fmt"
	"proj/ChatServer"	
	"proj/Handlers"
)

var (
	Router           *mux.Router = globalRouter()
	setting          *mux.Router = settingSubrouter(Router)
	localconnections *mux.Router = localConnectionsSubrouter(Router)
	chill            *mux.Router = chillSubrouter(Router)
	jobs             *mux.Router = jobsSubrouter(Router)
)

// this will be the whole application

// all handlers here are defined in Handlers.go
func SetUpRouter() *mux.Router { //this should be called in main package
	Router.HandleFunc("/", Handlers.LoginHandler).Methods("GET", "POST")


	Router.HandleFunc("/Test", Handlers.TestHandler).Methods("GET")

	Router.HandleFunc("/logout", Handlers.LogOutHandler).Methods("GET", "POST")
	Router.HandleFunc("/signup", Handlers.SignUpHandler).Methods("GET", "POST")
	Router.HandleFunc("/forgotpassword", Handlers.ForgotPassHandler)
	Router.HandleFunc("/homepage", Handlers.HomePageHandler)
	Router.HandleFunc("/confiration", Handlers.ForgotHandler)
	Router.HandleFunc("/SignupConfirmation", Handlers.SignupConfirmationHandler).Methods("GET")
	Router.HandleFunc("/Active", Handlers.ActiveHandler).Methods("GET")
	Router.HandleFunc("/InvalidCred", Handlers.InvalidCredentials).Methods("GET")
	Router.HandleFunc("/Logs", Handlers.LogHandler).Methods("GET")

	return Router

}

// Prefix /settings
func SetUpSettings() {
	setting.HandleFunc("", Handlers.SettingsHome).Methods("GET")
	setting.HandleFunc("/Profile", Handlers.ProfileHomePage)
}

// Prefix /local-connections
func SetUpLocalConnections() {
	localconnections.HandleFunc("", Handlers.LocalConnectionHome).Methods("GET")
}

// Prefix /chill
func SetUpChill() {
	chill.HandleFunc("", Handlers.ChillHome)
	chill.HandleFunc("/wsConnection", ChatServer.WebSocketConnection)
}

// Prefix /jobs
func SetUpJobs() {
	jobs.HandleFunc("", Handlers.JobsHomePage)
}

// Start main router and subrouters
func StartAllRouters() *mux.Router {
	mainRouter := SetUpRouter()
	SetUpSettings()
	SetUpJobs()
	SetUpChill()
	SetUpRouter()
	SetUpLocalConnections()
	return mainRouter
}
