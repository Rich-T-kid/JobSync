package DB

import (
	"fmt"
	"crypto/sha256"
    	//"encoding/hex"
//	"strings"

	"time"
//	"strconv"
)
var (

	hashObject = sha256.New()
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
	password = GenerateHash(password)
	formatedQuery := fmt.Sprintf("insert into Users(Username,PasswordHash,Email,Phone_number) values (\"%s\",\"%s\",\"%s\",\"%s\")", Username, password, Email, Phone_number)
	_, er := db.Exec(formatedQuery)
	return  er
}
/*
func GenerateHash(password string) string {
	password = strings.TrimSpace(password)
	hashObject.Write([]byte(password))
	hashedBytes := hashObject.Sum(nil)
	//hashedString  := fmt.Sprintf("%x" , hashedBytes)
	//fmt.Println("produced string" , hashedString)
	//return hashedString	
	fmt.Println("password: ", password , "Generated hash: " , hex.EncodeToString(hashedBytes))
	return hex.EncodeToString(hashedBytes)
}
*/
func GenerateHash(password string) string{
	return password}

func RealLogin(username , password string)  (string , error)  {
	db , err := DBConnection()
	if err != nil{
		return "", fmt.Errorf("Database Connection down")
	}
	fmt.Println("pre password db will check" , password)
	password = GenerateHash(password) 
	query := "SELECT username FROM Users WHERE Username = ? AND PasswordHash = ?"
	row , err := db.Query(query , username , password)
	if err != nil{
		return "", err
	}
	rowcount := 0
	var DbEntry UserDB
	for row.Next() {
		err := row.Scan(&DbEntry.Username)
		if err != nil{
		return "", err}
		rowcount++ }
	if rowcount == 0{
	return "", fmt.Errorf("User Doesn't Exist")
	}
	return  DbEntry.Username, nil
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
