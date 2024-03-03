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
        _, err := r.Cookie("SessionID")
        if err != nil { // If session ID is expired or doesn't exist, redirect to "/InvalidCred"
            switch r.URL.Path {
            case "/", "/signup":
                // Valid cookie session or access to the login page, allow the request to proceed
                next.ServeHTTP(w, r)
            default:
                http.Redirect(w, r, "/InvalidCred", http.StatusSeeOther)
                return
            }
        } else {
            // Valid cookie session or access to the login page, allow the request to proceed
            next.ServeHTTP(w, r)
        }
    })
}
