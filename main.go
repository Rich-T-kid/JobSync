package main

import ("fmt"
	"io"
	"proj/DB"
//	"os"
//	"errors"
	"net/http")

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("root")
	io.WriteString(w, "This is my website!\n")
}


func main(){
	/*
	http.HandleFunc("/" , getRoot)

	fmt.Print("lets start")
	server := http.ListenAndServe(":8080",nil)
	if server != nil{
		fmt.Println("error occured" , server)
		
	}else{
		fmt.Println("starting on port 8080" , server)
		
	}*/
	db , err := DB.DBConnection()
	if err != nil{
		fmt.Println(err)}
	fmt.Println("db Connection ->", db)
	DB.GrabData(db)
	
}
