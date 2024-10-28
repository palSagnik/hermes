package models

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

// users table
// id, email, number, name, password, expense
type User struct {
	UserID      int          `json:"userid"`
	Email       string       `json:"email"           form:"email"`
	Name        string       `json:"name"            form:"name"`
	Number      string       `json:"number"          form:"number"`
	Expenses  	[]Expense    `json:"expenses"        form:"expenses"`
	Password    string       `json:"password"        form:"password"`
	ConfirmPass string       `json:"confirm"         form:"confirm"`
}

// toverify --> vid, username, email, password, timestamp
// helpful for sending verification emails and updating users database
type Verification struct {
	VerificationID int       `json:"vid"`
	Email          string    `json:"email"    form:"email"`
	Name           string    `json:"name"     form:"name"`
	Number         string    `json:"number"   form:"number"`
	Password       string    `json:"password" form:"password"`
	CreatedAt      time.Time
}

type UserDetails struct {
	UserID      int          `json:"userid"`
	Email       string       `json:"email"           form:"email"`
	Name        string       `json:"name"            form:"name"`
	Number      string       `json:"number"          form:"number"`
	Expenses  	[]Expense    `json:"expenses"        form:"expenses"`
}

// struct for credentials
// will be used during login
type Credentials struct {
	Email    string `json:"email"    form:"email"`
	Password string `json:"password" form:"password"`
}

// midlleware token verification struct
type VerifyClaims struct {
	jwt.RegisteredClaims
	Email string  `json:"email"`
}
