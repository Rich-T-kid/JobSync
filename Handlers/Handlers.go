package Handlers 
import (//"time"
	"fmt"  
	"net/http"
	"path/filepath"
	"html/template"
         "proj/DB"

)




func LoginHandler(w http.ResponseWriter,r *http.Request){
	if r.Method == "GET"{
	renderTemplate(w,"Login.html",nil)
	}else{ //post request
		err := r.ParseForm()
		if err != nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return }
	Username := r.Form.Get("Username")
	Password := r.Form.Get("Password")
	fmt.Println(Username,Password)
	if DB.ValidLogin(Username , Password){
		http.Redirect(w,r,"/homepage",http.StatusSeeOther)
									 //	http.Redirect(w,"/homepage",http.StatusSeeOther) // add a intermediate redirect
		}                                                              // that last like half a sec
								// and then redirects to homepage
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
	renderTemplate(w,"ForgotConfirm.html",nil)
}

func ForgotPassHandler(w http.ResponseWriter, r *http.Request){
	renderTemplate(w,"forgotPassword.html.html",nil)
}
func HomePageHandler(w http.ResponseWriter, r *http.Request){
	renderTemplate(w,"homepage.html",nil)
}
func SignUpHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        // Render the signup form for GET requests
        renderTemplate(w, "Signup.html", nil)
    } else if r.Method == "POST" {
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
