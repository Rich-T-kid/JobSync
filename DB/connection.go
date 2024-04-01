package DB
import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "sync"
    "fmt"
	env"proj/env"
)



var (Once sync.Once
	db *sql.DB
)



func createDBConnection() (*sql.DB, error) {
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
		return nil, err
	}
	db = connection 	}
	return db , nil
	
}
func StartConnection(){
	DBConnection()
	go CleanDB()
	env.InitENV()
	awsConnection()

}
func awsConnection(){
	fmt.Println("create aws connection ricahrd")
 //ToDo: create and establish cloud aws connection for image uploads to s3 buckets
}


