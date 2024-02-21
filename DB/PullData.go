package DB

import (
	"fmt"
	"time"
//	"strconv"
)

//var users []User
type User struct{
	Name string
	Password interface{}
	Email string
	PhoneNumber string
}

type Session struct{
	SessionID string 
	UserID int 
	CreatedAt time.Time
	ExpiresAt time.Time
}

func CurrentTime() time.Time{
	return time.Now()
}


func IsDBCOnnectionstillAlive() bool{
	err := db.Ping()
	if err != nil{
	return false
	}
	return true
}

// create a user in database for the first time
func InputUser(Username , password , Email string,Phone_number interface{}) error{
	fmt.Println("called")
	db , _ := DBConnection() // handle error later
	//PhoneNumberStr := strconv.Itoa(Phone_number)
	formatedQuery := fmt.Sprintf("insert into Users(Username,PasswordHash,Email,Phone_number) values (\"%s\",\"%s\",\"%s\",\"%s\")", Username, password, Email, Phone_number)
	_, err := db.Exec(formatedQuery)
	return  err
}
/*
func GenerateHash(password string) string {
	return "not implemented currently"
}
func InputSession() error {}

func ValidPassword() bool {}


// always check if a usernmae is valid first and then check if the hashed password exist in database
func UsernameExist() bool{}



func ValidLogin(username string , password String) bool {}
*/

/*


not finished stil have to validate  the input beforeinputing into database 
also go over the the types of querying for sql and go lbray they are different.
this should return an err.
}
*/
