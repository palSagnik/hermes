package utils

import (
	"crypto/sha256"
	"fmt"
	"net/mail"

	"github.com/palSagnik/daily-expenses-application/config"
	"github.com/palSagnik/daily-expenses-application/models"
)

func GenerateHash(secret string) string {
	hash := sha256.New()
	hash.Write([]byte(secret))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func VerifySignupInput(signup *models.User) (bool, string) {
	
	// password verification
	password := signup.Password
	confirmPassword := signup.ConfirmPass

	passLength := len(password)
	if passLength > config.PASS_LEN || passLength < 8 {
		return false, fmt.Sprintf("Password should be of length 8-%d characters", config.PASS_LEN)
	}

	if password != confirmPassword {
		return false, "your passwords do not match, please try again"
	}

	// number verification
	number := signup.Number
	if len(number) != 10 {
		return false, "number must be exactly 10 digits long"
	}

	// name verification
	name := signup.Name
	if len(name) > config.NAME_LEN {
		return false, fmt.Sprintf("name must be less than %d characters", config.NAME_LEN)
	}

	// email verification
	email := signup.Email
	if len(email) > config.MAIL_LEN {
		return false, fmt.Sprintf("email should not exceed %d characters", config.MAIL_LEN)
	}

	// checking if the email address is valid
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false, "not a valid email address"
	}

	return true ,""
}


func VerifyLoginInput(creds *models.Credentials) (bool, string) {

	// email length verification
	emailLength := len(creds.Email)
	if emailLength > config.MAIL_LEN {
		return false, fmt.Sprintf("Email should not exceed %d characters", config.MAIL_LEN)
	}

	// password length verification
	passLength := len(creds.Password)
	if passLength > config.PASS_LEN || passLength < 8 {
		return false, fmt.Sprintf("Password should be of length 8-%d characters", config.PASS_LEN)
	}

	if _, err := mail.ParseAddress(creds.Email); err != nil {
		return false, "Not a valid email address"
	}
	return true, ""
}