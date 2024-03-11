package Sessions 

import ("time"
	"errors"
	"github.com/google/uuid"
"net/http"	
	"github.com/gorilla/sessions")

var ( 
	ErrNoCookie = errors.New("http: named cookie not present")
	key = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func generateSessionID() string {
	return uuid.NewString()

}
func FormatedTime() string{

	currentTime := time.Now()

    	// Define the layout for the desired format
    	layout := "January 2 03:04 PM" // Month Day Hour:Minute AM/PM

    	// Format the current time using the layout
    	formattedTime := currentTime.Format(layout)
	return formattedTime


}


func CreateSessionCookie(username string, password string) *http.Cookie {

	sessionIDString := generateSessionID()
	cookie := http.Cookie{
   	     	 Name:     "SessionID",
       		 Value:    sessionIDString,
       		 Path:     "/",
           	 MaxAge:   3600,
        	 HttpOnly: true,
        	 SameSite: http.SameSiteLaxMode, }
	return &cookie
	
}
	

