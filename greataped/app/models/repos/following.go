package repos

import (
	"db"

	"gorm.io/gorm"
)

// Following struct defines an account that the user follows
type Following struct {
	gorm.Model
	Target string `gorm:"not null"`
	Handle string `gorm:"not null"`
}

// CreateFollowing creates a new entry in the following's table
func CreateFollowing(following *Following) error {
	if err := db.Executor.Create(following).Error; err != nil {
		return err
	}

	return nil
}

// FindFollowing finds what accounts the user is following
func FindFollowing(userIden interface{}) ([]Following, error) {
	followings := &[]Following{}
	if err := db.Executor.Model(&Following{}).Find(followings, "`target` = ?", userIden).Error; err != nil {
		return *followings, err
	}

	return *followings, nil
}
