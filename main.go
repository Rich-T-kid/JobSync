package main

import (
	"log"
	"net/http"
	"proj/Autho"
	"proj/DB"
	"proj/MiddleWare"
<<<<<<< HEAD
=======
	"net/http"
>>>>>>> Cookies
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
	UniversalLogger.Info.Output("running server on local host 8080")
=======
	fmt.Println("running server on local host 8080")
>>>>>>> Cookies
	log.Fatal(http.ListenAndServe(":8080", chain))

}
