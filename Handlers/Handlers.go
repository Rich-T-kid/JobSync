package Handlers 
import ("time"
	"fmt"  
	"net/http"
	"path/filepath"
	"html/template"
         "proj/DB"
	 "proj/Sessions"

)




func LoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		renderTemplate(w, "Login.html", nil)
	case "POST":
		// Post request
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		Username := r.Form.Get("Username")
		Password := r.Form.Get("Password")
		if DB.ValidLogin(Username, Password) {
			cookie := Sessions.CreateSessionCookie(Username,Password)
			http.SetCookie(w,cookie)

			http.Redirect(w, r, "/homepage", http.StatusSeeOther)
			return
		}
		// Invalid login, handle appropriately
		// For example, render a login error message
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func LogOutHandler(w http.ResponseWriter , r *http.Request){
	switch r.Method{
	case "GET":
		_ , err := r.Cookie("SessionID")
    		if err != nil {
        		// If the error is due to the cookie missing, inform the user
        	if err == http.ErrNoCookie {
            		http.Error(w, "Session cookie is missing", http.StatusUnauthorized)
            		return 	}
        	http.Error(w, err.Error(), http.StatusInternalServerError)
        		return   }  
   	 	// If the cookie is found, render the logout page
   		renderTemplate(w, "LogOut.html", nil)
		
	case "POST":
		SessionCookie , err := r.Cookie("SessionID")
		if err != nil{
		http.Redirect(w,r,"/",http.StatusSeeOther)		
		}
		SessionCookie.Expires = time.Now().AddDate(0,0,-1)
		http.SetCookie(w,SessionCookie)
		http.Redirect(w,r,"/",http.StatusSeeOther)
	default:
		http.Error(w,"Method not allowed",http.StatusMethodNotAllowed)
	}
}
// dud. Work on getting this to work but not imporant.
func WelcomeBackHandler(w http.ResponseWriter , r *http.Request){ // later on input the users name in here so it can be  passed to the template	
	/*renderTemplate(w,"Welcomeback.html",nil) // pass in the users name later to be passed into the template here 
	time.Sleep(2500 * time.Millisecond)
	http.Redirect(w,r,"/homepage",http.StatusSeeOther)
*/
}

func ForgotHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method{
	case "GET":
		renderTemplate(w,"ForgotConfirm.html",nil)
		default:
			http.Error(w,"Method not allowed",http.StatusMethodNotAllowed)
		}
	}
func ForgotPassHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method{
	case "GET":
	renderTemplate(w,"forgotPassword.html.html",nil)
	default:
		http.Error(w,"Method not allowed", http.StatusMethodNotAllowed)
}}
func HomePageHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method{
	case "GET":
		renderTemplate(w,"homepage.html",nil)
	default:
		http.Error(w,"Method notallowed" , http.StatusMethodNotAllowed)
	}}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Render the signup form for GET requests
		renderTemplate(w, "Signup.html", nil)
	case "POST":
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Get form values
		email := r.Form.Get("email")
		username := r.Form.Get("Username")
		password := r.Form.Get("psw")
		phoneNumber := r.Form.Get("PhoneNumber")

		// Insert the user into the database (assuming DB.InputUser is correct)
		err = DB.InputUser(username, password, email, phoneNumber)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Redirect to the root URL after successful form submission
		http.Redirect(w, r, "/SignupConfirmation", http.StatusSeeOther)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}


func SignupConfirmationHandler(w http.ResponseWriter , r *http.Request){
	renderTemplate(w,"SignUpconfirmation.html",nil)
}

func LogHandler(w http.ResponseWriter, r *http.Request){}





func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) { // this is for static html page rendering. 
	staticDir := "static"
	// Get the absolute path by joining the current working directory with the relative path
	absStaticDir, err := filepath.Abs(staticDir)
	if err != nil{
	
	}
	fmt.Println("abs statoc dir: ", absStaticDir)

	tmplPath := filepath.Join(absStaticDir,tmpl)
	t, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
