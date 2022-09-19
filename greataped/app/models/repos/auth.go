package repos

import (
	"contracts"
	"db"
	"errors"

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
func CreateUser(user *User) error {
	if err := db.Executor.Create(user).Error; err != nil {
		return err
	}

	return nil
}

// FindUser searches the user's table with the condition given
func FindUser(conds ...any) (*User, error) {
	dest := &User{}
	if err := db.Executor.Model(dest).Take(dest, conds...).Error; err != nil {
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

	return dest, nil
}

// FindUserById searches the user's table with the id given
func FindUserById(id uint) (*User, error) {
	return FindUser("id = ?", id)
}

// FindUserByEmail searches the user's table with the email given
func FindUserByEmail(email string) (*User, error) {
	return FindUser("email = ?", email)
}

// FindUserByUsername searches the user's table with the name given
func FindUserByUsername(username string) (*User, error) {
	return FindUser("username = ?", username)
}

// UpdateProfile updates the user's profile with the info given
func UpdateProfile(userId interface{}, data interface{}) error {
	if err := db.Executor.Model(&User{}).Where("id = ?", userId).Updates(data).Error; err != nil {
		return err
	}

	return nil
}
