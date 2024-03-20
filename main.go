package main

import (
	"fmt"
	"log"
	"proj/DB"
	"proj/MiddleWare"

	"net/http"
	"proj/Routing"
)

var chain http.Handler

func init() {
	Router := Routing.StartAllRouters()
	chain = MiddleWare.New(MiddleWare.LogRequest, MiddleWare.StandardHeaders).Then(Router)
	fmt.Print("Set Up Routers and middlware.")
	fmt.Println("Sett up database connection")
	DB.DBConnection()
}


func main() {

	fmt.Println("running server on local host 8080")

	log.Fatal(http.ListenAndServe(":8080", chain))

}
