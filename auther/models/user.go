package models

import (
	"encoding/json"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
)

// There are stuff I want in the JWT, stuff I want in the DB and stuff I want the user to POST to me
type User struct {
	// https://gorm.io/docs/models.html#Fields-Tags
	ID            uint64 `json:"id" gorm:"primaryKey,autoIncrement"`
	Name          string `json:"name" gorm:"not null"`
	Email         string `json:"email" gorm:"uniqueIndex,not null" validate:"required,email"`
	Password      string `json:"password" validate:"required,password"`
	IsActive      bool   `json:"is_active"`
	IsAdmin       bool   `json:"is_admin" gorm:"default:false"`
	hashed        bool
	ConfirmedMail bool
	CreatedAt     time.Time
	EditedAt      time.Time
	//DeletedOn     time.Time
}

type UserRegistration struct {
	Name     string `json:"name" validate:"required,alpha"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
	hashed   bool
}

const PasswordCost = 10

func (u *User) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("password", validateUserPasswordField)
	return validate.Struct(u)
}
func (u *UserRegistration) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("password", validateUserPasswordField)
	return validate.Struct(u)
}

func validateUserPasswordField(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9_+=:;.,<> -]{9,20}$`)
	matches := re.FindAllString(fl.Field().String(), -1)

	return len(matches) == 1
}

// hashUserPassword hashes a user's password and returns it
func (u *UserRegistration) HashUserPassword() (string, error) {
	if !u.hashed {
		bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), PasswordCost)
		u.Password = string(bytes)
		return u.Password, err
	}
	u.hashed = true
	return u.Password, nil
}
func (u *User) HashUserPassword() (string, error) {
	if !u.hashed {
		bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), PasswordCost)
		u.Password = string(bytes)
		return u.Password, err
	}
	u.hashed = true
	return u.Password, nil
}

// ComparePassword checks user password
// What about logging, here?
func ComparePassword(hash, providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(providedPassword))
	return err == nil
}

func (u *User) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(u)
}
func (u *User) ToJSON(r io.Writer) error {
	d := json.NewEncoder(r)
	return d.Encode(u)
}
