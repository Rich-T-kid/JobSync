package MiddleWare

import "net/http"
/*
authenticate User by checking cookie first
if cookie doesnt exist return/is invalid user to the login page
otherwise alllow user to make their request 
grab the user from their cookie and make a AuthHead that will be the users actaully id/primary key in the databse
*/
func AuthMiddleWare(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path == "/" {
            return
        } else {
            _ , err := r.Cookie("SessionID")
            if err != nil { // invalid session cookie
                http.Redirect(w, r, "/", http.StatusSeeOther)
                return
            } else {
                // otherwise allow request to all other endpoints if there is a valid session id
                next.ServeHTTP(w, r)
            }
        }
    })
}

