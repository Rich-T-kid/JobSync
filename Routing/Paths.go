package Routing
import ("github.com/gorilla/mux"
//	"proj/Handlers"
)
var Router *mux.Router = globalRouter() 
// this will be the whole application 

func SetUpRouter() *mux.Router{ //this should be called in main package
	Router.HandleFunc("/")
	Router.HandleFunc("/signup")
	Router.HandleFunc("/login")
	Router.HandleFunc("/forgotpassword")
	Router.HandleFunc("/homepage")
	Router.HandleFunc("/confiration")
	
	return Router

}
