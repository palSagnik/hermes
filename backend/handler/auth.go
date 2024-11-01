package handler

import (
	"errors"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/palSagnik/hermes/config"
	"github.com/palSagnik/hermes/database"
	"github.com/palSagnik/hermes/middleware"
	"github.com/palSagnik/hermes/models"
	"github.com/palSagnik/hermes/utils"
)

func Signup(c *fiber.Ctx) error {
	signup := new(models.User)

	limiter := middleware.GetVisitor(c.IP())
	if !limiter.Allow() {
		return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{"status": "failed", "message": "too many requests"})
	}

	// assigning form values
	signup.Email = c.FormValue("email")
	signup.Name = c.FormValue("name")
	signup.Password = c.FormValue("password")
	signup.ConfirmPass = c.FormValue("confirm")

	// handling error if any of the fields are empty
	if signup.Email == "" || signup.Name == "" || signup.Password == "" || signup.ConfirmPass == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed", "message": "all fields must be filled"})
	}

	// removing extra space in form fields
	signup.Email = strings.TrimSpace(signup.Email)
	signup.Password = strings.TrimSpace(signup.Password)
	signup.ConfirmPass = strings.TrimSpace(signup.ConfirmPass)
	signup.Name = strings.TrimSpace(signup.Name)

	// converting email to lowercase
	signup.Email = strings.ToLower(signup.Email)

	// checking whether signup information is valid or not
	isOk, status := utils.VerifySignupInput(signup)
	if !isOk {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed", "message": status})
	}

	// Store the hash of the password in the database
	signup.Password = utils.GenerateHash(signup.Password)

	// Add to verification table to send mail
	err := database.AddUserToVerify(c, signup)
	if err != nil {
		log.Warn(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed", "message": "please contact admin"})
	}

	// send verification mail
	if err := utils.SendVerificationMail(signup); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed", "message": "error in sending verification email"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "check your email for verification"})
}

func Login(c *fiber.Ctx) error {
	creds := new(models.Credentials)
	user := new(models.User)

	limiter := middleware.GetVisitor(c.IP())
	if !limiter.Allow() {
		return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{"status": "failed", "message": "too many requests"})
	}

	// form values
	creds.Email = c.FormValue("email")
	creds.Password = c.FormValue("password")

	// empty field checks
	if creds.Email == "" || creds.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed", "message": "all fields must be filled"})
	}

	// formatting form fields
	creds.Email = strings.TrimSpace(creds.Email)
	creds.Password = strings.TrimSpace(creds.Password)
	creds.Email = strings.ToLower(creds.Email)

	// verifying login input
	if isOk, statusMsg := utils.VerifyLoginInput(creds); !isOk {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed", "message": statusMsg})
	}

	// hashing the password
	creds.Password = utils.GenerateHash(creds.Password)

	// checking email and password
	if err := database.ValidateCreds(c, creds); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "failure", "message": "invalid credentials"})
	}

	// token generation and assigning
	token, err := middleware.GenerateToken(user)
	if err != nil {
		log.Warn(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failure", "message": "error in token generation. contact admin"})
	}

	// assigning cookie values
	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.SameSite = fiber.CookieSameSiteStrictMode
	cookie.Expires = time.Now().Add(72 * time.Hour)

	c.Cookie(cookie)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "login successful"})
}

// function to verify token
func Verify(c *fiber.Ctx) error {
	
	limiter := middleware.GetVisitor(c.IP())
	if !limiter.Allow() {
		return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{"status": "failed", "message": "too many requests"})
	}
	
	token := c.Query("token")
	if token == "" {
		return c.Status(fiber.StatusBadRequest).SendString("missing token! Register again")
	}

	claims := new(models.VerifyClaims)
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("different signing method for token")
		}

		return []byte(config.TOKEN_SECRET), nil
	})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if claims.Email == "" {
		return c.Status(fiber.StatusBadRequest).SendString("error in token, register again!")
	}

	if msg, err := database.AddUser(c, claims.Email); err != nil {
		log.Warn(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failure", "message": msg})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "user added! proceed to login"})
}
