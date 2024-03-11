package Emails

import (
	"fmt"
	"os"
	"proj/Sessions"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

const APiKey = "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"
const GenericTemplate = "We're thrilled to have you on board with JobSynce! Get ready to dive deeper into exciting career opportunities and vibrant developer communities. We can't wait to see where this journey takes you"
const GenerticHtmlTemplate = "<strong>We're thrilled to have you on board with JobSynce! Get ready to dive deeper into exciting career opportunities and vibrant developer communities. We can't wait to see where this journey takes you.</strong><br/><i>This is an automated email - Please do not respond.</i>"

func SendEmail(Username, UserGmail, plaintext, html string) (int, error) {
    from := mail.NewEmail("JobsSync", "jobsyncrichard@gmail.com")
    subject := "You've Successfully signed up for JobSync"
    to := mail.NewEmail("Valued Customer", UserGmail)
    plainTextContent := "Hello " + Username + "\n\n" + plaintext
    htmlContent := "<p>Hello " + Username + "</p><p>" + "" + "</p>" + html
    message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
    client := sendgrid.NewSendClient(APiKey)
    response, err := client.Send(message)
    if err != nil {
        return response.StatusCode, err
    } else {
        return response.StatusCode, nil
    }
}

func ReportError(status int ,er error ) error {
    f, err := os.OpenFile("EmailErrorLogs.txt", os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil{
    	return err} 
    defer f.Close() // Ensure file is closed after use

    // Write data to the file
    _, err = fmt.Fprint(f, status,er)
    if err !=nil{
	    return err}
    return nil // Return nil if the operation is successful
}

func ReportSuccess(email  , path string ) error {
    f, err := os.OpenFile("SuccessEmailLog.txt", os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil{
    	return err} 
    defer f.Close() // Ensure file is closed after use
    currentTime := Sessions.FormatedTime()
    message := fmt.Sprintf("Sent email to %s at %s at this time %s\n", email, path, currentTime)
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
/*
func main() {
    fmt.Println("starting")
    statuscode, err := SendEmail("Richard", "rbb98@scarletmail.rutgers.edu", GenericTemplate, GenerticHtmlTemplate)
    fmt.Println(statuscode, err)
    fmt.Println("ending here ")
}*/
