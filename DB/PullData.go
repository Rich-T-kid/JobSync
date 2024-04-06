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
<<<<<<< HEAD
	return er
}
func GenerateHashfake(password string) string {
	return password
=======
	err = inputUserInDB(Username)
	if err != nil{
		fmt.Println(err)
		return err}
	return  er
>>>>>>> Cookies
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

func inputUserInDB(Username string)error {
	  stmt0, err := db.Prepare("INSERT INTO privacy_settings (user_id) VALUES (?)")
    if err != nil {
        fmt.Println("Error preparing SQL statement:InputUserINDB", err)
        return err
    }
    defer stmt0.Close()

	stmt1, err := db.Prepare("INSERT INTO permission_table (user_id, permission) VALUES (? , ?) ")
    if err != nil {
        fmt.Println("Error preparing SQL statement:", err)
        return err
    }
    defer stmt1.Close()
    
stmt2, err := db.Prepare("INSERT INTO notification_settings (user_id) VALUES (?)")
    if err != nil {
        fmt.Println("Error preparing SQL statement:", err)
        return err
    }
    defer stmt2.Close()

stmt3, err := db.Prepare("INSERT INTO appearance_settings (user_id, content_filters) VALUES (?, ?) ")
    if err != nil {
        fmt.Println("Error preparing SQL statement:", err)
        return err
    }
    defer stmt3.Close()

    // Define the user ID value
    yourUserID := idFromUserName(Username) // Replace with the actual user ID

    // Execute the SQL statement with the user ID
    _, err = stmt0.Exec(yourUserID)

   if err != nil{fmt.Println(err)} 
   defaultPermissionsStr , err := stringsToJSON("Read","Write")
   if err != nil{fmt.Println(err)}
   fmt.Println("jsonnnn string" , defaultPermissionsStr)
    _, err = stmt1.Exec(yourUserID, defaultPermissionsStr)

    fmt.Println("error from perimson insertion: , ", err)
   if err != nil{fmt.Println(err)} 
    _, err = stmt2.Exec(yourUserID)

   if err != nil{fmt.Println(err)} 
   defualtFilters := `["Explicit"]`
    _, err = stmt3.Exec(yourUserID,defualtFilters)
    
   if err != nil{fmt.Println(err)} 
   return nil
}


/*
just for security reasons, overight all of the UserID return values from
Scanning. Set it all to -1 as this cannot exist in the databse 

*/

func (p PrivacySettings) DbtoStruct(username string ) (interface{} ,error)  {
	var pSetting PrivacySettings

	fmt.Println("privacy settings  before ; ", pSetting)
	userid := idFromUserName(username)
	query := "select * from privacy_settings where user_id = ?"
	row , err := db.Query(query,userid)
	if err != nil{
	return nil , err
	}

	var date string
	for row.Next(){
	err := row.Scan(&pSetting.userID , &pSetting.UsernameVisibility , &pSetting.FriendRequestsVisibility , &pSetting.ContentVisibility,&date)
	if err != nil{
		return nil , err}
	}
	validTime , err := parseTimeString(date)
	PP := &pSetting
	PP.userID = -1
	PP.LastUpdated = validTime
	return &pSetting , nil
}


func (a AppearanceSettings) DbtoStruct(username string) (interface{} , error) {
	var pAppearanceSettings  AppearanceSettings
	userid := idFromUserName(username)
	query := "select * from appearance_settings where user_id = ?"
	row , err := db.Query(query , userid)
	if err != nil{
		return nil , err
	}
	for row.Next(){
		// Todo fill in the fields
		err := row.Scan(
    &pAppearanceSettings.userID,
    &pAppearanceSettings.Theme,
    &pAppearanceSettings.FontSize,
    &pAppearanceSettings.ColorScheme,
    &pAppearanceSettings.BackgroundImage,
    &pAppearanceSettings.Language,
    &pAppearanceSettings.ContentFilters,
)
	if err != nil{
		return nil , err }

	}

	PP := &pAppearanceSettings
	PP.userID = -1
	return &pAppearanceSettings , nil
}

func (n NotificationSettings) DbtoStruct(username string)(interface{} , error)  {
	var pNotificationSettings NotificationSettings

	userid := idFromUserName(username)
	query := "select * from notification_settings where user_id = ?"
	row , err := db.Query(query,userid)
	if err != nil{
		return nil , err
	}
	for row.Next(){

	err := row.Scan(&pNotificationSettings.userID, &pNotificationSettings.EmailNotifications ,&pNotificationSettings.PushNotifications ,&pNotificationSettings.NotificationFrequency )
	if err != nil{
		return nil , err}
	}
	PP := &pNotificationSettings
	PP.userID = -1
	return &pNotificationSettings , nil
}


func (p Permissions) DbtoStruct(username string) (interface{} , error) {
	var pPermissions Permissions
	userid := idFromUserName(username)
	query := "select * from permission_table where user_id = ?"
	row , err := db.Query(query , userid)
	if err != nil{
		return nil ,err}
	var date string
	for row.Next(){
		err := row.Scan(&pPermissions.userID,&pPermissions.ID,&pPermissions.Permissions,&date)
		if err != nil{
			return nil , err}
	}
	ValidDate , err := parseTimeString(date) 
	PP := &pPermissions
	PP.userID = -1
	PP.LastUpdated = ValidDate
	return &pPermissions , nil
} 



