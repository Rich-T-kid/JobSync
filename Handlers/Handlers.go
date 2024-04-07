package Handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"proj/DB"
	"proj/Emails"
	"proj/Sessions"
	"time"
)

const (
	Logfpath = "../JobSyncLogs.txt"
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
		_, Response := DB.RealLogin(Username, Password)
		if Response != nil {
			http.Error(w, "You must login fisrt", http.StatusUnauthorized)
		} else {
			Cookies, err := Sessions.GatherUserCookies(Username) //move both of these into the gatherUser cookies function. pass in username
			if err != nil {
				fmt.Fprint(w, err)
				return
			}
			cookie := Sessions.CreateSessionCookie(Username)
			Temp := DB.User{Name: Username, Password: Password, SessID: cookie.Value}
			DB.UserSlice = append(DB.UserSlice, Temp)
			if err := DB.StoreCookie(cookie, Username); err != nil {
				//this is only for the storagee cookie. other cookies already stored in database
				fmt.Fprintln(w, err)
			}
			Cookies = append(Cookies, cookie)
			for _, cookie := range Cookies {
				fmt.Println(cookie.Name, " cookie values  : ", cookie.Value)
				http.SetCookie(w, cookie)
			}
			http.SetCookie(w, cookie) // if done well wont need to do this by hand
			http.Redirect(w, r, "/homepage", http.StatusSeeOther)
			return
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func LogOutHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":

		renderTemplate(w, "LogOut.html", nil)

	case "POST":

		SessionCookie, err := r.Cookie("SessionID")

		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
		SessionCookie.Expires = time.Now().AddDate(0, 0, -1)
		http.SetCookie(w, SessionCookie)
		DB.RemoveUserSessionSlice(SessionCookie.Value)
		// need to remove all the cookies but mbaey it wont matter actaully
		// remove the session from the database
		DB.DeleteCookieSession(SessionCookie)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// move this to its ownfile once finished setting up the routers and subrouters.
// figure out how youll model user prefrences with cookies and perist this to the databse
func ActiveHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		renderTemplate(w, "ActiveSesion.html", DB.UserSlice)
	default:
		http.Error(w, "Method not allowed ", http.StatusMethodNotAllowed)
	}
}

func ForgotHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		renderTemplate(w, "ForgotConfirm.html", nil)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
func ForgotPassHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		renderTemplate(w, "forgotPassword.html.html", nil)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		UserNameCookie, err := r.Cookie("UserNameCookie")
		if err != nil {
			fmt.Print(err)
			return
		}
		username := UserNameCookie.Value
		data := struct {
			Username string
		}{
			Username: username}
		fmt.Println(data, data.Username)
		renderTemplate(w, "homepage.html", data)
	default:
		http.Error(w, "Method notallowed", http.StatusMethodNotAllowed)
	}
}

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
			if err.Error() == "User already exists" {
				http.Error(w, "Username  already exists.Try using a different Username", http.StatusConflict)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Need a way to validate that its a real email being passed.
		status, err := Emails.SendEmail(username, email, Emails.GenericTemplate, Emails.GenerticHtmlTemplate)
		fmt.Println(status, err)
		if err != nil {
			fmt.Println("error occured sending email, status err: ", status, err)
		}
		http.Redirect(w, r, "/SignupConfirmation", http.StatusSeeOther)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func SignupConfirmationHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "SignUpconfirmation.html", nil)
}

func LogHandler(w http.ResponseWriter, r *http.Request) {
	data, err := DB.TopLogs()
	if err != nil {
		fmt.Fprint(w, "Error had occured loading Logs")
		return
	}
	DataBytes := []byte(data)
	_, err = w.Write(DataBytes)
	if err != nil {
		http.Error(w, "Error Occured", http.StatusInternalServerError)
	}
}

func InvalidCredentials(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "InvalidCred.html", nil)
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	data, err := DB.AllActiveUsers()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
	fmt.Fprint(w, data)
	return

}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) { // this is for static html page rendering.
	staticDir := "static"
	// Get the absolute path by joining the current working directory with the relative path
	absStaticDir, err := filepath.Abs(staticDir)
	if err != nil {

	}
	tmplPath := filepath.Join(absStaticDir, tmpl)
	t, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(data)
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
