package DB

import ("time")


var UserSlice []User
type User struct{
	Name string
	Password interface{}
	SessID  string
	/*Email string
	PhoneNumber string
	*/
}

type UserDB struct{
UserID int
Username string
password string
email string
phone string

}

type Session struct{
	SessionID string 
	UserID int 
	CreatedAt time.Time
	ExpiresAt time.Time
}

func RemoveUserSessionSlice(Usid string) {
    for idx, user := range UserSlice {
        if user.SessID == Usid {
            UserSlice = append(UserSlice[:idx], UserSlice[idx+1:]...)
}}}
