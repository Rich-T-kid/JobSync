package middleware

import (
    "fmt"
    "net/http"
    "time"
    "os"
)
func handErr(e error){
	if e != nil{
		panic(e)}
}

func writeToFile(data string, filepath string) error {
    f, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, 0644)
    handErr(err) 
    defer f.Close() // Ensure file is closed after use

    // Write data to the file
    _, err = fmt.Fprint(f, data)
    handErr(err)
    return nil // Return nil if the operation is successful
}

func LogRequest(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Get the current date and time
        now := time.Now()

        // Log the HTTP method, URL path, and current date and time
	formatedString := fmt.Sprintf("[%s] %s at %s\n", now.Format("2006-01-02 15:04:05"), r.Method, r.URL.Path)
	writeToFile(formatedString,"JobSyncLogs.txt")
        fmt.Printf(formatedString)

        // Call the next handler in the chain
        next.ServeHTTP(w, r)
    })
}

