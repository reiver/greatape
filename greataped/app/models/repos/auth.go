package repos

import (
	"contracts"
	"db"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Access int64

const (
	ACCESS_PUBLIC Access = iota
	ACCESS_PRIVATE
)

// User struct defines the user
type User struct {
	gorm.Model
	Username    string `gorm:"uniqueIndex;not null"`
	Email       string `gorm:"uniqueIndex;not null"`
	Password    string `gorm:"not null"`
	DisplayName string
	Bio         string
	Github      string
	ApiKey      string
	PrivateKey  string
	PublicKey   string
	Avatar      string
	Banner      string
	Access      Access
}

// CreateUser create a user entry in the user's table
func CreateUser(user *User) *gorm.DB {
	return db.Executor.Create(user)
}

// FindUser searches the user's table with the condition given
func FindUser(dest interface{}, conds ...interface{}) *gorm.DB {
	return db.Executor.Model(&User{}).Take(dest, conds...)
}

// FindUserById searches the user's table with the id given
func FindUserById(id uint) (*User, error) {
	user := &User{}
	if err := FindUser(user, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &fiber.Error{
				Code:    contracts.StatusNotFound,
				Message: "user not found",
			}
		} else {
			return nil, &fiber.Error{
				Code:    contracts.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	}

	return user, nil
}

// FindUserByEmail searches the user's table with the email given
func FindUserByEmail(email string) (*User, error) {
	user := &User{}
	if err := FindUser(user, "email = ?", email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &fiber.Error{
				Code:    contracts.StatusNotFound,
				Message: "user not found",
			}
		} else {
			return nil, &fiber.Error{
				Code:    contracts.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	}

	return user, nil
}

// FindUserByUsername searches the user's table with the name given
func FindUserByUsername(username string) (*User, error) {
	user := &User{}
	if err := FindUser(user, "username = ?", username).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &fiber.Error{
				Code:    contracts.StatusNotFound,
				Message: fmt.Sprintf("user '%s' not found", username),
			}
		} else {
			return nil, &fiber.Error{
				Code:    contracts.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	}

	return user, nil
}

// UpdateProfile updates the user's profile with the info given
func UpdateProfile(userId interface{}, data interface{}) *gorm.DB {
	return db.Executor.Model(&User{}).Where("id = ?", userId).Updates(data)
}
