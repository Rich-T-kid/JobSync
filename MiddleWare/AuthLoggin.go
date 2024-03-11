package MiddleWare

import ("net/http"
	"time"

)
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
// StandardHeaders is a middleware function to set standard response headers.
func StandardHeaders(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        time := FormatedHeaderDate()
        // Set response headers
        w.Header().Set("Last-Request", time)
        w.Header().Set("Location", r.RemoteAddr) //  can use API to grab user's real location

        // Call the next handler in the chain
        next.ServeHTTP(w, r)
    })
}

func formatDatehelper(t time.Time) string {
    // Format the date using the desired layout
    layout := "Mon, 02 Jan 2006 15:04:05 MST"
    formattedDate := t.Format(layout)

    return formattedDate
}

func FormatedHeaderDate()string{
	CurrentUnformatedTime := time.Now()
	RealTime := formatDatehelper(CurrentUnformatedTime)
	return RealTime
}
