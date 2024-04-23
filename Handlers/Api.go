package Handlers

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func FileUpload(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		renderTemplate(w, "UploadFile.html", nil)
	case "POST":
		fmt.Println("got post request")
	default:
		fmt.Println("not allowed")

	}
}
func Uploadtest(w http.ResponseWriter, r *http.Request) {
	// Ensure that only GET and POST requests are allowed
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Check if the request is a POST request
	if r.Method == http.MethodPost {
		// Parse the multipart form containing the file
		err := r.ParseMultipartForm(10 << 20) // Set max memory for file size
		if err != nil {
			http.Error(w, "Failed to parse multipart form", http.StatusInternalServerError)
			return
		}

		// Retrieve the file from the form data
		file, handler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Failed to retrieve file", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// Create a temporary file to save the uploaded file
		tempFile, err := ioutil.TempFile("", "uploaded-*.jpg")
		if err != nil {
			http.Error(w, "Failed to create temporary file", http.StatusInternalServerError)
			return
		}
		defer tempFile.Close()

		// Copy the contents of the uploaded file to the temporary file
		_, err = io.Copy(tempFile, file)
		if err != nil {
			http.Error(w, "Failed to copy file contents", http.StatusInternalServerError)
			return
		}

		// Print the uploaded file name (for testing purposes)
		println("Uploaded file:", handler.Filename)

		// Send a success response to the client with 200 status code
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Uploaded file successfully")
		return
	}

	// Handle GET requests (if needed)
	// Add GET request handling logic here
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "GET request received")
}

/*
func Uploadtest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost, http.MethodGet:
		// Handle POST request
		// You can add your logic for handling POST requests here
		// For now, let's just respond with a 200 OK status
		w.WriteHeader(http.StatusOK)
		// Optionally, you can write a response body here if needed
		// Example: w.Write([]byte("POST request received"))
	default:
		// For all other request methods, respond with a 405 Method Not Allowed status
		w.WriteHeader(http.StatusMethodNotAllowed)
		// Optionally, you can provide a message indicating that only POST requests are allowed
		// Example: w.Write([]byte("Only POST requests are allowed"))
	}
}
*/
func ApiHomepage(w http.ResponseWriter, r *http.Request) {
	info := []byte("Api portion of site homepage")
	w.Write(info)
}
