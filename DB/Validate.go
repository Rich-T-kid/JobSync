package DB

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
<<<<<<< HEAD
=======
	"encoding/json"
    "crypto/sha256"
    "encoding/hex"
>>>>>>> Cookies
)

var seed = []byte("db746edd-40a4-4015-a1e3-1572c64ba865") // Define a fixed seed

func hashPassword(password string) string {
	hash := sha256.New()
	hash.Write(seed) // Write the seed to the hash function
	hash.Write([]byte(password))
	hashedPassword := hex.EncodeToString(hash.Sum(nil))
	return hashedPassword
}

//func CheckPasswordHash(password, hash string) bool { return HashPassword(password) == hash}

func userNameAlreadyExist() bool { return false } //use this as validation for the front end javascript code

// this goroutine runs for the whole lifespan of the program cleaning up user sessions. could move to its own api
func CleanDB() error {
	for {
		prestatment := "DELETE FROM user_cookie_sessions WHERE expiration_timestamp <= NOW();"
		delStatment, err := db.Prepare(prestatment)
		if err != nil {
			fmt.Println("error cleaning database cookies:", err)
			return err
		}
		defer delStatment.Close()

		_, err = delStatment.Exec()
		if err != nil {
			return err
		}

		// Sleep for a while before running the cleanup again
		time.Sleep(1 * time.Minute)
	}
}
<<<<<<< HEAD
=======
func stringsToJSON(args ...string) (string, error) {
    // Convert variadic input strings to a slice of strings
    stringsSlice := []string(args)

    // Marshal the slice into JSON format
    jsonString, err := json.Marshal(stringsSlice)
    if err != nil {
        return "", err
    }

    return  string(jsonString) , err 
}
>>>>>>> Cookies
