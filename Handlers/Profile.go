package Handlers
import "net/http"


func ProfileHomePage(w http.ResponseWriter,r *http.Request){
	info := []byte("profile Home Page")
	w.Write(info)
}
