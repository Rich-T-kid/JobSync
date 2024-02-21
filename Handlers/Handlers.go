package Handlers 
import ("fmt"  
	"net/http"
	"path/filepath"
	"html/template"
         "proj/DB"

)

func LoginHandler(w http.ResponseWriter,r *http.Request){
	fmt.Println(r.Method, "at", r.URL.Path)
	if r.Method == http.MethodGet{
	renderTemplate(w,"Login.html",nil)}
}
	/* else{ //post request
		err := r.ParseForm()
		if err != nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return }
	//Username := r.Form.Get("Username")
	//Password := r.Form.Get("Password")
	//if DB.ValidLogin(Username , Password){
	// 	http.Redirect(w,"/homepage",http.StatusSeeOther) // add a intermediate redirect
//}                                                              // that last like half a sec
								// and then redirects to homepage
	}
}
*/
//ForgotConfirm.html  forgotPassword.html  home.html  homepage.html  Signup.html

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
       /// http.Redirect(w, r, "/", http.StatusSeeOther)
    }
}

/*

Need to implement ogic to store user in database (autheticfication can come after)
write function to store user in the DB package and use here 
then after that validdation and adding user to db we can send them to the home page
work on assigning cookies after as well with the autheticaationa dn identifyation of users 

*/
//func LogHandler(w http.ResponseWriter, r *http.Request){}





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
