package ChatServer
import(
	"net/http"
	
	)
// Same endpoint will be used for Uploading proifile pictures as well as file uploads in chat room
// upload images to aws s3 bucket
func UploadImage(){
	/*
	store in S3 bucket for now

	*/

}

// retirve image from aw3
func RetriveUploadedImage(w http.ResponseWriter , r *http.Response){
 	/*
	store link to these image in databse or somthing of that nature
	*/

}
