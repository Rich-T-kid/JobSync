package main

import (
	"log"
	"net/http"
<<<<<<< HEAD
	"proj/Autho"
	"proj/DB"
	"proj/MiddleWare"
<<<<<<< HEAD
=======
	"net/http"
>>>>>>> Cookies
=======
	"os"
	"proj/DB"
	"proj/MiddleWare"
>>>>>>> ChatApp
	"proj/Routing"
)

var chain http.Handler
var UniversalLogger = Autho.NewGlobalLogger("SystemLogs")

func init() {
	Router := Routing.StartAllRouters()
	chain = MiddleWare.New(MiddleWare.LogRequest, MiddleWare.StandardHeaders).Then(Router)
<<<<<<< HEAD
	DB.StartConnection()

	defer UniversalLogger.CleanUp()
	UniversalLogger.Info.Output("Started Routers")
	UniversalLogger.Info.Output("Set up Database connection")
=======
	DB.StartConnection()	
	fmt.Println("set up routers and database connection")
>>>>>>> Cookies
}

func main() {
<<<<<<< HEAD
<<<<<<< HEAD
	UniversalLogger.Info.Output("running server on local host 8080")
=======
	fmt.Println("running server on local host 8080")
>>>>>>> Cookies
=======
	value := os.Getenv("EmailAPIKey")
	fmt.Println("email api  =" , value)
	fmt.Println("running server on local host 8080")
>>>>>>> ChatApp
	log.Fatal(http.ListenAndServe(":8080", chain))

}
