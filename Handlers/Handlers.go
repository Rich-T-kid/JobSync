package Handlers 
import ("fmt"  
	"net/http"
	"path/filepath"
	"html/template"


)

func HomeHandler(w http.ResponseWriter,r *http.Request){
	renderTemplate(w,"home.html",nil)
}

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
func SignUpHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
	renderTemplate(w,"Signup.html",nil)
	}else{
		err := r.ParseForm()
		if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return		}
	email := r.Form.Get("email")
	username := r.Form.Get("Username")
	password := r.Form.Get("psw")
	passwordr := r.Form.Get("psw-repeat")
	PhoneNumber := r.Form.Get("PhoneNumber")
	fmt.Fprint(w,email,username,password,passwordr,PhoneNumber)
}}

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
