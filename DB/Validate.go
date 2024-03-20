package DB


import (
    "crypto/sha256"
    "encoding/hex"
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

func userNameAlreadyExist() bool {return false}//use this as validation for the front end javascript code
