package Sessions 

import ("errors"
	"github.com/gorilla/sessions")

var ( 
	ErrNoCookie = errors.New("http: named cookie not present")
	key = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func GenerateSessionID(){
}
 
