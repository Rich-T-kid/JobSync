package ChatServer

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// Same endpoint will be used for Uploading proifile pictures as well as file uploads in chat room
// upload images to aws s3 bucket
func UploadImage(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to upload file", http.StatusInternalServerError)
		return
	}
	defer file.Close()
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
	// Save the uploaded file to the server
	// For simplicity, we are just printing the file name
	println("Uploaded file:", handler.Filename)
	fmt.Fprint(w, "uploaded to s3Bucket")
}

// retirve image from aw3
func RetriveUploadedImage(w http.ResponseWriter, r *http.Response) {
	/*
		store link to these image in databse or somthing of that nature
	*/

}
