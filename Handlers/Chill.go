package Handlers
import "net/http"


func ChillHome(w http.ResponseWriter,r *http.Request){
	info := []byte("chill  Home page")
	w.Write(info)
}
