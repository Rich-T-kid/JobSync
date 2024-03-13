package Handlers
import "net/http"


func JobsHomePage(w http.ResponseWriter,r *http.Request){
	info := []byte("jobs portion of site homepage")
	w.Write(info)
}
