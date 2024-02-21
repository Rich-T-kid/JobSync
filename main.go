package main

import ("fmt"
	"io"
//	"proj/DB"
	"proj/Routing"
//	"os"
//	"errors"
	"net/http")

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("root")
	io.WriteString(w, "This is my website!\n")
}
type data struct{
	name, password, email  string
	number int
}

func main(){

	Router := Routing.SetUpRouter()
	fmt.Println("running server on local host 9000")
	
	server := http.ListenAndServe(":9000",Router)
	if server != nil{
		fmt.Println("error :" , server)
	}
}

