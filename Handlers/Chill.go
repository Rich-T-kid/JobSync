package Handlers
import ("net/http"
"fmt"
)
func ChillHome(w http.ResponseWriter,r *http.Request){
	Inputlist := []string{"first test" , "second tesT" ,"third test"}
	UserNameCookie , err := r.Cookie("UserNameCookie")
	if err != nil{fmt.Println(err)}
	username := UserNameCookie.Value
 	data := struct {
            Username    string
            UserFriends []string
       		 }{
            Username:    username,
            UserFriends: Inputlist,}
		renderTemplate(w,"chill-chat/ChatApplication.html",data)
}
