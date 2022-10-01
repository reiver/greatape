package repos

import (
	"contracts"
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
func (repo *repository) CreateUser(user *User) error {
	if err := repo.Storage.Create(user).Error; err != nil {
		return err
	}

	return nil
}

// FindUser searches the user's table with the condition given
func (repo *repository) FindUser(conds ...any) (*User, error) {
	dest := &User{}
	if err := repo.Storage.Model(dest).Take(dest, conds...).Error; err != nil {
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
func (repo *repository) FindUserById(id uint) (*User, error) {
	return repo.FindUser("id = ?", id)
}

// FindUserByEmail searches the user's table with the email given
func (repo *repository) FindUserByEmail(email string) (*User, error) {
	return repo.FindUser("email = ?", email)
}

// FindUserByUsername searches the user's table with the name given
func (repo *repository) FindUserByUsername(username string) (*User, error) {
	return repo.FindUser("username = ?", username)
}

// UpdateProfile updates the user's profile with the info given
func (repo *repository) UpdateProfile(userId interface{}, data interface{}) error {
	if err := repo.Storage.Model(&User{}).Where("id = ?", userId).Updates(data).Error; err != nil {
		return err
	}

	return nil
}
