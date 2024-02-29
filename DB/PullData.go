package DB

import (
	"fmt"
	"hash/fnv"
	"time"
//	"strconv"
)
var (

	hashObject  = fnv.New128a() 
)


func CurrentTime() time.Time{
	return time.Now()
}


func IsDBCOnnectionstillAlive() bool{
	err := db.Ping()
	if err != nil{
		return false }
	return true
}

// create a user in database for the first time
func InputUser(Username , password , Email string,Phone_number interface{}) error{
	fmt.Println("called")
	db , err := DBConnection() // handle error later
	if err != nil{
		return err
	}
	var count int
    	err = db.QueryRow("SELECT COUNT(*) FROM Users WHERE Username = ?", Username).Scan(&count)
    	if err != nil {
        	return err
    	}
    	if count > 0 {
        	return fmt.Errorf("User already exists")
    	}
	//all the hashing will be done internally no need to call it from handlers
	password = GenerateHash(password)
	formatedQuery := fmt.Sprintf("insert into Users(Username,PasswordHash,Email,Phone_number) values (\"%s\",\"%s\",\"%s\",\"%s\")", Username, password, Email, Phone_number)
	_, er := db.Exec(formatedQuery)
	return  er
}
func GenerateHash(password string) string {
	hashObject.Write([]byte(password))
	hashedBytes := hashObject.Sum(nil)
	hashedString  := fmt.Sprintf("%x" , hashedBytes)
	return hashedString	
}
func RealLogin(username , password string) (bool ,error)  {
	db , err := DBConnection()
	if err != nil{
		return false ,fmt.Errorf("Database Connection down")
	}
	password = GenerateHash(password) // convert to hash to check in DB
	query := "SELECT * FROM Users WHERE Username = ? AND PasswordHash = ?"
	row , err := db.Query(query , username , password)
	if err != nil{
		return false , err
	}
	rowcount := 0
	for row.Next(){
	rowcount++ 
	}
	if rowcount == 0{
	return false, fmt.Errorf("User Doesn't Exist")
	}
	return true , nil
}
func ValidLogin(username string , password string) bool {
	return true

}
/*
func InputSession() error {}
 Might just handle this with some vanila java script and html validaton. Waste of server computing to implement this ehre 
func ValidPassword() bool {}




*/

/*


not finished stil have to validate  the input beforeinputing into database 
also go over the the types of querying for sql and go lbray they are different.
this should return an err.
}
*/
