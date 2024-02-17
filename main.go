package main

import ("fmt"
	"io"
//	"os"
//	"errors"
	"net/http")

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("root")
	io.WriteString(w, "This is my website!\n")
}


func main(){
	http.HandleFunc("/" , getRoot)
	fmt.Print("lets start")
	server := http.ListenAndServe(":8080",nil)
	if server != nil{
		fmt.Println("error occured" , server)
		
	}else{
		fmt.Println("starting on port 8080" , server)
		
	}

	
}
