package repos

import (
	"db"

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
func FindUserById(dest interface{}, id uint) *gorm.DB {
	return FindUser(dest, "id = ?", id)
}

// FindUserByEmail searches the user's table with the email given
func FindUserByEmail(dest interface{}, email string) *gorm.DB {
	return FindUser(dest, "email = ?", email)
}

// FindUserByUsername searches the user's table with the name given
func FindUserByUsername(dest interface{}, name string) *gorm.DB {
	return FindUser(dest, "username = ?", name)
}

// UpdateProfile updates the user's profile with the info given
func UpdateProfile(userId interface{}, data interface{}) *gorm.DB {
	return db.Executor.Model(&User{}).Where("id = ?", userId).Updates(data)
}
