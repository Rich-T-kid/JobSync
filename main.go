package main

import (
	"fmt"
	"log"
	"proj/MiddleWare"

	"proj/Routing"
	"net/http"
)

var chain http.Handler

func init(){
	Router := Routing.StartAllRouters()
	chain = MiddleWare.New(MiddleWare.LogRequest , MiddleWare.StandardHeaders).Then(Router)
	fmt.Print("Set Up Routers and middlware")
}


func main(){

	fmt.Println("running server on local host 9000")

	log.Fatal(http.ListenAndServe(":9000",chain))
	
}
