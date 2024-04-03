package DB

import (
	"crypto/sha256"
	"fmt"
	"time"
)

var (
	hashObject = sha256.New()
)

func CurrentTime() time.Time {
	return time.Now()
}

func IsDBCOnnectionstillAlive() bool {
	err := db.Ping()
	if err != nil {
		return false
	}
	return true
}

// create a user in database for the first time
func InputUser(Username, password, Email string, Phone_number interface{}) error {
	db, err := DBConnection() // handle error later
	if err != nil {
		return err
	}
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE Username = ?", Username).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("User already exists")
	}
	// Validate.go Has a working has function include that here and in the valid login function
	password = hashPassword(password)
	formatedQuery := fmt.Sprintf("insert into users(Username,password,Email,Phone_number) values (\"%s\",\"%s\",\"%s\",\"%s\")", Username, password, Email, Phone_number)
	_, er := db.Exec(formatedQuery)
	return er
}
func GenerateHashfake(password string) string {
	return password
}

func RealLogin(username, password string) (string, error) {
	db, err := DBConnection()
	if err != nil {
		return "", fmt.Errorf("Database Connection down")
	}
	password = hashPassword(password)
	query := "SELECT username FROM users WHERE Username = ? AND password = ?"
	row, err := db.Query(query, username, password)
	if err != nil {
		return "", err
	}
	rowcount := 0
	var DbEntry UserDB
	for row.Next() {
		err := row.Scan(&DbEntry.Username)
		if err != nil {
			return "", err
		}
		rowcount++
	}
	if rowcount == 0 {
		return "", fmt.Errorf("User Doesn't Exist")
	}
	return DbEntry.Username, nil
}
func ValidLogin(username string, password string) bool {
	return true

}
