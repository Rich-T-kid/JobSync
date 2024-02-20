package Routing
import ("github.com/gorilla/mux"
	"proj/Handlers"
)
var Router *mux.Router = globalRouter() 
// this will be the whole application 

func SetUpRouter() *mux.Router{ //this should be called in main package
	Router.HandleFunc("/",Handlers.HomeHandler)

	Router.HandleFunc("/signup",Handlers.SignUpHandler).Methods("GET","POST")
	Router.HandleFunc("/forgotpassword",Handlers.ForgotPassHandler)
	Router.HandleFunc("/homepage",Handlers.HomePageHandler)
	Router.HandleFunc("/confiration",Handlers.ForgotHandler)

	return Router

}
