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

type  DBMarshall interface{
	DbtoStruct(string) (interface{}, error ) 
}
type UserDB struct{
UserID int
Username string
password string
email string
phone string

}


// UserCookieSession represents the user_cookie_sessions table.
type UserCookieSession struct {
    SessionID           string
    UserID              int
    CreatedAt           string
    ExpirationTimestamp string
}

// PrivacySettings represents the privacy_settings table.
type PrivacySettings struct {
    userID                int
    UsernameVisibility    string
    FriendRequestsVisibility string
    ContentVisibility     string
    LastUpdated time.Time
}

// AppearanceSettings represents the appearance_settings table.
type AppearanceSettings struct {
    userID        int
    Theme         string
    FontSize      string
    ColorScheme   string
    BackgroundImage string
    Language      string
    ContentFilters string
}

type Permissions struct {
	userID int
	ID int
	Permissions string
	LastUpdated time.Time
}
// NotificationSettings represents the notification_settings table.
type NotificationSettings struct {
    userID               int
    EmailNotifications  string
    PushNotifications    string
    NotificationFrequency string
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

