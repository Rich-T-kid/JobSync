package DB

import (
	"database/sql"
<<<<<<< HEAD
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	env "proj/env"
	"sync"
=======
	"strconv"
	"fmt"
	"log"
	"sync"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
>>>>>>> ChatApp
)

var (
	Once sync.Once
	db   *sql.DB
)

func createDBConnection() (*sql.DB, error) {
<<<<<<< HEAD
	// MySQL connection details
	host := "127.0.0.1" // Docker container host
	port := 3305        // Docker container port
	username := "root"
	password := "richard"
	databaseName := "jobsync"

	// Create the data source name (DSN)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, databaseName)

	// Open a connection to the MySQL server
	db, err := sql.Open("mysql", dsn)
	if err != nil {
=======
    // MySQL connection details
    host := os.Getenv("host")
    port := os.Getenv("port")
    username := os.Getenv("username")
    password := os.Getenv("password")
    databaseName := os.Getenv("databaseName")
    // Create the data source name (DSN)
    port1 , _ := stringToNumber(port) 
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port1, databaseName)




    fmt.Println("current sql connection string", dsn)
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
func DBConnection() (*sql.DB,error)  {
	if db == nil{
	connection , err := createDBConnection()
	if err != nil{
		fmt.Println("Richard Doubel check that the Dokcer container running the DB instance is running")
		fmt.Println("error in creating DB instance", err)
>>>>>>> ChatApp
		return nil, err
	}

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
<<<<<<< HEAD
func DBConnection() (*sql.DB, error) {
	if db == nil {
		connection, err := createDBConnection()
		if err != nil {
			fmt.Println("Richard Doubel check that the Dokcer container running the DB instance is running")
			fmt.Println("error in creating DB instance", err)
			return nil, err
		}
		db = connection
	}
	return db, nil

}
func StartConnection() {
=======
func loadEnviromentVars(){
	log.Println("loading enviroment varibaleS")
	godotenv.Load()
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


func StartConnection(){
	loadEnviromentVars()
>>>>>>> ChatApp
	DBConnection()
	go CleanDB()
	awsConnection()

}
func awsConnection(){
	fmt.Println("create aws connection ricahrd")
 //ToDo: create and establish cloud aws connection for image uploads to s3 buckets
}


