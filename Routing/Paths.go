package Routing
import ("github.com/gorilla/mux"
	"proj/Handlers"
)
var Router *mux.Router = globalRouter() 
// this will be the whole application 

func SetUpRouter() *mux.Router{ //this should be called in main package
	Router.HandleFunc("/",Handlers.LoginHandler).Methods("GET","POST")

	Router.HandleFunc("/logout",Handlers.LogOutHandler).Methods("GET","POST")
	Router.HandleFunc("/signup",Handlers.SignUpHandler).Methods("GET","POST")
	Router.HandleFunc("/forgotpassword",Handlers.ForgotPassHandler)
	Router.HandleFunc("/homepage",Handlers.HomePageHandler)
	Router.HandleFunc("/confiration",Handlers.ForgotHandler)
	Router.HandleFunc("/SignupConfirmation",Handlers.SignupConfirmationHandler).Methods("GET")
	Router.HandleFunc("/welcomeback",Handlers.WelcomeBackHandler).Methods("GET")
	Router.HandleFunc("/Active",Handlers.ActiveHandler).Methods("GET")
	return Router

}
