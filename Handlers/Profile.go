package Handlers

<<<<<<< HEAD
import "net/http"

func ProfileHomePage(w http.ResponseWriter, r *http.Request) {
	info := []byte("profile Home Page")
	w.Write(info)
=======
import (
	"fmt"
	"net/http"
)


func ProfileHomePage(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"Profile_page lol")
>>>>>>> ChatApp
}
