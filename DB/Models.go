package DB

import ("time")


var UserSlice []User
type User struct{
	Name string
	Password interface{}
	/*Email string
	PhoneNumber string
	*/
}

type Session struct{
	SessionID string 
	UserID int 
	CreatedAt time.Time
	ExpiresAt time.Time
}
