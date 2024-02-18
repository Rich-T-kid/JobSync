package DB

import "time"

 
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
func GenerateHash(password string) string {
	return "not implemented currently"
}
