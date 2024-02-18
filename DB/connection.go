package DB
import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "sync"
    "fmt"
)



var (Once sync.Once
	db *sql.DB
)


func DBConnection() (*sql.DB,error)  {
	if db == nil{
	connection , err := createDBconnection()
	if err != nil{
	fmt.Println("error in creating DB instance", err)
	return nil, err	}
	db = connection 	}
	return db , nil
	
}
func createDBconnection() (*sql.DB, error) {
    username := "root"
    password := "1019026"
    databaseName := "table"
    // Create the data source name (DSN)
    dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", username, password, databaseName)
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }
    return db, nil
}
type human struct{
	id int
	job string
}
func GrabData(db *sql.DB) {
	rows , err := db.Query("select id , job from users")
	if err != nil{
		fmt.Println(err)
		return}
	defer rows.Close()	
	for rows.Next(){
	var person human
	er  := rows.Scan(&person.id,&person.job)
	if er != nil{
		fmt.Println(er)
	}else{
		fmt.Println(person)
	}
}}
