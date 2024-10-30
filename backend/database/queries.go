package database

import (
	"database/sql"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	_ "github.com/lib/pq"
	"github.com/palSagnik/daily-expenses-application/models"
)

// user queries
func AddUserToVerify(c *fiber.Ctx, user *models.User) error {
	
	// deleting any previous record of user
	query := `DELETE FROM verifications WHERE email = $1;`
	_, err := DB.Exec(query, user.Email)
	if err != nil {
		log.Warn(err)
		return err
	}
	
	// adding user to verification table
	query = `INSERT INTO verifications (email, name, password, created_at) VALUES ($1, $2, $3, $4);`
	_, err = DB.Exec(query, user.Email, user.Name, user.Password, time.Now())
	if err != nil {
		log.Warn(err)
		return err
	}

	log.Infof("added user with email '%s' for verification", user.Email)
	return nil
}

// deleting user by email -> unique property
func DeleteUser(c *fiber.Ctx) error {
	email := c.Params("email")
	
	query := `DELETE FROM users WHERE email = $1;`
	_, err := DB.Exec(query, email)
	if err != nil {
		log.Warn(err)
		return err
	}

	log.Infof("deleted user with email '%s'", email)
	return nil
}


func AddUser(c *fiber.Ctx, email string) (string, error) {

	// checking if user already exists
	found, err := doesEmailExist(email)
	if found {
		return "user already exists", errors.New("token already verified")
	}
	if err != nil {
		return err.Error(), err
	}

	// email does not exist in the user table hence fetch from verification table
	var verifiedUser models.Verification

	query := `SELECT email, name, password FROM verifications WHERE email=$1`
	row := DB.QueryRow(query, email)
	if err := row.Scan(&verifiedUser.Email, &verifiedUser.Name, &verifiedUser.Password); err != nil {
		log.Warn(err)
		return  err.Error(), err
	}

	// creating user
	query = `INSERT INTO users (email, name, password) VALUES ($1, $2, $3);`
	_, err = DB.Exec(query, verifiedUser.Email, verifiedUser.Name, verifiedUser.Password)
	if err != nil {
		log.Warn(err)
		return err.Error(), err
	}

	// deleting from verification table
	query = `DELETE FROM verifications WHERE email=$1`
	_, err = DB.Exec(query, email)
	if err != nil {
		log.Warn(err)
		return err.Error(), err
	}
	
	log.Infof("added user with email '%s' for users", email)
	return "", nil
}

// gettting user details
func GetUserDetails (c *fiber.Ctx, userid int) (*models.User, error) {
	var user models.User
	
	log.Infof("fetching user details of userid '%d'", userid)
	query := `SELECT email, name FROM users WHERE user_id=$1`
	row := DB.QueryRow(query, userid)
	switch err := row.Scan(&user.Email, &user.Name); err {
	case sql.ErrNoRows:
		return nil, sql.ErrNoRows
	case nil:
		return &user, err
	default:
		return nil, err
	}
}

func GetUsers (c *fiber.Ctx) ([]*models.UserDetails, error) {

	query := `SELECT user_id, name, email FROM users ORDER BY user_id;`
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.UserDetails
	for rows.Next() {
		user := &models.UserDetails{}

		err := rows.Scan(
			&user.UserID,
			&user.Name,
			&user.Email,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// misc queries
// to validate creds -> during login
func ValidateCreds(c *fiber.Ctx, creds *models.Credentials) error {

	query := `SELECT user_id FROM users WHERE email=$1 AND password=$2;`
	_, err := DB.Exec(query, creds.Email, creds.Password)
	if err != nil {
		log.Warn(err)
		return err
	}

	log.Infof("verified credentials for '%s'", creds.Email)
	return nil
}

// to check whether there is a duplicate email at the time of signup
func doesEmailExist(email string) (bool, error) {
	var id int

	query := `SELECT user_id FROM users WHERE email=$1`
	row := DB.QueryRow(query, email)
	switch err := row.Scan(&id); err {
	case sql.ErrNoRows:
		return false, nil
	case nil:
		return true, err
	default:
		return false, err
	}
}