package DB

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	logFilePath = "JobSyncLogs.txt"
)

func TopLogs() (string, error) {
	file, err := ioutil.ReadFile(logFilePath)
	if err != nil {
		return "", fmt.Errorf("unable to read file: %v", err)
	}

	// Split the file content into lines
	lines := strings.Split(string(file), "\n")

	// Get the first 50 lines
	var top50 []string
	for i := 0; i < 50 && i < len(lines); i++ {
		top50 = append(top50, lines[i])
	}

	// Combine the lines back into a single string
	response := strings.Join(top50, "\n")
	return response, nil
}
func idFromUserName(inputusername string) int {
	stmt, err := db.Prepare("SELECT userid FROM users WHERE username = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Username to search
	username := inputusername

	// Execute the SQL query
	var userID int
	err = stmt.QueryRow(username).Scan(&userID)
	if err != nil {
		log.Fatal(err)

	}
	return userID
}

func generateDBCookie(c *http.Cookie, username string) *UserCookieSession {
	var dbCookie UserCookieSession
	userid := idFromUserName(username)
	currentTime := time.Now()
	mysqlTimestamp := currentTime.Format("2006-01-02 15:04:05")
	dbCookie.SessionID = c.Value
	dbCookie.UserID = userid
	dbCookie.CreatedAt = mysqlTimestamp
	dbCookie.ExpirationTimestamp = addOneHour(dbCookie.CreatedAt)
	return &dbCookie
}
func StoreCookie(c *http.Cookie, username string) error {
	cookiestruct := generateDBCookie(c, username)
	stmt, err := db.Prepare("INSERT INTO user_cookie_sessions (session_id, user_id, created_at, expiration_timestamp) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(cookiestruct.SessionID, cookiestruct.UserID, cookiestruct.CreatedAt, cookiestruct.ExpirationTimestamp)
	if err != nil {
		return err
	}
	return nil
}

func addOneHour(mysqlTimestamp string) string {
	// Define the layout for parsing the MySQL timestamp format
	layout := "2006-01-02 15:04:05"
	// Parse the MySQL timestamp string into a time.Time value
	parsedTime, _ := time.Parse(layout, mysqlTimestamp)
	// Add 1 hour to the parsed time
	oneHourLater := parsedTime.Add(time.Hour)
	// Format the new time value back into a string with the same layout
	newTimestamp := oneHourLater.Format(layout)
	return newTimestamp
}
func DeleteCookieSession(c *http.Cookie) error {
	prestatment := " DELETE FROM  user_cookie_sessions where session_id = ?"
	delStatment, err := db.Prepare(prestatment)
	if err != nil {
		return err
	}
	defer delStatment.Close()
	sessionID := c.Value
	_, err = delStatment.Exec(sessionID)
	if err != nil {
		return err
	}
	return nil
}

func AllActiveUsers() ([]string, error) {
	db, err := DBConnection()
	if err != nil {
		return nil, fmt.Errorf("failed to establish database connection: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT username FROM users INNER JOIN user_cookie_sessions ON users.userid = user_cookie_sessions.user_id")
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var Users []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, fmt.Errorf("error occurred scanning names: %v", err)
		}
		Users = append(Users, name)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred while iterating over rows: %v", err)
	}

	return Users, nil
}

// this says store but it more of an update function.since in db certain values are set by defualt. this willa ct as an update. Pass in cookie
// and the database will be updated to match the cookies values and return an error.
func StoreProfileSettings() {}

func StoreNotificationSettings() {}

func StoreContentPrefences() {}

func StoreappeanceSettings() {}

func StoreprivacySettings() {}
