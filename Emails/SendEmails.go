package Emails

import (
	"os"
	"path/filepath"
	"proj/Sessions"
	"errors"
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var APiKey = os.Getenv("EmailAPIKey")
const GenericTemplate = "We're thrilled to have you on board with JobSynce! Get ready to dive deeper into exciting career opportunities and vibrant developer communities. We can't wait to see where this journey takes you"
const GenerticHtmlTemplate = "<strong>We're thrilled to have you on board with JobSynce! Get ready to dive deeper into exciting career opportunities and vibrant developer communities. We can't wait to see where this journey takes you.</strong><br/><i>This is an automated email - Please do not respond.</i>"

func SendEmail(Username, UserGmail, plaintext, html string) (int, error) {
    from := mail.NewEmail("JobsSync", "jobsyncrichard@gmail.com")
    subject := "You've Successfully signed up for JobSync"
    to := mail.NewEmail("Valued Customer", UserGmail)
    plainTextContent := "Hello " + Username + "\n\n" + plaintext
    htmlContent := "<p>Hello " + Username + "</p><p>" + "" + "</p>" + html
    fmt.Println("the api key that was placed in here is :", APiKey)
    message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
    client := sendgrid.NewSendClient(APiKey)
    response, err := client.Send(message)
    if err != nil{
	    reportError(-1,err)
	    return -1,err
    }
	if response.StatusCode < 200 ||response.StatusCode > 299  {
        // Handle non-200 status code
        reportError(response.StatusCode, errors.New("non-200 status code returned"))
        return response.StatusCode, errors.New("non-200 status code returned")
    }

    // If everything is successful, report success
    reportSuccess(UserGmail)
    return response.StatusCode, nil
}    

func reportError(status int ,er error ) error {
	curpwd := currentwd()
	emailsDir := filepath.Join(curpwd, "Emails")

    // Construct the path to the success email log file within the "Emails" directory
    filePath := filepath.Join(emailsDir, "ErrorEmailLog.txt")
    f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil{
    	return err} 
    defer f.Close() // Ensure file is closed after use
    currentTime := Sessions.FormatedTime()
    _, err = fmt.Fprint(f, "Error occurred sending email at ", currentTime, " with the status code of ", status, " and error of ", er,"\n" )
    if err !=nil {
	    return err}
    return nil 
}
func reportSuccess(email string ) error {
	curpwd := currentwd()
	emailsDir := filepath.Join(curpwd, "Emails")

    // Construct the path to the success email log file within the "Emails" directory
    filePath := filepath.Join(emailsDir, "SuccessEmailLog.txt")
    f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil{
    	return err} 
    defer f.Close() 
    currentTime := Sessions.FormatedTime()
    message := fmt.Sprintf("Sent email to %s at %s at this time %s\n", email, currentTime, currentTime)
    // Write data to the file
    _, err = fmt.Fprint(f,message)
    if err !=nil{
	    return err}
    return nil // Return nil if the operation is successful
}

func FakeSendEmail(Username, UserGmail, plaintext, html string) (int, error) {
	fmt.Println("JUst Sent Email. Not fr tho dont want to waste API REquest. Line 60/Emails")	
	return 0,nil
}

func currentwd() string {
    // Get the current working directory
    cwd, err := os.Getwd()
    if err != nil {
        fmt.Println("Error getting curent working dir for emails:", err)
        return "" 
    }

     return cwd
}
