package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

const saltRounds = 10

func SetPassword(password string) string {
    // Generate hashed password with salt
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), saltRounds)
    if err != nil {
        log.Println("Error generating hashed password:", err)
        return ""
    }
    return string(hashedPassword)
}