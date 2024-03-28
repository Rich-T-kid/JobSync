package Handlers
import "net/http"

func ChillHome(w http.ResponseWriter,r *http.Request){
	renderTemplate(w,"chill-chat/ChatApplication.html",nil)
}
