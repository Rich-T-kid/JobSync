package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"proj/DB"
	"proj/MiddleWare"
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
	value := os.Getenv("EmailAPIKey")
	fmt.Println("email api  =" , value)
	fmt.Println("running server on local host 8080")
	log.Fatal(http.ListenAndServe(":8080", chain))

}
