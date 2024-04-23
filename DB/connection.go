package DB

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"path/filepath"
	"mime"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	Once sync.Once
	db   *sql.DB
	s3Bucket *s3.S3
    AccessKeyId     string
    SecretAccessKey string
    AwsRegion       string
    BucketName      string
)



func createDBConnection() (*sql.DB, error) {
	// Open a connection to the MySQL server
	// MySQL connection details
	host := os.Getenv("host")
	port := os.Getenv("port")
	username := os.Getenv("username")
	password := os.Getenv("password")
	databaseName := os.Getenv("databaseName")
	// Create the data source name (DSN)
	port1, _ := stringToNumber(port)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port1, databaseName)
	// Open a connection to the MySQL server
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
func DBConnection() (*sql.DB, error) {
	if db == nil {
		connection, err := createDBConnection()
		if err != nil {
			fmt.Println("Richard, double check that the Docker container running the DB instance is running")
			fmt.Println("error in creating DB instance:", err)
			return nil, err
		}
		db = connection // Assigning the connection to the global variable db
	}
	return db, nil
}

func loadEnviromentVars() {
    log.Println("loading enviroment varibaleS")
    godotenv.Load()
    AccessKeyId = os.Getenv("AccessKeyId")
    SecretAccessKey = os.Getenv("SecretAccessKey")
    AwsRegion = os.Getenv("awsRegion")
    BucketName = os.Getenv("bucketName")
    log.Println(AccessKeyId,SecretAccessKey,AwsRegion,BucketName)
}
func stringToNumber(input string) (int, error) {
	// Convert string to integer
	number, err := strconv.Atoi(input)
	if err != nil {
		// Return 0 and the error if conversion fails
		return 0, err
	}
	// Return the converted number and no error
	return number, nil
}

func StartConnection() {
	loadEnviromentVars()
	DBConnection()
	go CleanDB()
	initglobalS3bucket()

}
func  newS3Client(creds *credentials.Credentials) (*s3.S3,error) {
	sess, err := session.NewSession(&aws.Config{
	 Credentials: creds,
	    Region: aws.String("us-east-2")},
	)
	if err != nil{return nil, err}
	s3LiveSession := s3.New(sess)
	return s3LiveSession , nil 

}


func (m *msgUploader) Upload(file *os.File) error {
    svc := m.s3Bucket
    // Get file info
    fileInfo, err := file.Stat()
    if err != nil {
        return err
    }
    // Set content type based on file extension
    contentType := mime.TypeByExtension(filepath.Ext(fileInfo.Name()))

    // Upload file to S3 bucket in specific folder
    _, err = svc.PutObject(&s3.PutObjectInput{
        Bucket:      aws.String(m.parentBucketName), //jobsynce gernal bucker
        Key:         aws.String(m.bucketName + fileInfo.Name()), // specfic peer to peer bucket
        ContentType: aws.String(contentType),
        Body:        file,
    })
    if err != nil {
        return err
    }

    return nil
}

func  NewmsgUploader(name string,s3Connection *s3.S3) *msgUploader{
	return &msgUploader{parentBucketName:"job-sync-bucket",bucketName:name,s3Bucket:s3Connection}
}


// implement aws connection as well as s3 bucket connection and set up a clean and eeasy to use interface from here. Encapsulate all the s3 bucket storage logic into a wrapper struct
func initglobalS3bucket() (*s3.S3,error) {
	if s3Bucket == nil{
		creds := credentials.NewStaticCredentials(AccessKeyId, SecretAccessKey, "")
		S3client , err := newS3Client(creds)	
		if err != nil {return nil , fmt.Errorf("Error setting up aws bucket connection: %v", err)}
	s3Bucket = S3client }
	return s3Bucket,nil

}
func BucketInstance(name string)*msgUploader{
	s3Bucket, err := initglobalS3bucket()
	if err != nil{panic(err)}
	return NewmsgUploader(name,s3Bucket)
}


