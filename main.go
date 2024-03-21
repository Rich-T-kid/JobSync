package main

import (
	"fmt"
	"log"
	"proj/DB"
	"proj/MiddleWare"
	"os"
	"net/http"
	"proj/Routing"
)

var chain http.Handler

func init() {
	Router := Routing.StartAllRouters()
	chain = MiddleWare.New(MiddleWare.LogRequest, MiddleWare.StandardHeaders).Then(Router)
	DB.StartConnection()
	
	fmt.Println("set up routers")
	fmt.Println("Sett up database connection")
}


func main() {
	//os.Setenv("EmailAPIKey", "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")
	var APiKey = os.Getenv("EmailAPIKey")
	fmt.Println("running server on local host 8080")
	fmt.Println(APiKey)
	log.Fatal(http.ListenAndServe(":8080", chain))

}
