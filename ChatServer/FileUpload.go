package ChatServer
import(
	"net/http"
	
	)
// Same endpoint will be used for Uploading proifile pictures as well as file uploads in chat room

func UploadImage(w http.ResponseWriter , r *http.Response){
	/*
	store in S3 bucket for now

	*/

}


func RetriveUploadedImage(w http.ResponseWriter , r *http.Response){
 	/*
	store link to these image in databse or somthing of that nature
	*/

}
