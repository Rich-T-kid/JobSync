package main

import ("fmt"
	"proj/MiddleWare"
//	"io"
//	"proj/DB"
	"proj/Routing"
//	"os"
//	"errors"
	"net/http")


func main(){

	Router := Routing.SetUpRouter()

	fmt.Println("running server on local host 9000")


	server := http.ListenAndServe(":9000",MiddleWare.StandardHeaders(MiddleWare.LogRequest(Router)))
	if server != nil{
		fmt.Println("error :" , server)
	}
}
