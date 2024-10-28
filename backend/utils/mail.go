package utils

import (
	"bytes"
	"fmt"
	"net/smtp"
	"text/template"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/palSagnik/daily-expenses-application/config"
	"github.com/palSagnik/daily-expenses-application/models"
)


func getToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"email": user.Email,
		"expiry": time.Now().Add(time.Minute * time.Duration(config.TOKEN_EXPIRY)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.TOKEN_SECRET))
	if err != nil {
		return "", err
	}
	
	return tokenString, nil
}

func SendVerificationMail(user *models.User) error {

	token, err := getToken(user)
	if err != nil {
		return err
	}

	secret := config.EMAIL_AUTH
	fromEmail := config.EMAIL_ID
	toEmail := []string{
		user.Email,
	}
	auth := smtp.PlainAuth("", fromEmail, secret, config.SMTP_HOST)
	templ, _ := template.ParseFiles("template/mail.html")

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Hermes Backend Auth\n%s\n\n", mimeHeaders)))
	
	templ.Execute(&body, struct {
		Username string
		Link     string
	}{
		Username: user.Name,
		Link: config.AUTH_URL + token,
	})

	err = smtp.SendMail(config.SMTP_HOST + ":" + config.SMTP_PORT, auth, fromEmail, toEmail, body.Bytes())
	if err != nil {
		return err
	}

	return nil
}