package main

import (
	"log"
	"net/http"
	"proj/Autho"
	"proj/DB"
	"proj/MiddleWare"
	"proj/Routing"
)

var chain http.Handler
var UniversalLogger = Autho.NewGlobalLogger("SystemLogs")

func init() {
	Router := Routing.StartAllRouters()
	chain = MiddleWare.New(MiddleWare.LogRequest, MiddleWare.StandardHeaders).Then(Router)
	DB.StartConnection()

	defer UniversalLogger.CleanUp()
	UniversalLogger.Info.Output("Started Routers")
	UniversalLogger.Info.Output("Set up Database connection")
	DB.StartConnection()
}

func main() {
	UniversalLogger.Info.Output("running server on local host 8080")
	log.Fatal(http.ListenAndServe(":8080", chain))

}
