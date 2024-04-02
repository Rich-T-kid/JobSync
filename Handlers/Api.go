package Handlers


import ("fmt"	
	"net/http")


func FileUpload(w http.ResponseWriter , r *http.Request){
	switch r.Method{
case "GET":
	renderTemplate(w,"UploadFile.html",nil)
case "POST":
	fmt.Println("got post request")
default:
	fmt.Println("not allowed")

}}
