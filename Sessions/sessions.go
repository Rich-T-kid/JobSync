package Sessions

import (
	"errors"
	"net/http"
	"proj/DB"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
)

var (
	ErrNoCookie = errors.New("http: named cookie not present")
	key         = []byte("super-secret-key")
	store       = sessions.NewCookieStore(key)
)

type CreateCookie interface {
	createcookie() *http.Cookie
}

/*
add in structs that will represesnt the setting cookies so that it can be converted to json
*/
func GatherUserCookies(username string) ([]*http.Cookie, error) {
	var CookieJar []*http.Cookie
	var PrivStruct DB.PrivacySettings
	var AppeaStruct DB.AppearanceSettings
	var PermStruct DB.Permissions
	var NotifStruct DB.NotificationSettings

	PrivacyData, err := StructToJson(PrivStruct, username)
	AppearanceData, err := StructToJson(AppeaStruct, username)
	PermiData, err := StructToJson(PermStruct, username)
	NotificationData, err := StructToJson(NotifStruct, username)
	if err != nil {
		return nil, err
	}
	UserNameCookie := CreateNameCookie(username)
	PrivacyCookie := PrivacySettingCookie(PrivacyData)
	AppearenceCookie := AppearanceSettingscookie(AppearanceData)
	PermCookie := PermissionsCookie(PermiData)
	NotifsCookie := NotificationCookie(NotificationData)
	CookieJar = append(CookieJar, UserNameCookie, PrivacyCookie, AppearenceCookie, PermCookie, NotifsCookie)
	return CookieJar, nil
}

func generateSessionID() string {
	return uuid.NewString()

}

func FormatedTime() string {

	currentTime := time.Now()

	// Define the layout for the desired format
	layout := "January 2 03:04 PM" // Month Day Hour:Minute AM/PM

	// Format the current time using the layout
	formattedTime := currentTime.Format(layout)
	return formattedTime
}

func CreateSessionCookie(username string) *http.Cookie {
	sessionIDString := generateSessionID()
	cookie := http.Cookie{
		Name:     "SessionID",
		Value:    sessionIDString,
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode}
	return &cookie

}

func CreateNameCookie(DBUserName string) *http.Cookie {
	cookie := http.Cookie{
		Name:     "UserNameCookie",
		Value:    DBUserName,
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: false,
		SameSite: http.SameSiteLaxMode,
	}
	return &cookie
}

func ProfieSettingCookie() {

} //all will return a cookie pointer  *http.Cookie{}
func NotificationCookie(String_NotifsData string) *http.Cookie {
	cookie := http.Cookie{
		Name:     "NotificationCookie",
		Value:    EncodeJSONToBase64(String_NotifsData[1 : len(String_NotifsData)-1]),
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	return &cookie
}

func PermissionsCookie(String_PrivData string) *http.Cookie {
	cookie := http.Cookie{
		Name:     "PermissionsCookie",
		Value:    EncodeJSONToBase64(String_PrivData[1 : len(String_PrivData)-1]),
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	return &cookie
}

// DB.StoreCookie()
func AppearanceSettingscookie(String_AppData string) *http.Cookie {
	cookie := http.Cookie{
		Name:     "AppearanceSetting",
		Value:    EncodeJSONToBase64(String_AppData[1 : len(String_AppData)-1]),
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	return &cookie
}

func PrivacySettingCookie(String_PrivData string) *http.Cookie {
	cookie := http.Cookie{
		Name:     "PrivacySettings",
		Value:    EncodeJSONToBase64(String_PrivData[1 : len(String_PrivData)-1]),
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	return &cookie
}
